FROM golang:1.7 

MAINTAINER Sebastian Stjerne "s.stjerne@gmail.com"

RUN go get github.com/gin-gonic/gin

RUN mkdir /app 

ADD . /app/ 

WORKDIR /app 

EXPOSE 80

RUN go build -o main . 

CMD ["/app/main"]


