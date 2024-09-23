FROM golang:1.22-alpine

WORKDIR /app

RUN apk add --update nodejs npm

RUN npm install -g nodemon

COPY go.sum go.mod ./

COPY . .

RUN go mod download

RUN go build -o digital-bank cmd/main.go

EXPOSE 8080

CMD ["nodemon", "--exec", "go", "run", "./cmd/main.go", "--signal", "SIGTERM"]

