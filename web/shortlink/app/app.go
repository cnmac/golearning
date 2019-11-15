package app

import (
	"encoding/json"
	"fmt"
	"github.com/cnmac/golearning/web/shortlink/merror"
	"github.com/cnmac/golearning/web/shortlink/middleware"
	"github.com/gorilla/mux"
	"github.com/justinas/alice"
	"github.com/smartwalle/validator"
	"log"
	"net/http"
)

type App struct {
	Router      *mux.Router
	middlewares *middleware.Middleware
}

type shortenReq struct {
	URL                 string `json:"url" validate:"nonzero"`
	ExpirationInMinutes int64  `json:"expiration_in_minutes" validate:"min=0"`
}

type shortlinkResp struct {
	Shortlink string `json:"shortlink"`
}

func (a *App) Initialize() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	a.Router = mux.NewRouter()
	a.middlewares = &middleware.Middleware{}
	a.initializeRoutes()
}

func (a *App) initializeRoutes() {
	//a.Router.HandleFunc("/api/shorten", a.createShortlink).Methods("POST")
	//a.Router.HandleFunc("/api/info", a.getShortlinkInfo).Methods("GET")
	//a.Router.HandleFunc("/{shortlink:[a-zA-Z0-9]{1,11}}", a.redirect).Methods("GET")
	m := alice.New(a.middlewares.LoggingHandler, a.middlewares.RecoverHandler)
	a.Router.Handle("/api/shorten", m.ThenFunc(a.createShortlink)).Methods("POST")
	a.Router.Handle("/api/info", m.ThenFunc(a.getShortlinkInfo)).Methods("GET")
	a.Router.Handle("/{shortlink:[a-zA-Z0-9]{1,11}}", m.ThenFunc(a.redirect)).Methods("GET")
}

func (a App) createShortlink(w http.ResponseWriter, r *http.Request) {
	var req shortenReq
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		respondWithError(w, merror.StatusError{
			Code: http.StatusBadRequest,
			Err:  fmt.Errorf("parse parameters faild %v", r.Body),
		})
		return
	}
	if err := validator.Validate(req); err != nil {
		respondWithError(w, merror.StatusError{
			Code: http.StatusBadRequest,
			Err:  fmt.Errorf("validate parameters faild %v", req),
		})
		return
	}
	defer r.Body.Close()

	fmt.Println("%v", req)
}

func (a App) getShortlinkInfo(w http.ResponseWriter, r *http.Request) {
	vals := r.URL.Query()
	s := vals.Get("shortlink")
	fmt.Println("%s", s)
}

func (a App) redirect(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Println("%s", vars["shortlink"])
}

func respondWithError(w http.ResponseWriter, err interface{}) {
	switch e := err.(type) {
	case merror.Error:
		log.Printf("HTTP %d - %s", e.Status(), e)
		respondWithJson(w, e.Status(), e.Error())
	default:
		respondWithJson(w, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
	}
}

func respondWithJson(w http.ResponseWriter, code int, text interface{}) {
	resp, _ := json.Marshal(text)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(resp)
}

// Run starts listen and server
func (a App) Run(addr string) {
	log.Fatal(http.ListenAndServe(addr, a.Router))
}
