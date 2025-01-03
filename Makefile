up:
	docker-compose up -d
down:
	docker-compose down
sh:
	docker-compose exec ec-server bash
rm-img:
	docker-compose down --rmi all

run:
	go run cmd/server/main.go

gen-oapi:
	oapi-codegen -package api -o gen/api/api.gen.go docs/swagger/swagger.yaml

swagger-up:
	docker-compose -f docs/swagger/docker-compose.yml up -d
	open http://localhost:8002

gen-mock:
	mockgen -source=domain/repository/products.go -destination=domain/repository/products_mock.go -package=repository
	mockgen -source=domain/repository/variants.go -destination=domain/repository/variants_mock.go -package=repository
	mockgen -source=infra/db/connector.go -destination=infra/db/connector_mock.go -package=db
