FROM golang:1.19.5

# ホストの./backend/をコンテナの/appにコピー
COPY backend/ /app/ 

# ポートを開放
EXPOSE 8080
# カレントディレクトリをセット
WORKDIR /app

RUN go mod download
RUN go mod tidy
RUN go build
