FROM golang:latest AS env

COPY ./go.mod /app/go.mod
COPY ./go.sum /app/go.sum

WORKDIR /app
RUN go mod download

FROM env AS build

COPY . /app
WORKDIR /app

RUN CGO_ENABLED=0 go build -o ./auth ./...

FROM scratch
COPY --from=build /app/auth /auth
COPY --from=build /app/config.ini /config.ini

CMD ["/auth"]
