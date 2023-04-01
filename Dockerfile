FROM golang:1.16-alpine
MAINTAINER Vikas Sharma
WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY *.go ./

RUN go build -o /RestAPI-SellerApp

EXPOSE 8080

CMD [ "/RestAPI-SellerApp" ]