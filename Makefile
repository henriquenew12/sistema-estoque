APP=sistema-estoque
VERSION=v2.0
BACKENDPORT=8080
FRONTENDPORT=80

backend: docker-build-backend docker-run-backend
frontend: docker-build-frontend docker-run-frontend
all: backend frontend

docker-build-backend:
	docker build -t $(APP)-backend:$(VERSION) ./backend/

docker-run-backend:
	docker run -d -p $(BACKENDPORT):$(BACKENDPORT) --name $(APP)-backend -v $(PWD)/backend/produtos.db:/app/produtos.db $(APP)-backend:$(VERSION)

docker-build-frontend:
	docker build -t $(APP)-frontend:$(VERSION) ./frontend/

docker-run-frontend:
	docker run -d -p $(FRONTENDPORT):$(FRONTENDPORT) --name $(APP)-frontend $(APP)-frontend:$(VERSION)
