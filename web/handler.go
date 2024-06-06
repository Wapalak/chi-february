package web

import (
	"chi_test_second"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/google/uuid"
	"html/template"
	"net/http"
)

func NewHandler(store chi_test_second.Store) *Handler {
	h := &Handler{
		Mux:   chi.NewMux(),
		store: store,
	}

	h.Use(middleware.Logger)
	h.Route("/threads", func(r chi.Router) {
		r.Get("/", h.ThreadList())
		r.Get("/new", h.ThreadsCreate())
		r.Post("/", h.ThreadsStore())

		// Применяем UUIDFromURLParam только к маршрутам, которым требуется UUID
		r.With(UUIDFromURLParam).Route("/{id}", func(r chi.Router) {
			r.Post("/delete", h.ThreadsDelete())
			r.Get("/posts", h.PostView())
		})
	})

	return h
}

type Handler struct {
	*chi.Mux

	store chi_test_second.Store
}

func (h *Handler) ThreadList() http.HandlerFunc {
	type data struct {
		Threads []chi_test_second.Thread
	}
	tmpl := template.Must(template.New("threadListHtml.html").ParseFiles("C:\\Users\\User\\GolandProjects\\" +
		"chi_test_second\\web\\templates\\threadListHtml.html"))
	return func(w http.ResponseWriter, r *http.Request) {
		tt, err := h.store.Threads()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		err = tmpl.Execute(w, data{Threads: tt})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func (h *Handler) ThreadsCreate() http.HandlerFunc {
	tmpl := template.Must(template.New("threadCreateHTML.html").ParseFiles("C:\\Users\\User\\GolandProjects\\" +
		"chi_test_second\\web\\templates\\threadCreateHTML.html"))
	return func(w http.ResponseWriter, r *http.Request) {
		tmpl.Execute(w, nil)
	}
}

func (h *Handler) ThreadsStore() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		title := r.FormValue("title")
		description := r.FormValue("description")

		if err := h.store.CreateThread(&chi_test_second.Thread{
			ID:          uuid.New(),
			Title:       title,
			Description: description,
		}); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		http.Redirect(w, r, "/threads", http.StatusFound)
	}
}

func (h *Handler) ThreadsDelete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr := chi.URLParam(r, "id")

		id, err := uuid.Parse(idStr)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		if err := h.store.DeleteThread(id); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		http.Redirect(w, r, "/threads", http.StatusFound)
	}
}

func (h *Handler) PostView() http.HandlerFunc {
	type data struct {
		Posts []chi_test_second.Post
	}
	tmpl := template.Must(template.New("postList.html").ParseFiles("C:\\Users\\User\\GolandProjects\\" +
		"chi_test_second\\web\\templates\\postList.html"))
	return func(w http.ResponseWriter, r *http.Request) {
		id, ok := r.Context().Value("threadID").(uuid.UUID)
		if !ok {
			http.Error(w, "Thread ID not found in context", http.StatusInternalServerError)
			return
		}
		tt, err := h.store.PostsByThread(id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		err = tmpl.Execute(w, data{Posts: tt})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}
