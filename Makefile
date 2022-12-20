build:
	@go build -o bin/restapi

run: build
	docker-compose up -d
	@./bin/restapi
	
stop: 
	docker-compose down

test:
	@go test -v ./..