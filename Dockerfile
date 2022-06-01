FROM golang:1.17-alpine

WORKDIR /app

# Download necessary Go modules
COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . ./

RUN go build -o /go-products-example-ddd

EXPOSE 8080

CMD [ "/go-products-example-ddd" ]