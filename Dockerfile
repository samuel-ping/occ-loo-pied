FROM golang:1.23.4

WORKDIR /app

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /occ-loo-pied cmd/main.go

EXPOSE 3333

CMD ["/occ-loo-pied"]