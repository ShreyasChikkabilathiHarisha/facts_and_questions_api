package handler

import (
	"encoding/json"
	"net/http"

	"facts_and_questions_api/controller"
	"facts_and_questions_api/entity"
)

type Option func(*Handler)

type Logger interface {
	Printf(format string, v ...interface{})
}

func LogWith(logger Logger) Option {
	return func(h *Handler) {
		h.logger = logger
	}
}

type Handler struct {
	logger     Logger
	mux        *http.ServeMux
	controller *controller.Controller
}

func NewHandler(options ...Option) *Handler {
	h := &Handler{}

	for _, o := range options {
		o(h)
	}
	h.controller = controller.NewController(controller.LogWith(h.logger))

	h.mux = http.NewServeMux()
	h.mux.HandleFunc("/", h.index)

	h.mux.HandleFunc("/fetch", h.handleFetch)
	h.mux.HandleFunc("/create", h.handleCreate)

	return h
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.log("%s %s", r.Method, r.URL.Path)

	h.mux.ServeHTTP(w, r)
}

func (h *Handler) log(format string, v ...interface{}) {
	if h.logger != nil {
		h.logger.Printf(format+"\n", v...)
	}
}

func (h *Handler) index(w http.ResponseWriter, r *http.Request) {
	// if r.Method == "GET" {
	// 	h.mux.HandleFunc("/fetch", h.handleFetch)
	// 	return
	// }

	w.Write([]byte("Please specify the correct path!"))
	w.WriteHeader(http.StatusMethodNotAllowed)
}

func (h *Handler) handleFetch(w http.ResponseWriter, r *http.Request) {
	var request entity.Request
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		w.WriteHeader(http.StatusFailedDependency)
	}

	if len(request.Question) > 0 {
		res := h.controller.Fetch(request.Question)
		json.NewEncoder(w).Encode(res)
		return
	}

	res := h.controller.FetchAllQuestions()
	json.NewEncoder(w).Encode(res)
}

func (h *Handler) handleCreate(w http.ResponseWriter, r *http.Request) {
	var request entity.Request
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		w.WriteHeader(http.StatusFailedDependency)
	}

	res := h.controller.Create(request)
	json.NewEncoder(w).Encode(res)
}
