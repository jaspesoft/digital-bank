# Usar la imagen oficial de Go 1.22 basada en Debian Bullseye
FROM golang:1.22-bullseye

# Instalación de dependencias necesarias
RUN apt-get update && apt-get install -y \
    nodejs \
    npm \
    wkhtmltopdf \
    fonts-liberation \
    fontconfig \
    xvfb \
    libxrender1 \
    libxext6 \
    && apt-get clean \
    && rm -rf /var/lib/apt/lists/*

# Instalar nodemon
RUN npm install -g nodemon

# Establecer el directorio de trabajo
WORKDIR /app

# Copiar los archivos de Go y descargar las dependencias
COPY go.mod go.sum ./
RUN go mod download

# Copiar el código restante de la aplicación
COPY . .

# Compilar la aplicación
RUN go build -o digital-bank cmd/main.go

# Exponer el puerto que usará la aplicación
EXPOSE 8080

# Comando para ejecutar la aplicación con nodemon
CMD ["nodemon", "--exec", "go", "run", "./cmd/main.go", "--signal", "SIGTERM"]
