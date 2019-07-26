package controller

import (
	"encoding/json"
	"net/http"
	"rest-example-go/entity"
	"rest-example-go/repository"
	"strconv"

	"github.com/gorilla/mux"
)

// UserController type ...
type UserController struct {
	r *repository.UserRepository
}

// NewUser is a function that setups the handlers.
func NewUser(r *repository.UserRepository) *UserController {
	return &UserController{r}
}

// FindAll lists Users.
func (u *UserController) FindAll(w http.ResponseWriter, r *http.Request) {
	data, err := u.r.FindAll()
	if err != nil {
		json.NewEncoder(w).Encode(err)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(data)
}

// Add user
func (u *UserController) Add(w http.ResponseWriter, r *http.Request) {

	//defer r.Body.Close()

	var user entity.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		json.NewEncoder(w).Encode(err)
		return
	}

	err := u.r.Add(&user)
	if err != nil {
		json.NewEncoder(w).Encode(err)
		return
	}
	//respondWithJson(w, http.StatusOK, map[string]string{"result": "success"})
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"result": "success"})
}

// Delete user
func (u *UserController) Delete(w http.ResponseWriter, r *http.Request) {

	//defer r.Body.Close()
	params := mux.Vars(r)

	id, err := strconv.ParseInt(params["id"], 10, 64)
	if err != nil {
		json.NewEncoder(w).Encode(err)
		return
	}

	err = u.r.Delete(id)

	if err != nil {
		json.NewEncoder(w).Encode(err)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"result": "success"})
}

// Update user
func (u *UserController) Update(w http.ResponseWriter, r *http.Request) {

	//defer r.Body.Close()
	params := mux.Vars(r)
	id, err := strconv.ParseInt(params["id"], 10, 64)
	if err != nil {
		json.NewEncoder(w).Encode(err)
		return
	}
	var user entity.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		json.NewEncoder(w).Encode(err)
		return
	}

	err = u.r.Update(id, &user)
	if err != nil {
		json.NewEncoder(w).Encode(err)
		return
	}
	//respondWithJson(w, http.StatusOK, map[string]string{"result": "success"})
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"result": "success"})
}

// FindOneByID user
func (u *UserController) FindOneByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.ParseInt(params["id"], 10, 64)
	if err != nil {
		json.NewEncoder(w).Encode(err)
		return
	}
	data, err := u.r.FindOneByID(id)
	if err != nil {
		json.NewEncoder(w).Encode(err)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(data)
}

// respondWithJson server answer
func respondWithJson(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
