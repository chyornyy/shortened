#!/bin/bash

sudo passwd admin

sudo apt update -y
sudo apt install postgresql postgresql-contrib nginx snapd -y

sudo ufw allow 'Nginx Full'
sudo ufw allow OpenSSH
sudo ufw enable -y
sudo ufw status
sudo systemctl start nginx
sudo rm /etc/nginx/sites-enabled/default
sudo nano /etc/nginx/sites-enabled/default
sudo systemctl reload nginx

/bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/HEAD/install.sh)"
eval "$(/home/linuxbrew/.linuxbrew/bin/brew shellenv)"
brew install go
git clone https://github.com/chyornyy/shortened.git
cd shortened && nano .env

cd shorturl_server && go build main.go
cd ..
cd shorturl_client && go build main.go
cd

sudo bash -c 'echo "[Unit]
Description=client daemon
After=network.target

[Service]
User=admin
WorkingDirectory=/home/admin/shortened
ExecStart= /home/admin/shortened/shorturl_server/main -database=postgresql
[Install]
WantedBy=multi-user.target" > /etc/systemd/system/shorturl_server.service'

sudo bash -c 'echo "[Unit]
Description=client daemon
After=network.target

[Service]
User=admin
WorkingDirectory=/home/admin/shortened/shorturl_client/
ExecStart= /home/admin/shortened/shorturl_client/main
[Install]
WantedBy=multi-user.target" > /etc/systemd/system/shorturl_client.service'


sudo systemctl start shorturl_server
sudo systemctl start shorturl_client

sudo systemctl enable shorturl_server.service
sudo systemctl enable shorturl_client.service

sudo snap install core; sudo snap refresh core
sudo snap install --classic certbot
sudo ln -s /snap/bin/certbot /usr/bin/certbot 
sudo certbot --nginx
sudo systemctl reload nginx

sudo reboot