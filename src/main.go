package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"strings"
	"text/template"

	_ "github.com/mattn/go-sqlite3"
)

var (
	tmpl  *template.Template
	db    *sql.DB
)


func init() {
	var err error
	tmpl, err = template.ParseGlob("./static/*.html")
	if err != nil {
		panic(err)
	}

	db, err = sql.Open("sqlite3", "./db/data.db")
	if err != nil {
		log.Fatal(err)
	}

	// Initialiser la base de données
	err = initDB(db)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	help()
	route(
		handlerAccueil,
		handlerAbout,
		handlerLogin,
		registerHandler,
		logoutHandler,
		postHandler,
		commentHandler,
	)
}

func handleFormData(r *http.Request) {
	text := strings.ToLower(r.FormValue("String"))
	if text == "" {
		fmt.Println("Aucune valeur de formulaire n'a été trouvée")
		return
	}
	fmt.Println(text)
}
