FROM golang:1.24.1-bullseye

RUN mkdir /go/src/golang-clean-architecture

WORKDIR /go/src/golang-clean-architecture

RUN apt-get update && apt-get install -y \
    ca-certificates \
    curl \
    git \
    && rm -rf /var/lib/apt/lists/*

# マイグレーションファイルを実行するためのライブラリ
# golang-migrateのバイナリをダウンロードしてインストールする
RUN curl -L https://github.com/golang-migrate/migrate/releases/download/v4.17.1/migrate.linux-amd64.tar.gz | tar xvz \
    && mv migrate /usr/local/bin/migrate

ENV ENV=docker

# oapi-codegen
RUN go install github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen@latest
