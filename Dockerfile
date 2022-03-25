FROM golang:alpine as build-env
COPY . /src
WORKDIR /src
RUN go build -o gowon-talk

FROM alpine:3.15.2
WORKDIR /app
COPY --from=build-env /src/gowon-talk /app/
ENTRYPOINT ["./gowon-talk"]
