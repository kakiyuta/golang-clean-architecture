up:
	docker-compose up -d
down:
	docker-compose down
bash:
	docker-compose exec ec-server bash
rm-img:
	docker-compose down --rmi all

# migrate-up:
# 	docker run --rm -v migrations:/migrations --network host migrate/migrate \
# 	-path=/migrations/ \
# 	-database mysql://user:password@tcp(localhost:3306)/ec \
# 	up
