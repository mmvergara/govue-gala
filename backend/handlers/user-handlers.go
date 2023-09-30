package handlers

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/mmvergara/golang-gala/backend/repository/user"
)



type User struct {
	Repo *user.RedisRepo
}


func (u *User) GetByID(w http.ResponseWriter, r *http.Request) {
	userID:= chi.URLParam(r, "id")
	id, err := uuid.Parse(userID)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	user, err := u.Repo.FindById(r.Context(), id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(user.Username))
}
