FROM golang:1.19.3-alpine

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

# RUN env GOOS=linux GOARCH=arm64 go build -o bin/mixko-pay main.go
RUN go build -o ./bin/mixko-pay .

EXPOSE 8080

CMD [ "./bin/mixko-pay" ]