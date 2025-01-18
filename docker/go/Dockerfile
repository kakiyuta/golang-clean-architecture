FROM golang:1.23.4-bullseye

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

# hot reload
RUN go install github.com/air-verse/air@latest
