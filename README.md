# Shortened
TinyUrl-like project. Shorten long links into shareable format for easier sharing and cleaner appearance.

## Stack:

### Backend:
[![Go](https://img.shields.io/badge/Go-00ADD8?style=for-the-badge&logo=go&logoColor=white)](https://go.dev/)
[![PostgreSQL](https://img.shields.io/badge/PostgreSQL-316192?style=for-the-badge&logo=postgresql&logoColor=white)](https://www.postgresql.org/)
[![Nginx](https://img.shields.io/badge/NGINX-008000?style=for-the-badge&logo=NGINX)](https://nginx.org/ru/)

### Infrastructure and Automation:
[![Ubuntu](https://img.shields.io/badge/Ubuntu-E95420?style=for-the-badge&logo=ubuntu&logoColor=white)](https://ubuntu.com/)
[![Docker](https://img.shields.io/badge/Docker-87cefa?style=for-the-badge&logo=docker)](https://www.docker.com/)
[![Yandex.Cloud](https://img.shields.io/badge/-Yandex.Cloud-1e90ff?style=for-the-badge&logo=Yandex.Cloud)](https://cloud.yandex.ru/)

## Installation

### Docker
```zsh
docker compose up
```
```zsh 
docker compose up --build
```

### Installation for Linux & MacOS
1. Open a terminal and clone the repository to your local machine.
``` zsh
git clone https://github.com/chyornyy/shortened.git
```
2. Navigate to the project directory.
``` zsh
cd shortened
```
3. Run the setup script.
``` zsh
chmod +x setup.sh && ./setup.sh
```

This script installs and configures a shortened gRPC-server and gRPC-client by performing the following tasks:

1. Sets the password for the admin user.
2. Installs necessary packages, including PostgreSQL, Nginx, and Snapd.
3. Enables and configures a firewall, Nginx, and SSL certificate using Certbot.
4. Installs Homebrew and Golang.
5. Clones the shortened URL repository and builds the server and client applications.
6. Defines and starts two systemd unit services for the shortened URL server and client applications.
7. Enables the systemd services to start automatically on boot.

## Usage

### Endpoints
```
GET /links/:shorturl - retrieves the original URL associated with the given short URL.
```
```
POST /links/ - creates a new short URL and associates it with the provided original URL.
```
```
DELETE /links/:shorturl - deletes the association between the given short URL and its original URL.
```
### Use [Postman](https://www.postman.com/) or alternatively use [curl](https://curl.se/)

![POST](https://github.com/chyornyy/shortened/blob/main/images/POST-request.png)
``` zsh
curl -X POST '{"originalurl": "ya.ru"}' https://shortened.ru/links/
```
![GET](https://github.com/chyornyy/shortened/blob/main/images/GET-request.png)
``` zsh
curl https://shortened.ru/links/fE54KN4v_4
```
![DELETE](https://github.com/chyornyy/shortened/blob/main/images/DELETE-request.png)
``` zsh
curl -X DELETE https://shortened.ru/links/fE54KN4v_4
```

## License
This project is licensed under the MIT License - see the LICENSE file for details.

## Development Team
#### [Aleksandr Chyornyy](https://github.com/chyornyy) - Backend
