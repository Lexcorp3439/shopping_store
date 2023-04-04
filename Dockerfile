FROM golang:1.19-alpine

WORKDIR /app

# Install make
# Install dependencies
RUN apk update && apk add make


COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY ./ /app

RUN go build -o /shopping_store

EXPOSE 8080

CMD ["make", "bin-deps"]
CMD ["make", "db-up"]
CMD [ "/shopping_store" ]
