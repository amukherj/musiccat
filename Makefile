IMAGE_NAME=musiccat
IMAGE_VERSION=latest
DHUB_USER=amukher1

all: musiccat
	sudo docker rmi musiccat:latest >/dev/null 2>&1 || echo "No older image found"
	sudo docker build -t musiccat:latest .
	sudo docker tag musiccat:latest $(DHUB_USER)/musiccat:latest

musiccat:
	CGO_ENABLED=0 GOOS=linux go build -a -o musiccat github.com/amukherj/musiccat/cmd/musiccat

