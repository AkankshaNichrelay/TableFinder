package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/AkankshaNichrelay/TableFinder/internal/restaurant"
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
)

// GetAllRestaurants returns list of all Restaurants
func (h *Handler) GetAllRestaurants(w http.ResponseWriter, r *http.Request) {
	resList := h.restaurants.GetAllRestaurants()
	render.JSON(w, r, resList)
}

// GetRestaurant returns the given restaurant details if found
func (h *Handler) GetRestaurant(w http.ResponseWriter, r *http.Request) {
	restaurant := chi.URLParam(r, "restaurant")
	res, err := h.restaurants.GetRestaurant(restaurant)
	if err != nil {
		render.Status(r, http.StatusNotFound)
		render.JSON(w, r, fmt.Sprintf("%+v", err))
		return
	}
	render.JSON(w, r, res)
}

// AddRestaurants adds the given list of restaurants
func (h *Handler) AddRestaurants(w http.ResponseWriter, r *http.Request) {
	res := restaurant.Restaurants{}

	err := json.NewDecoder(r.Body).Decode(&res.List)
	if err != nil {
		log.Println("addRestaurants Err while Decoding. err:", err)
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, "Internal Server Error")
		return
	}

	resList := h.restaurants.AddRestaurants(res.List)
	render.JSON(w, r, resList)
}

// RemoveRestaurants ...
// func (h *Handler) RemoveRestaurants(w http.ResponseWriter, r *http.Request) {
// 	resIds := []int{}
// 	err := json.NewDecoder(r.Body).Decode(&resIds)
// 	if err != nil {
// 		log.Println("removeRestaurants Err while Decoding. err:", err)
// 		w.Write([]byte("Internal Server Error"))
// 		return
// 	}

// 	var deletedRes []Restaurant
// 	for ii, res := range Restaurants {
// 		for i := 0; i < len(resIds); i++ {
// 			if res.Id == resIds[i] {
// 				deletedRes = append(deletedRes, res)
// 				// TODO: buggy logic
// 				for j := ii; j < len(Restaurants)-1; j++ {
// 					Restaurants[j] = Restaurants[j+1]
// 				}
// 				Restaurants = Restaurants[:len(Restaurants)-1]
// 				// TODO: Can remove deleted from resIds to make more efficient
// 			}
// 		}
// 	}

// 	log.Println("removeRestaurants", Restaurants)
// 	render.JSON(w, r, deletedRes)
// }
