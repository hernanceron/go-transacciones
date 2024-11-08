# Etapa de compilación
FROM golang:1.23-alpine AS builder

# Instalar dependencias de compilación
RUN apk add --no-cache git

# Establecer el directorio de trabajo
WORKDIR /app

# Copiar los archivos go.mod y go.sum primero para aprovechar la cache de dependencias
COPY go.mod go.sum ./
RUN go mod download

# Copiar el código fuente
COPY . .

# Compilar la aplicación
# CGO_ENABLED=0 para una compilación estática
# -ldflags="-w -s" para reducir el tamaño del binario
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o main ./main.go

# Etapa final
FROM alpine:3.19

# Añadir certificados CA para HTTPS
RUN apk --no-cache add ca-certificates

# Crear un usuario no root para mayor seguridad
RUN adduser -D -g '' appuser

WORKDIR /app/

# Copiar el binario compilado desde la etapa builder
COPY --from=builder /app/main .

# Cambiar la propiedad del binario al usuario no root
RUN chown -R appuser:appuser /app

# Cambiar al usuario no root
USER appuser

# Exponer el puerto que usa tu aplicación
EXPOSE 8080

# Comando para ejecutar la aplicación
CMD ["./main"]