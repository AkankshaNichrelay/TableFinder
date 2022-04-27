package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/AkankshaNichrelay/TableFinder/internal/restaurant"
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
)

func (h *Handler) GetAllRestaurants(w http.ResponseWriter, r *http.Request) {
	resList := h.restaurants.GetAllRestaurants()
	render.JSON(w, r, resList)
}

func (h *Handler) GetRestaurant(w http.ResponseWriter, r *http.Request) {
	restaurant := chi.URLParam(r, "restaurant")
	res, err := h.restaurants.GetRestaurant(restaurant)
	if err != nil {
		json.NewEncoder(w).Encode(err)
		return
	}
	render.JSON(w, r, res)
}

func (h *Handler) AddRestaurants(w http.ResponseWriter, r *http.Request) {
	res := restaurant.Restaurants{}

	err := json.NewDecoder(r.Body).Decode(&res.List)
	if err != nil {
		log.Println("addRestaurants Err while Decoding. err:", err)
		w.Write([]byte("Internal Server Error"))
		return
	}

	resList := h.restaurants.AddRestaurants(res.List)
	render.JSON(w, r, resList)
}

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
