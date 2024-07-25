run:
	go run cmd/server/main.go

run-digtest:
	go run cmd/digtest/main.go

gen-oapi:
	oapi-codegen -package api -o gen/api.gen.go docs/swagger/swagger.yaml

swagger-up:
	docker-compose -f docs/swagger/docker-compose.yml up -d
	open http://localhost:8002