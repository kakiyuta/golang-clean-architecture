FROM golang:1.22.5-bullseye

RUN mkdir /go/src/app

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
