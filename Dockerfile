# ==========================================
# Etapa 1: Build (Compilação)
# ==========================================
FROM golang:1.26.8-alpine AS builder

WORKDIR /app

# Copia dependências e faz o download (otimiza o cache do Docker)
COPY go.mod go.sum ./
RUN go mod download

# Copia todo o código fonte
COPY . .

# Compila a aplicação gerando um binário estático chamado "cli_app"
RUN CGO_ENABLED=0 GOOS=linux go build -o cli_app .

# ==========================================
# Etapa 2: Imagem Final (Execução)
# ==========================================
FROM alpine:latest

# Certificados SSL para caso o app faça requisições web no futuro
RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copia apenas o binário pronto da etapa anterior
COPY --from=builder /app/cli_app .

# Define o programa como ponto de entrada
ENTRYPOINT ["./cli_app"]

# Comando padrão caso nenhum argumento seja passado
CMD ["--help"]