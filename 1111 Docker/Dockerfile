FROM golang:1.17.3-alpine3.14

RUN mkdir /app

WORKDIR /app

ADD . /app

RUN go get -u -d gocv.io/x/gocv

RUN go get -u github.com/jinzhu/gorm

RUN go get -u github.com/gin-gonic/gin

RUN go get -u github.com/jinzhu/gorm/dialects/mysql

RUN go build -o main

EXPOSE 8080

CMD [ "/app/main" ]

# CMD [ "./wait-for-it.sh", "mysql:3306", "--","/app/main" ]