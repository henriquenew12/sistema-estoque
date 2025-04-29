APP=sistema-estoque
VERSION=v1.0
PORT=8080

all: docker-build-backend docker-run-backend

docker-build-backend:
	docker build -t $(APP)-backend:$(VERSION) ./backend/

docker-run-backend:
	docker run -p $(PORT):$(PORT) $(APP)-backend:$(VERSION)
