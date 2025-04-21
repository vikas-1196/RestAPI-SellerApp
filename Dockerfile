FROM golang:1.16-alpine
MAINTAINER Vikas Sharma
WORKDIR /app

COPY go.mod ./
RUN go test mod download1

COPY *.go ./

RUN go build -o /Test /RestAPI-SellerApp1

EXPOSE 8080

CMD [ "/Test /RestAPI-SellerApp1" ]
