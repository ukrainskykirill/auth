FROM golang:1.22.2-alpine AS builder

COPY . /github.com/ukrainskykirill/auth/source
WORKDIR /github.com/ukrainskykirill/auth/source

RUN go mod download
RUN go build -o ./bin/auth cmd/auth/main.go

FROM alpine:latest

WORKDIR /root/
COPY --from=builder /github.com/ukrainskykirill/auth/source/bin/auth .

ARG DB_USER
ARG DB_PASSWORD
ARG DB_HOST
ARG DB_PORT
ARG DB_DATABASE_NAME
ARG GRPC_PORT

ENV DB_USER=$DB_USER
ENV DB_PASSWORD=$DB_PASSWORD
ENV DB_HOST=$DB_HOST
ENV DB_PORT=$DB_PORT
ENV DB_DATABASE_NAME=$DB_DATABASE_NAME
ENV GRPC_PORT=$GRPC_PORT

CMD ["./auth"]