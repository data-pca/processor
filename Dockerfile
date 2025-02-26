FROM docker.io/golang:1.22.7-alpine as build

WORKDIR /app

COPY ./go.mod ./go.sum /app/
COPY . /app

RUN go build -v -o bin ./cmd/main.go

FROM docker.io/alpine:3.19

WORKDIR /app

COPY --from=build /app/bin /app/bin

CMD ["/app/bin"]

EXPOSE 8080