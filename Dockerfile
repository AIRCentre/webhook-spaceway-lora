FROM golang:1.20 AS builder

WORKDIR /app

COPY go.* ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

FROM scratch

COPY --from=builder /app/main /usr/local/bin/app

WORKDIR /app

EXPOSE 3000

ENTRYPOINT ["app"]
