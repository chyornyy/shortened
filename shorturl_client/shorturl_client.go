package main

import (
	"flag"
	"log"
	"net/http"

	md "shortened/model"
	pb "shortened/shorturl_proto"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	addr = flag.String("addr", "localhost:50051", "the address to connect to")
)

func main() {
	flag.Parse()
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	defer conn.Close()
	client := pb.NewLinkServiceClient(conn)

	router := gin.Default()
	publicRoutes := router.Group("/links")
	publicRoutes.GET("/:shorturl", func(ctx *gin.Context) {
		shorturl := ctx.Param("shorturl")
		res, err := client.GetLink(ctx, &pb.ReadLinkRequest{Shorturl: shorturl})
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{
				"message": err.Error(),
			})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{
			"link": res.Link,
		})
	})
	publicRoutes.POST("/", func(ctx *gin.Context) {
		var link md.Link

		err := ctx.ShouldBind(&link)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err,
			})
			return
		}
		data := &pb.Link{
			Shorturl:    link.ShortUrl,
			Originalurl: link.OriginalUrl,
		}
		res, err := client.CreateLink(ctx, &pb.CreateLinkRequest{
			Link: data,
		})
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err,
			})
			return
		}
		ctx.JSON(http.StatusCreated, gin.H{
			"link": res.Link,
		})
	})
	publicRoutes.DELETE("/:shorturl", func(ctx *gin.Context) {
		shorturl := ctx.Param("shorturl")
		res, err := client.DeleteLink(ctx, &pb.DeleteLinkRequest{Shorturl: shorturl})
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		if res.Success == true {
			ctx.JSON(http.StatusOK, gin.H{
				"message": "Link deleted successfully",
			})
			return
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": "error deleting link",
			})
			return
		}
	})
	router.Run(":8080")
}
