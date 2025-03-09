# Golang Clean Architecture Sample Project

## Tools

- [golang-migrate](https://github.com/golang-migrate/migrate) : DBマイグレーション
- [golangci-lint](https://golangci-lint.run/) : linter

## マイグレーション

マイグレーションは [golang-migrate](https://github.com/golang-migrate/migrate) を使用。


```Shell
# Create Migration file
migrate create -ext sql -dir migrations/ -seq create_variants_table

# Up
migrate -path /go/src/app/migrations/ -database "mysql://user:password@tcp(ec-db:3306)/ec" up

# Down
migrate -path /go/src/app/migrations/ -database "mysql://user:password@tcp(ec-db:3306)/ec" down 1

```

## linter

linterは [golangci-lint](https://golangci-lint.run/) を使用している。
現時点ではデフォルトの設定のまま運用し定義調整する。

## TODO
