FROM golang:1.23-alpine3.20 as builder

WORKDIR /tmp/zadanie5
COPY . .

RUN git submodule update --recursive --init

RUN make clean
RUN make swagger
RUN make build

FROM golang:1.23-alpine3.20

WORKDIR /zadanie5
COPY --from=builder /tmp/Wawel/main .

CMD ["./main"]
EXPOSE 8080