FROM golang:alpine

WORKDIR /build

COPY go.mod .

COPY go.sum .

RUN go mod download

COPY . . 

RUN go build

EXPOSE 8080

CMD ["./backend"]