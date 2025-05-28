APP_NAME = quote

up:
	docker-compose up -d

down:
	docker-compose down -v
install:
	go install github.com/lib/pq
	go install github.com/gorilla/mux
