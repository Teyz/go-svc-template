FROM golang:1-alpine

WORKDIR /src/app

# Copiez le fichier go.mod et go.sum pour télécharger les dépendances
COPY go.mod .
COPY go.sum .

# Téléchargez les dépendances du module Go
RUN go mod download

# Copiez le code source dans le conteneur
COPY . .

# Compilez le code source
RUN go build -o main ./cmd

# Commande par défaut pour exécuter l'application
CMD ["./main"]