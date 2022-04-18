package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
)

type Restaurant struct {
	Name     string   `json:"name"`
	Id       int      `json:"id"`
	Location string   `json:"address"`
	Ratings  float32  `json:"ratings"`
	Cuisine  string   `json:"cuisine"`
	Hours    []string `json:"open_hours"`
}

var restaurants []Restaurant

func getHome(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("Welcome To TableFinder!"))
}

func getAllRestaurants(w http.ResponseWriter, r *http.Request) {
	render.JSON(w, r, restaurants)
}

func getRestaurant(w http.ResponseWriter, r *http.Request) {
	restaurant := chi.URLParam(r, "restaurant")
	for _, val := range restaurants {
		if val.Name == restaurant {
			render.JSON(w, r, val)
			return
		}
	}
	json.NewEncoder(w).Encode(fmt.Errorf("restaurant not found: %s", restaurant))
}

func addRestaurants(w http.ResponseWriter, r *http.Request) {
	resList := []Restaurant{}
	err := json.NewDecoder(r.Body).Decode(&resList)
	if err != nil {
		log.Println("addRestaurants Err while Decoding. err:", err)
		w.Write([]byte("Internal Server Error"))
		return
	}

	for i := 0; i < len(resList); i++ {
		if len(restaurants) != 0 {
			resList[i].Id = restaurants[len(restaurants)-1].Id + 1
		} else {
			resList[i].Id = 1
		}
		restaurants = append(restaurants, resList[i])
	}
	log.Println("addRestaurants", restaurants)
	render.JSON(w, r, resList)
}

func removeRestaurants(w http.ResponseWriter, r *http.Request) {
	resIds := []int{}
	err := json.NewDecoder(r.Body).Decode(&resIds)
	if err != nil {
		log.Println("removeRestaurants Err while Decoding. err:", err)
		w.Write([]byte("Internal Server Error"))
		return
	}

	var deletedRes []Restaurant
	for ii, res := range restaurants {
		for i := 0; i < len(resIds); i++ {
			if res.Id == resIds[i] {
				deletedRes = append(deletedRes, res)
				// TODO: buggy logic
				for j := ii; j < len(restaurants)-1; j++ {
					restaurants[j] = restaurants[j+1]
				}
				restaurants = restaurants[:len(restaurants)-1]
				// TODO: Can remove deleted from resIds to make more efficient
			}
		}
	}

	log.Println("removeRestaurants", restaurants)
	render.JSON(w, r, deletedRes)
}

func main() {

	res := Restaurant{
		Name:     "Anatolia",
		Id:       1,
		Location: "123 Magazine St, New Orleans",
		Ratings:  4.5,
		Cuisine:  "Mediterranean",
		Hours:    []string{"Mon-Fri: 10:00AM-9:00PM", "Sat-Sun: 11:00AM-10:00PM"},
	}
	restaurants = append(restaurants, res)

	mux := chi.NewRouter()
	mux.Get("/", getHome)
	mux.Get("/getAllResturants", getAllRestaurants)
	mux.Get("/getRestaurant/{restaurant}", getRestaurant)
	mux.Post("/addRestaurants", addRestaurants)
	mux.Delete("/removeRestaurants", removeRestaurants)

	go http.ListenAndServe(":3000", mux)

	fmt.Println("Listening on localhost:3000...")
	for {

	}
}
