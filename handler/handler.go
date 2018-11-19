package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/akamiko/entity-sample2/driver"
	repository "github.com/akamiko/entity-sample2/repository"
	user "github.com/akamiko/entity-sample2/repository/user"
	"github.com/go-chi/chi"
)

// NewPostHandler ...
func NewUserHandler(db *driver.DB) *User {
	return &User{
		repo: user.NewSQLPostRepo(db.SQL),
	}
}

// User ...
type User struct {
	repo repository.UserRepository
}

// GetByID returns a post details
func (p *User) GetByID(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	payload, err := p.repo.GetByID(r.Context(), int(id))
	if err != nil {
		respondWithError(w, http.StatusNoContent, "Content not found")
	}

	respondwithJSON(w, http.StatusOK, payload)
}

// respondwithJSON write json response format
func respondwithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

// respondwithError return error message
func respondWithError(w http.ResponseWriter, code int, msg string) {
	respondwithJSON(w, code, map[string]string{"message": msg})
}
