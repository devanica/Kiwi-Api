package smoothieRepository

import (
	"database/sql"
	"kiwi/model"
	"log"
)

//REPOSITORY makes sql query
func logFatal(err error)  {
	if err != nil {
		log.Fatal(err)
	}
}

type SmoothieRepository struct {

}

func (sr SmoothieRepository) GetSmoothies(db *sql.DB, smoothie model.Smoothie, smoothies []model.Smoothie) ([]model.Smoothie, error) {
	rows, err := db.Query("select * from smoothies")
	if err != nil {
		return []model.Smoothie{}, err
	}
	for rows.Next(){
		// IT IS VERY IMPORTANT TO RESPECT ORDER OF ELEMENTS HERE, THEY MUST MATCH THE ORDER OF OBJECT
		err := rows.Scan(&smoothie.MSmoothie, &smoothie.Base,  &smoothie.Source, &smoothie.Recipe, &smoothie.Ingredient, &smoothie.Image, &smoothie.ID)
		logFatal(err)
		smoothies = append(smoothies, smoothie)
	}

	return smoothies, nil
}

func (sr SmoothieRepository) FilterSmoothiesByName(db *sql.DB, mSmoothie model.Smoothie, smoothies []model.Smoothie, smoothie string) ([]model.Smoothie, error) {
	rows, err := db.Query("select * from smoothies where lower(smoothie) like '%' || $1 || '%'", smoothie)
	if err != nil {
		return []model.Smoothie{}, err
	}
	for rows.Next(){
		err := rows.Scan(&mSmoothie.MSmoothie, &mSmoothie.Base,  &mSmoothie.Source, &mSmoothie.Recipe, &mSmoothie.Ingredient, &mSmoothie.Image, &mSmoothie.ID)
		logFatal(err)
		smoothies = append(smoothies, mSmoothie)
	}

	return smoothies, nil
}

func (sr SmoothieRepository) FilterSmoothiesByBase(db *sql.DB, smoothie model.Smoothie, smoothies []model.Smoothie, base string) ([]model.Smoothie, error) {
	rows, err := db.Query("select * from smoothies where base=$1", base)
	if err != nil {
		return []model.Smoothie{}, err
	}
	for rows.Next(){
		err := rows.Scan(&smoothie.MSmoothie, &smoothie.Base,  &smoothie.Source, &smoothie.Recipe, &smoothie.Ingredient, &smoothie.Image, &smoothie.ID)
		logFatal(err)
		smoothies = append(smoothies, smoothie)
	}

	return smoothies, nil
}

func (sr SmoothieRepository) FilterSmoothiesByIngredient(db *sql.DB, smoothie model.Smoothie, smoothies []model.Smoothie, ingredient string) ([]model.Smoothie, error) {
	rows, err := db.Query("select * from smoothies where ingredients like '%' || $1 || '%'", ingredient)
	if err != nil {
		return []model.Smoothie{}, err
	}
	for rows.Next(){
		err := rows.Scan(&smoothie.MSmoothie, &smoothie.Base,  &smoothie.Source, &smoothie.Recipe, &smoothie.Ingredient, &smoothie.Image, &smoothie.ID)
		logFatal(err)
		smoothies = append(smoothies, smoothie)
	}

	return smoothies, nil
}

