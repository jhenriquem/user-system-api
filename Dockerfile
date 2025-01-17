# Etapa 1: Construção da aplicação
FROM golang:1.23-alpine

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . /app

EXPOSE 8000

# Comando de inicialização do container
CMD ["go","run","./server/main.go"]
