package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"log"
	"net"
	"sync"

	db "shortened/database"
	md "shortened/model"
	pb "shortened/shorturl_proto"
	ul "shortened/utils"

	"google.golang.org/grpc"
)

var (
	database string
	port     = flag.Int("port", 50051, "gRPC server port")
)

type server struct {
	pb.UnimplementedLinkServiceServer
	mu    sync.Mutex
	links map[string]*md.Link
}

func (s *server) CreateLink(ctx context.Context, req *pb.CreateLinkRequest) (*pb.CreateLinkResponse, error) {
	log.Println("Create Link")
	link := req.GetLink()
	shortUrl, err := ul.EncodeToBase63(link.GetOriginalurl())
	if err != nil {
		return nil, fmt.Errorf("couldn't encode url: %w", err)
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	if _, ok := s.links[shortUrl]; ok {
		return nil, errors.New("link with this shorturl already exists")
	}

	data := md.Link{
		ShortUrl:    shortUrl,
		OriginalUrl: link.GetOriginalurl(),
	}
	s.links[shortUrl] = &data

	return &pb.CreateLinkResponse{
		Link: &pb.Link{
			Shorturl: shortUrl,
		},
	}, nil
}

func (s *server) GetLink(ctx context.Context, req *pb.ReadLinkRequest) (*pb.ReadLinkResponse, error) {
	log.Println("Read Link", req.GetShorturl())
	s.mu.Lock()
	defer s.mu.Unlock()

	if link, ok := s.links[req.GetShorturl()]; ok {
		return &pb.ReadLinkResponse{
			Link: &pb.Link{
				Originalurl: link.OriginalUrl,
			},
		}, nil
	}

	return nil, errors.New("link not found")
}

func (s *server) DeleteLink(ctx context.Context, req *pb.DeleteLinkRequest) (*pb.DeleteLinkResponse, error) {
	log.Println("Delete Link", req.GetShorturl())
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, ok := s.links[req.GetShorturl()]; !ok {
		return nil, errors.New("link not found")
	}

	delete(s.links, req.GetShorturl())

	return &pb.DeleteLinkResponse{
		Success: true,
	}, nil
}

func init() {
	flag.StringVar(&database, "database", "in-memory", "select database: postgresql or in-memory")
	flag.Parse()

	switch database {
	case "postgresql":
		db.DatabaseConnection()
	default:
		log.Println("Using In-Memory database")
	}
}

func main() {
	log.Println("gRPC server running ...")
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	server := server{links: make(map[string]*md.Link)}
	pb.RegisterLinkServiceServer(s, &server)

	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
