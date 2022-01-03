FROM golang:1.17.5-alpine AS builder
WORKDIR /build
COPY . .
RUN go build .

FROM alpine
WORKDIR /app
COPY --from=builder /build/commutator commutator
EXPOSE 8080
ENTRYPOINT [ "./commutator" ]