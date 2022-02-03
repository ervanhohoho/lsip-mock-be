FROM golang:1.16-alpine as base

WORKDIR /app

ENV GO111MODULE=on CGO_ENABLED=0

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

RUN go build -o /lsip-mock

# runner image
FROM gcr.io/distroless/static:latest
WORKDIR /app
COPY --from=base /lsip-mock main

EXPOSE 8080

CMD [ "/app/main" ]