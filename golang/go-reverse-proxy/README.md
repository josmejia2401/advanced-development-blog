# go

A simple CRUD  application using GO and mongodb.

# Run Docker

- build image:
sudo docker build -t go-reverse-proxy --tag go-reverse-proxy .

- Run container in backgrond
mkdir -p ~/logs/go-docker
sudo docker run v ~/logs/go-docker:/app/logs --name go-reverse-proxy -d -p 443:443 -p 80:80 go-reverse-proxy
sudo docker run v ./logs:/app/logs --name go-reverse-proxy -p 443:443 -p 80:80 go-reverse-proxy