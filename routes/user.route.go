package routes

import (
	"encoding/json"
	"net/http"

	"git.hub/villalobosbms/tprueba/db"
	"git.hub/villalobosbms/tprueba/models"
	"github.com/gorilla/mux"
	"golang.org/x/crypto/bcrypt"
)

func GetUsers(w http.ResponseWriter, r *http.Request){
	var user []models.User
	db.DB.Find(&user)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&user)
}

func GetUser(w http.ResponseWriter, r *http.Request){
	params := mux.Vars(r)
	var user models.User

	db.DB.First(&user, params["ID"])

	if user.ID == 0{
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("No se encontro el usuario"))
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&user)
}

func UpdateUser(w http.ResponseWriter, r *http.Request){
	params :=  mux.Vars(r)
	var user, nuevo models.User

	err := json.NewDecoder(r.Body).Decode(&nuevo)
	if err != nil{
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("JSON invalido"))
		return
	}

	db.DB.First(&user, params["id"])

	if user.ID == 0{
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("No se encontro el usuario"))
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(nuevo.Password), bcrypt.DefaultCost)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("No se pudo encriptar la contraseña por lo tanto no se guardara el usuario intentelo de nuevo"))
		return
	}

	user.FirstName = nuevo.FirstName
	user.SecondName = nuevo.SecondName
	user.LastNames = nuevo.LastNames
	user.Age = nuevo.Age
	user.Password = string(hashedPassword)

	db.DB.Save(&user)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&user)

}

func CreateUser(w http.ResponseWriter, r *http.Request){
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("JSON invalido"))
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("No se pudo encriptar la contraseña por lo tanto no se guardara el usuario intentelo de nuevo"))
		return
	}

	user.Password = string(hashedPassword)

	createdUser := db.DB.Create(&user)
	if createdUser.Error != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("No se pudo crear el usuario"))
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&user)
}

func DeleteUser(w http.ResponseWriter, r *http.Request){
	params := mux.Vars(r)
	var user models.User

	db.DB.First(&user, params["id"])

	if user.ID == 0{
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("No se encontro el usuario"))
		return
	}

	db.DB.Delete(&user, params["id"])
	json.NewEncoder(w).Encode(&user)
	w.WriteHeader(http.StatusOK)
}