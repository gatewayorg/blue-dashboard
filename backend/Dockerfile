FROM golang:1.14-alpine as builder
LABEL builder=blue-dashboard
WORKDIR /src

ADD . .

RUN GOOS=linux go build -o ./build/server ./cmd/server/

FROM alpine
COPY --from=builder /src/build/ /
ENTRYPOINT ["./server"]