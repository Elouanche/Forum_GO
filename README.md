# Forum - Documentation

Projet réalisé en binôme.

## Sommaire
I. [Rappel du Projet](#i-rappel-du-projet)  
II. [Cadre de développement](#ii-cadre-de-développement)  
III. [Stack Technique](#iii-stack-technique)  
IV. [Comment installer le projet](#iv-comment-installer-le-projet)  
V. [Utilisation](#v-utilisation)  
VI. [Visuels de l'application web](#vi-visuels-de-lapplication-web)  
VII. [Axes d'amélioration](#vii-axes-damélioration)  
VIII. [Post Scriptum](#viii-post-scriptum)

## I. Rappel du Projet

Ce projet Forum consiste à créer un serveur en golang permettant de lancer un site web (forum), géré avec une base de données SQLite. 
Le CRUD devait être développé en Golang permettant ainsi la récupération des données sur la DataBase et l'envoi de celles-ci sur le forum.  

Nous devions permettre aux utilisateurs de s'enregistrer et de se connecter, de créer des posts, laisser des commentaires et liker/disliker 
les posts s'ils sont connectés. Si un utilisateur n'est pas connecté, il peut voir les posts.

## II. Cadre de développement

- **Établissement** : Rennes Ynov Campus
- **Niveau d'études** : Première année en Bachelor Informatique
- **Contraintes** :
    * 1 mois
    * Travail en binôme
    * Serveur en golang
    * Gestion des données via SQLite
    * Gestion des utilisateurs
    * Gestion des cookies de session

## III. Stack Technique

- **Langage** : Golang, HTML/CSS et JavaScript

## IV. Comment installer le projet

### 1 - Prérequis

Avant toute chose, il vous faut un compilateur C. Si vous n'en avez pas, voici les étapes à suivre :  

- Installer le ZIP en cliquant sur ce [lien](https://github.com/brechtsanders/winlibs_mingw/releases/download/14.1.0posix-18.1.5-11.0.1-ucrt-r1/winlibs-x86_64-posix-seh-gcc-14.1.0-mingw-w64ucrt-11.0.1-r1.zip)  
- Décompresser le dossier à l'intérieur  
- Déplacer ce dossier dans **C:\Program Files**  
- Trouver le dossier **bin** puis copier le chemin  
- Dans votre barre de recherche Windows, chercher *"Variables d'environnement"* et cliquer sur **Modifier les variables d'environnement système**  
- Cliquer sur **Variables d'environnement**  
- Dans *"Variables système"*, sélectionner **Path** puis cliquer sur **Modifier**  
- Coller le lien préalablement copié à la fin du fichier puis cliquer sur *"Ok"* afin de fermer les différentes fenêtres  
- Ouvrez ensuite vôtre cmd, VisualStudioCode, ou votre IDE habituel
- Copier cette commande dans votre terminal VSCode : **$env:CGO_ENABLED=1**  
- Attendre la fin de l'exécution de la commande puis rédemarrer votre IDE  

Vous êtes désormais prêt à utiliser ce projet.

### 2 - Cloner le répertoire

Placez-vous dans le dossier de votre choix afin de cloner le répertoire.
Ouvrez ensuite vôtre cmd, VisualStudioCode, ou votre IDE habituel, puis utilisez la commande ``git clone https://github.com/Elouanche/Forum_GO.git``.

## V. Utilisation

Sasissez à la racine du projet : **make build** puis **make run** afin de compiler et lancer le projet. 
Une fois le jeu essayé, n'oublier pas d'utiliser **make clean**.
Si votre projet est compilé et que vous souhaitez restart le jeu vous pouvez faire : **make restart**.

## VI. Visuels de l'application web

Voici l'interface générale du Forum du Hash :  

![Page d'accueil](ressources/accueil.png)
![Page de connexion](ressources/login.png)

## VII. Axes d'amélioration

### 1 - Likes/Dislikes

L'implantation des likes et des dislikes a été faite, et nous pouvons ainsi liker un post. Ce like (ou dislike) sera retrouvé dans la base de données 
lié à un utilisateur et à un post.  
L'incrémentation sur le site ne se fait pas encore, ce qui est notre principal axe de progrès sur ce projet.

### 2 - Docker

Actuellement, il est possible de créer une image Docker via la commande :  
``docker build -t forum-app .``  
Le problème rencontré actuellement est que l'image ne se run pas encore.

## VIII. Remerciements

Je vous remercie d'avoir essayé mon projet en espérant que cela vous a plu !