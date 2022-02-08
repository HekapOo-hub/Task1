FROM golang as builder

RUN mkdir "/app"
WORKDIR "/app"

COPY . .
RUN  go mod download
RUN go build


FROM alpine:latest

RUN apk --no-cache add ca-certificates
RUN apk add --no-cache libc6-compat
WORKDIR /root/
COPY --from=builder /app/Task1 .

EXPOSE 1323

CMD ["./Task1"]

