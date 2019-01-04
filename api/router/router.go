package router

import (
	"encoding/json"
	"net/http"

	"github.com/davyj0nes/contract-testing-example/api/storage"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

// New creates a new router and attaches routes
func New() *mux.Router {
	mux := mux.NewRouter()

	h := initHandlers()
	mux.HandleFunc("/users/{user}", h.userHandler)

	return mux
}

type handlers struct {
	UsersStore storage.UserStorage
	Logger     *logrus.Logger
}

func initHandlers() handlers {
	us := storage.NewUserStore()
	logger := logrus.New()

	return handlers{
		UsersStore: us,
		Logger:     logger,
	}
}

func (h handlers) userHandler(w http.ResponseWriter, r *http.Request) {
	h.Logger.Infof("received request to %v", r.URL)
	vars := mux.Vars(r)
	username := vars["user"]

	user, err := h.UsersStore.GetByName(username)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	resp, _ := json.Marshal(user)

	w.Header().Set("Content-Type", "application/json")
	w.Write(resp)
}
