FROM golang:1.24

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
WORKDIR /app/auth-service
COPY ./auth-service/go.mod .
COPY ./auth-service/go.sum .

WORKDIR /app/router
COPY ./router/go.mod .

WORKDIR /app/auth-service
RUN go mod download

WORKDIR /app
COPY . .

WORKDIR /app/auth-service
RUN go build -v -o /usr/local/bin/app .

CMD ["app"]
