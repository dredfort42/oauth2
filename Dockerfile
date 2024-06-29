FROM golang:latest AS env

COPY ./go.mod /app/go.mod
COPY ./go.sum /app/go.sum

WORKDIR /app
RUN go mod download

FROM env AS build

COPY . /app
WORKDIR /app

RUN CGO_ENABLED=0 go build -o ./auth ./cmd/auth/main.go

FROM scratch
COPY --from=build /app/auth /app/auth
COPY --from=build /app/config.ini /app/config.ini

CMD ["/app/auth"]
