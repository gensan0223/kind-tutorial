# Goイメージをベースにする
FROM golang:1.24-alpine

# 作業ディレクトリを設定
WORKDIR /app

# ソースコードをコンテナにコピー
COPY . .

# Goアプリケーションをビルド
RUN go build -o main

# ポート8080を公開
EXPOSE 8080

# アプリケーションを実行
CMD ["./main"]

