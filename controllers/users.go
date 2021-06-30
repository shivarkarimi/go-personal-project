package controllers

import (
	"net/http"
	"regexp"
	"strconv"

	"github.com/shivarkarimi/go-personal-project/models"
)

type userController struct {
	userIDPattern *regexp.Regexp
}

func (uc userController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// w.Write([]byte("Hello from user controller!"))
	if r.URL.Path == "/users" {
		switch r.Method {
		case http.MethodGet:
			uc.GetAll(w, r)
		case http.MethodPost:
			uc.Post(w, r)
		default:
			w.WriteHeader(http.StatusNotImplemented)
		}
	} else {
		matches := uc.userIDPattern.FindStringSubmatch(r.URL.Path)
		if len(matches) == 0 {
			w.WriteHeader(http.StatusNotFound)
		}

		id, err := strconv.Atoi(matches[1])

		if err != nil {
			w.WriteHeader(http.StatusNotFound)
		}

		switch r.Method {
		case http.MethodGet:
			uc.Get(id, w)
		case http.MethodPost:
			uc.Post(w, r)
		case http.MethodDelete:
			uc.Delete(w, r)
		default:
			w.WriteHeader(http.StatusNotImplemented)
		}

	}
}

func (uc userController) GetAll(w http.ResponseWriter, r *http.Request) {
	encodeResponseAsJSON(models.GetUsers(), w)
}

func (uc userController) Get(id int, w http.ResponseWriter) {
	u, err := models.GetUserByID(id)
	if err != nil {
		w.WriteHeader((http.StatusNotFound))
		return
	}
	encodeResponseAsJSON(u, w)
}

func (uc userController) Post(w http.ResponseWriter, r *http.Request) {

}

func (uc userController) Put(w http.ResponseWriter, r *http.Request) {

}

func (uc userController) Delete(w http.ResponseWriter, r *http.Request) {

}

func newUserController() *userController {
	return &userController{
		userIDPattern: regexp.MustCompile(`^/users/(\d+)/?`),
	}
}
