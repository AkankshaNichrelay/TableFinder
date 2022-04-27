package handler

import (
	"log"
	"net/http"

	"github.com/AkankshaNichrelay/TableFinder/internal/restaurant"
	"github.com/go-chi/chi"
)

type Handler struct {
	Router      *chi.Mux
	log         *log.Logger
	restaurants *restaurant.Restaurants
}

func New(lg *log.Logger, res *restaurant.Restaurants) *Handler {
	mux := chi.NewRouter()

	h := Handler{
		log:         lg,
		restaurants: res,
		Router:      mux,
	}

	mux.Get("/", h.getHome)
	mux.Get("/getAllResturants", h.GetAllRestaurants)
	mux.Get("/getRestaurant/{restaurant}", h.GetRestaurant)
	mux.Post("/addRestaurants", h.AddRestaurants)
	// mux.Delete("/removeRestaurants", h.RemoveRestaurants)

	return &h
}

func (h *Handler) getHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome To TableFinder!"))
}
