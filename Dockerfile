FROM golang:1.17 AS builder
WORKDIR /app
COPY . .
RUN go build .

FROM alpine
WORKDIR /app
COPY --from=builder /app/commutator commutator
ENTRYPOINT [ "./commutator" ]