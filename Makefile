APP=sistema-estoque
VERSION=v2.0
BACKENDPORT=8080
FRONTENDPORT=80

all: docker-build-backend docker-build-frontend docker-run-backend docker-run-frontend

docker-build-backend:
	docker build -t $(APP)-backend:$(VERSION) ./backend/

docker-run-backend:
	docker run -d -p ${BACKENDPORT}:${BACKENDPORT} --name ${APP}-backend ${APP}-backend:${VERSION}

docker-build-frontend:
	docker build -t $(APP)-frontend:$(VERSION) ./frontend/

docker-run-frontend:
	docker run -d -p ${FRONTENDPORT}:${FRONTENDPORT} --name ${APP}-frontend ${APP}-frontend:${VERSION}
