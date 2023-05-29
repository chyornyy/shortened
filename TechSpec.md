# Technical Specification

Implement a service that provides an API for creating shortened links.

The link should be:
- Unique; only one shortened link should refer to one original URL;
- 10 characters long;
- Composed of Latin letters in both lowercase and uppercase, numbers, and the _ (underscore) symbol

The service should be written in Go and accept the following HTTP requests:
1. HTTP POST method that saves the original URL to the database and returns the shortened link.
2. HTTP GET method that accepts the shortened URL and returns the original URL.
Bonus task:
Implement the service using gRPC protocol, that is, create a proto file and implement the service with two corresponding endpoints.


The solution should meet the following requirements:
- The service is distributed as a Docker image.
- The expected storage solutions are in-memory and PostgreSQL. The storage type is specified as a parameter when launching the service.
-The implemented functionality is covered by unit tests.

The result should be provided in the form of a public repository on github.com