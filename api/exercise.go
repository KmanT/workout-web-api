package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"../model"

	"github.com/gorilla/mux"
)

// POST: 	/api/auth/exercise

func CreateExercise(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	exercise := &model.Exercise{}
	json.NewDecoder(r.Body).Decode(exercise)

	createdExercise := db.Create(&exercise)
	var errMessage = createdExercise.Error

	if createdExercise.Error != nil {
		fmt.Println(errMessage)
	}
	json.NewEncoder(w).Encode(createdExercise)
}

// GET: 	/api/auth/exercise

func FetchExercises(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	var exercises []model.Exercise
	db.Preload("auths").Find(&exercises)

	json.NewEncoder(w).Encode(exercises)
}

// GET: 	/api/auth/user/{id}/exercise

func FetchUserExercises(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	params := mux.Vars(r)
	var id = params["id"]
	var exercises []model.Exercise
	db.Where("user_id = ?", id).Find(&exercises)

	json.NewEncoder(w).Encode(exercises)
}

// GET: 	/api/auth/user/{id}/exercise/{date}

func FetchUserDateExercises(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	params := mux.Vars(r)
	var id = params["id"]
	var date = params["date"]
	var exercises []model.Exercise
	db.Where("user_id = ? AND date = ?", id, date).Find(&exercises)

	json.NewEncoder(w).Encode(exercises)
}

// GET:		/api/auth/user/{id}/dates

func FetchUserDates(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	params := mux.Vars(r)
	var id = params["id"]
	var exercises []model.Exercise
	var dates []time.Time

	db.Select("DISTINCT(date)").Order("date desc").Where("user_id = ?", id).Find(&exercises)

	for _, exercise := range exercises {
		dates = append(dates, exercise.Date)
	}

	json.NewEncoder(w).Encode(dates)
}

// PUT: 	/api/auth/exercise/{id}

func UpdateExercise(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	exercise := &model.Exercise{}
	params := mux.Vars(r)
	var id = params["id"]
	db.First(&exercise, id)
	json.NewDecoder(r.Body).Decode(exercise)
	db.Save(&exercise)
	json.NewEncoder(w).Encode(&exercise)
}

// DELETE: 	/api/auth/exercise/{id}

func DeleteExercise(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	params := mux.Vars(r)
	var id = params["id"]
	var exercise model.Exercise
	db.First(&exercise, id)
	db.Delete(&exercise)
	json.NewEncoder(w).Encode("Exercise deleted")
}

// GET: 	/api/auth/exercise/{id}

func GetExercise(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	params := mux.Vars(r)
	var id = params["id"]
	var exercise model.Exercise
	db.First(&exercise, id)
	json.NewEncoder(w).Encode(&exercise)
}
