# Estágio de compilação
FROM golang:1.21 as build

WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 go build -o server main.go

# Estágio de criação da imagem final
FROM alpine:latest

WORKDIR /app

# Copiar binário do estágio de compilação
COPY --from=build /app/server .

# Usar um usuário não privilegiado
USER nobody:nogroup

CMD ["./server"]

