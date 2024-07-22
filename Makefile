# Nom du binaire à produire
BINARY_NAME=forum.exe

DIR_NAME = static

#chemin pour static (build)
SOURCE_PATH = ./src/$(DIR_NAME)

# chemin pour static (clean)
SOURCE_PATH2 = ./$(DIR_NAME)

#chemin de destination après avoir build
DESTINATION_PATH = ./

#chemin de destination après avoir clean
DESTINATION_PATH2 = ./src/

# Liste explicite des fichiers source
SOURCES := src/*.go \
 
# Commande de build
build:
	@echo "Construction du projet..."
	CGO_ENABLED=1 go build -o $(BINARY_NAME) $(SOURCES)
	@mv $(SOURCE_PATH) $(DESTINATION_PATH)
	
 
# Commande pour nettoyer le projet (supprimer le binaire)
clean:
	@echo "Nettoyage..."
	@mv $(SOURCE_PATH2) $(DESTINATION_PATH2)
	rm $(BINARY_NAME)
 
# Commande pour exécuter le programme
run: 
	@echo "Execution du programme..."
	./$(BINARY_NAME)

restart: clean build run

# Option 'phony' pour indiquer que 'clean', 'run', et 'build' ne sont pas des fichiers
.PHONY: build clean run restart