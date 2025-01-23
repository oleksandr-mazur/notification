FROM golang:1.23.5 AS build

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .

RUN CGO_ENABLED=0 go build -o /notification ./cmd/main

FROM scratch

COPY --from=build /notification /app/notification

WORKDIR /app
EXPOSE 8080
CMD ["/app/notification"]