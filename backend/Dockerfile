FROM golang:1.20.2-alpine3.16 as builder
RUN mkdir /backend
ADD . /backend/
WORKDIR /backend
#RUN go build -o main ./src
RUN CGO_ENABLED=0 go build -ldflags="-s -w" -o main ./src

FROM alpine:3.17.2

COPY --from=builder /backend/main  /backend/

EXPOSE 6565
WORKDIR /backend
CMD ["./main"]
