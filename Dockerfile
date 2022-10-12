FROM golang:1.19

WORKDIR /usr/src/apps/math

COPY src/ .
COPY go.mod .

RUN go build -o math

CMD ["./math"]

