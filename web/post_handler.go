package web

//
//import (
//	"chi_test_second"
//	"github.com/go-chi/chi"
//	"github.com/go-chi/chi/middleware"
//)
//
//func NewPostHandler(store chi_test_second.Store) *PostHandler {
//	h := &Handler{
//		Mux:   chi.NewMux(),
//		store: store,
//	}
//
//	h.Use(middleware.Logger)
//	h.Route("/threads", func(r chi.Router) {
//
//	})
//
//	return h
//}
//
//type PostHandler struct {
//	*chi.Mux
//
//	post chi_test_second.Store
//}
