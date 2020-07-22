package controller

import (
	"database/sql"
	"github.com/gorilla/mux"

	"kiwi/model"
	"kiwi/repository/smoothie"
	"kiwi/util"

	"net/http"
)

//CONTROLLER it is there to manage sql requests
type Controller struct {}
var smoothies []model.Smoothie

func (c Controller) GetSmoothies(db *sql.DB) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		var smoothie model.Smoothie
		var mError model.Error

		//{} means some sort of initialization
		smoothies = []model.Smoothie{}
		smoothieRepo := smoothieRepository.SmoothieRepository{}

		smoothies, err := smoothieRepo.GetSmoothies(db, smoothie, smoothies)

		if err != nil {
			mError.Message = "Server Error"
			util.SendError(writer, http.StatusInternalServerError, mError)
			return
		}

		writer.Header().Set("Content-Type", "application/json")
		util.SendSuccess(writer, smoothies)
	}
}

func (c Controller) FilterSmoothiesByName(db *sql.DB) http.HandlerFunc {
	return func (w http.ResponseWriter, r *http.Request)  {
		var smoothie model.Smoothie
		var sError model.Error
		//Vars returns the route variables for the current request, if any.
		params := mux.Vars(r)

		smoothies = []model.Smoothie{}
		smoothieRepo := smoothieRepository.SmoothieRepository{}

		//conversions
		//log.Println(reflect.TypeOf(params["id"]))
		//& returns the memory address of the following variable. --> http://localhost:8080/books/3
		smoothieName, _ := params["smoothie"]
		smoothies, err := smoothieRepo.FilterSmoothiesByName(db, smoothie, smoothies, smoothieName)

		if err != nil {
			if err == sql.ErrNoRows {
				sError.Message = "Not found"
				util.SendError(w, http.StatusNotFound, sError)
				return
			}else {
				sError.Message = "Server error"
				util.SendError(w, http.StatusInternalServerError, sError)
				return
			}
		}

		w.Header().Set("Content-Type", "application/json")
		util.SendSuccess(w, smoothies)
	}
}

func (c Controller) FilterSmoothiesByBase(db *sql.DB) http.HandlerFunc {
	return func (w http.ResponseWriter, r *http.Request)  {
		var smoothie model.Smoothie
		var sError model.Error
		//Vars returns the route variables for the current request, if any.
		params := mux.Vars(r)

		smoothies = []model.Smoothie{}
		smoothieRepo := smoothieRepository.SmoothieRepository{}

		//conversions
		//log.Println(reflect.TypeOf(params["id"]))
		//& returns the memory address of the following variable. --> http://localhost:8080/books/3
		base, _ := params["base"]
		smoothies, err := smoothieRepo.FilterSmoothiesByBase(db, smoothie, smoothies, base)

		if err != nil {
			if err == sql.ErrNoRows {
				sError.Message = "Not found"
				util.SendError(w, http.StatusNotFound, sError)
				return
			}else {
				sError.Message = "Server error"
				util.SendError(w, http.StatusInternalServerError, sError)
				return
			}
		}

		w.Header().Set("Content-Type", "application/json")
		util.SendSuccess(w, smoothies)
	}
}

func (c Controller) FilterSmoothiesByIngredient(db *sql.DB) http.HandlerFunc {
	return func (w http.ResponseWriter, r *http.Request)  {
		var smoothie model.Smoothie
		var sError model.Error
		//Vars returns the route variables for the current request, if any.
		params := mux.Vars(r)

		smoothies = []model.Smoothie{}
		smoothieRepo := smoothieRepository.SmoothieRepository{}

		//conversions
		//log.Println(reflect.TypeOf(params["id"]))
		//& returns the memory address of the following variable. --> http://localhost:8080/books/3
		ingredient, _ := params["ingredient"]
		smoothies, err := smoothieRepo.FilterSmoothiesByIngredient(db, smoothie, smoothies, ingredient)

		if err != nil {
			if err == sql.ErrNoRows {
				sError.Message = "Not found"
				util.SendError(w, http.StatusNotFound, sError)
				return
			}else {
				sError.Message = "Server error"
				util.SendError(w, http.StatusInternalServerError, sError)
				return
			}
		}

		w.Header().Set("Content-Type", "application/json")
		util.SendSuccess(w, smoothies)
	}
}
