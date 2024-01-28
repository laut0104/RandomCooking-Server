FROM golang:1.19.5 as builder

# ホストの./backend/をコンテナの/appにコピー
COPY backend/ /app/ 

# カレントディレクトリをセット
WORKDIR /app

RUN go mod download
RUN go mod tidy
RUN GOOS=linux GOARCH=amd64 go build -o server .


FROM alpine:latest

# 必要なランタイムとライブラリを追加
RUN apk --no-cache add libc6-compat

# コピー元のビルダーイメージの/app/serverをコピー
COPY --from=builder /app/server /app/
COPY --from=builder /app/img/ /app/img/
COPY --from=builder /app/.env /app/
COPY --from=builder /app/.env.local /app/
COPY --from=builder /app/credentials.json /app/

# ポートを開放
EXPOSE 8080
WORKDIR /app

# 実行
CMD ["./server"]
