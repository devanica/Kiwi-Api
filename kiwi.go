package main

import (
	"database/sql"
	"github.com/gorilla/mux"
	"github.com/subosito/gotenv"
	"kiwi/controller"
	"kiwi/driver"
	"log"
	"net/http"
)

func init() {
	_ = gotenv.Load()
}

var db *sql.DB

func main() {
	db = driver.ConnectDB()
	c := controller.Controller{}

	router := mux.NewRouter()

	router.HandleFunc("/smoothies", c.GetSmoothies(db)).Methods("GET")
	router.HandleFunc("/smoothies/smoothie/{smoothie}", c.FilterSmoothiesByName(db)).Methods("GET")
	router.HandleFunc("/smoothies/base/{base}", c.FilterSmoothiesByBase(db)).Methods("GET")
	router.HandleFunc("/smoothies/ingredient/{ingredient}", c.FilterSmoothiesByIngredient(db)).Methods("GET")

	log.Fatal(http.ListenAndServe(":8080", router))
}