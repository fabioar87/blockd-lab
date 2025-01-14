ARG GO_VERSION=1.22

FROM golang:${GO_VERSION}-bullseye

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN GOARCH=amd64 GOOS=linux go build -o transactionservice .

ENTRYPOINT ["bash", "-c", "while true; do sleep 1000; done"]