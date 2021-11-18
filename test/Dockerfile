FROM golang:1.17.3-alpine3.14

RUN mkdir /app

WORKDIR /app

COPY . .
# 因為go.mod go.sum都是直接複製，所以不需要再下載，直接build
# RUN go get -u -d gocv.io/x/gocv
# RUN go get -u github.com/jinzhu/gorm
# RUN go get -u github.com/gin-gonic/gin
# RUN go get -u github.com/jinzhu/gorm/dialects/mysql

RUN go build -o main .

# 此設定只影響分享給別人使用時，告知別人你使用的port
# 不設定一樣可以執行
# EXPOSE 8080

CMD [ "/app/main" ]