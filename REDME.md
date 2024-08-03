# Golang Clean Architecture Sample Project

## Tools

- [golang-migrate](https://github.com/golang-migrate/migrate) : DBマイグレーション

## マイグレーション

マイグレーションは [golang-migrate](https://github.com/golang-migrate/migrate) を使用。


```Shell

# Up
migrate -path /go/src/app/migrations/ -database "mysql://user:password@tcp(ec-db:3306)/ec" up

# Down
migrate -path /go/src/app/migrations/ -database "mysql://user:password@tcp(ec-db:3306)/ec" down

```