FROM golang1.17:alpine AS builder
WORKDIR /app
RUN go build .

FROM alpine
WORKDIR /app
COPY --from=builder /app/commutator commutator
ENTRYPOINT [ "./commutator" ]