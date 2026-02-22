FROM golang:1.26-alpine

WORKDIR /app

RUN apk add --no-cache git make

RUN go install github.com/air-verse/air@latest
RUN go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

ENV PATH="/go/bin:${PATH}"

COPY go.mod go.sum ./
RUN go mod tidy

COPY . .

EXPOSE 8080

CMD ["air", "-c", ".air.toml"]
