FROM golang:1.16.7-buster

RUN apk add --no-cache git

WORKDIR /app/${NAME}

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN go build -o ./out/${NAME} .


EXPOSE 5050

CMD ["./out/${NAME}"]