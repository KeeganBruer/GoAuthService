FROM golang:1.24

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
WORKDIR /app/router
COPY ./router/go.mod .

WORKDIR /app/intervaljobs
COPY ./intervaljobs/go.mod .

WORKDIR /app/sql-querybuilder
COPY ./sql-querybuilder/go.mod .

WORKDIR /app/auth-service
COPY ./auth-service/go.mod .
COPY ./auth-service/go.sum .

RUN go mod download

# copy over full project for build
WORKDIR /app
COPY . .

# Build app
WORKDIR /app/auth-service
RUN go build -v -o /usr/local/bin/app .

# Run app on container start
CMD ["app"]
