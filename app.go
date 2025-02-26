package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

type App struct {
	DB     *sql.DB
	Router *mux.Router
}

func (app *App) InitialiseDB(DBUser string, DBPass string, DBName string) error {
	connectionString := fmt.Sprintf("%v:%v@tcp(127.0.0.1:3306)/%v", DBUser, DBPass, DBName)
	_, err := sql.Open("mysql", connectionString)

	if err != nil {
		return err
	}

	app.Router = mux.NewRouter().StrictSlash(true)
	app.handleRoutes()

	return nil
}

func (app *App) Run(addr string) {
	fmt.Printf("Server running on %v", addr)
	log.Fatal(http.ListenAndServe(addr, app.Router))
}

// private to app.go file
func (app *App) getProducts(w http.ResponseWriter, r *http.Request) {

	var response []byte

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(200)
	w.Write(response)

}

func (app *App) handleRoutes() {
	app.Router.HandleFunc("/products", app.getProducts).Methods("GET")
}
