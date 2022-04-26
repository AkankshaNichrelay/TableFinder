package restaurant

import (
	"fmt"
	"log"
)

type Restaurant struct {
	Name     string   `json:"name"`
	Id       int      `json:"id"`
	Location string   `json:"address"`
	Ratings  float32  `json:"ratings"`
	Cuisine  string   `json:"cuisine"`
	Hours    []string `json:"open_hours"`
}

type Restaurants struct {
	List []Restaurant
}

func New() *Restaurants {
	res := Restaurants{}
	return &res
}

func (r *Restaurants) GetAllRestaurants() *Restaurants {
	return r
}

func (r *Restaurants) GetRestaurant(restaurantName string) (*Restaurant, error) {
	for _, res := range r.List {
		if res.Name == restaurantName {
			return &res, nil
		}
	}
	return nil, fmt.Errorf("restaurant not found: %s", restaurantName)
}

func (r *Restaurants) AddRestaurants(resList []Restaurant) []Restaurant {
	for i := 0; i < len(resList); i++ {
		if len(r.List) != 0 {
			resList[i].Id = r.List[len(r.List)-1].Id + 1
		} else {
			resList[i].Id = 1
		}
		r.List = append(r.List, resList[i])
	}
	log.Println("addRestaurants", r.List)
	return resList
}
