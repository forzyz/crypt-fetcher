FROM golang:1.23.0

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

COPY . ./

RUN go build -o /cryptfetcher

EXPOSE 3000

CMD ["/cryptfetcher"]