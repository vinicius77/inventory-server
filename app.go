package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

type App struct {
	DB     *sql.DB
	Router *mux.Router
}

func (app *App) InitialiseDB(DBUser string, DBPass string, DBName string) error {
	connectionString := fmt.Sprintf("%v:%v@tcp(127.0.0.1:3306)/%v", DBUser, DBPass, DBName)
	var err error
	app.DB, err = sql.Open("mysql", connectionString) // Assign to app.DB

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

func (app *App) getProducts(w http.ResponseWriter, r *http.Request) {
	products, err := getProducts(app.DB)

	if err != nil {
		sendError(w, http.StatusInternalServerError, err.Error())
		return
	}

	sendResponse(w, http.StatusOK, products)
}

func (app *App) getProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key, err := strconv.Atoi(vars["id"])

	if err != nil {
		sendError(w, http.StatusBadRequest, err.Error())
	}

	p := Product{ID: key}

	fmt.Println(p)

	err = p.getProduct(app.DB)

	if err != nil {
		switch err {
		case sql.ErrNoRows:
			sendError(w, http.StatusNotFound, "Product not found.")
		default:
			sendError(w, http.StatusInternalServerError, err.Error())
		}
	}

	sendResponse(w, http.StatusOK, p)
}

func sendError(w http.ResponseWriter, statusCode int, err string) {
	error_message := map[string]string{"error": err}
	sendResponse(w, statusCode, error_message)
}

func sendResponse(w http.ResponseWriter, statusCode int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(response)
}

func (app *App) handleRoutes() {
	app.Router.HandleFunc("/products", app.getProducts).Methods("GET")
	app.Router.HandleFunc("/product/{id}", app.getProduct).Methods("GET")
}
