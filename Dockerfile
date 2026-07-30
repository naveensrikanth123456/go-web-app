FROM golang:1.23.10 AS base

WORKDIR /app

COPY go.mod .

RUN go mod download

COPY . .

RUN go build -o Dev .

FROM gcr.io/distroless/base

COPY --from=base /app/Dev .

COPY --from=base /app/static ./static

EXPOSE 8080

CMD ["/Dev"]
