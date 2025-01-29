FROM golang:1.21.13-bullseye AS builder

WORKDIR /usr/src/app

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN go build -v -o /usr/local/bin/app cmd/main.go

FROM gcr.io/distroless/base-debian11

WORKDIR /usr/src/app

COPY --from=builder /usr/local/bin/app /usr/local/bin/app

EXPOSE 8080

CMD ["app"]
