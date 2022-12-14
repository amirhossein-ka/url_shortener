FROM golang:1.19-alpine AS builder

WORKDIR /app

# copy seperatly for caching
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .

RUN go build -o ./bin/urlsh .


FROM alpine:latest AS runner

ENV GOOS linux
ENV CGO_ENABLED 0

COPY --from=builder /app/bin/urlsh /usr/bin/urlsh
USER app
CMD [ "/usr/bin/urlsh", "-c", "/config.json" ]

