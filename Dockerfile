# Utiliser une image de base officielle de Golang pour la construction
FROM golang:1.21.6 AS builder

# Définir le répertoire de travail dans le conteneur
WORKDIR /app

# Copier les fichiers go.mod et go.sum pour installer les dépendances
COPY go.mod go.sum ./
RUN go mod download

# Copier le reste des fichiers de l'application
COPY . .

# Construire l'application en activant CGO pour go-sqlite3
RUN CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build -o forum ./src

# Afficher le contenu du répertoire pour le débogage
RUN ls -al /app/

# Utiliser une image de base légère pour l'exécution
FROM alpine:latest

# Installer les dépendances nécessaires pour exécuter l'application
RUN apk --no-cache add ca-certificates

# Définir le répertoire de travail dans le conteneur final
WORKDIR /root/

# Copier les fichiers nécessaires depuis l'étape de construction
COPY --from=builder /app/forum .

# Assurer que le fichier forum est exécutable
RUN chmod +x ./forum

# Afficher le contenu du répertoire pour le débogage
RUN ls -al /root/

# Copier les fichiers statiques et la base de données
COPY --from=builder /app/static ./static
COPY --from=builder /app/db ./db

# Exposer le port sur lequel l'application écoute
EXPOSE 8080

# Démarrer l'application
CMD ["./forum"]
