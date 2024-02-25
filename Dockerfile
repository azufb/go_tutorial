FROM golang:1.22.0-alpine

# 作業ディレクトリ
WORKDIR /go/src/application

# ファイルを作業ディレクトリにコピー
COPY . .

# git導入（軽量イメージのalpineには含まれないため）
RUN apk upgrade --update && apk add git