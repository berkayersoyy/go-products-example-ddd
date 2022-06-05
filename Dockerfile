FROM golang:1.17-alpine

WORKDIR /app
ARG VERSION=dev
# Download necessary Go modules
COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . ./

RUN go build -o /go-products-example-ddd -ldflags=-X=main.version=${VERSION} cmd/main.go

EXPOSE 8080

CMD [ "/go-products-example-ddd" ]