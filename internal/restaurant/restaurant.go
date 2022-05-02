package restaurant

import (
	"fmt"
	"log"

	"github.com/AkankshaNichrelay/TableFinder/internal/cache"
	"github.com/AkankshaNichrelay/TableFinder/internal/database"
)

/*** Example Restaurant Values
	Restaurant{
		Name:     "Anatolia",
		Id:       1,
		Location: "123 Magazine St, New Orleans",
		Ratings:  4.5,
		Cuisine:  "Mediterranean",
		Hours:    []string{"Mon-Fri: 10:00AM-9:00PM", "Sat-Sun: 11:00AM-10:00PM"},
	}
***/

// Restaurant contains basic details of a restaurant
type Restaurant struct {
	Name     string   `json:"name"`
	Id       int      `json:"id"`
	Location string   `json:"address"`
	Ratings  float32  `json:"ratings"`
	Cuisine  string   `json:"cuisine"`
	Hours    []string `json:"open_hours"`
}

// Restaurants contains a list of all Restaurants
type Restaurants struct {
	cache *cache.Client
	db    database.DBAccessor
	List  []Restaurant // TODO: This will be replaced by cache & db. For now consider this as DB
}

// New creates new Restaurants instance
func New(cache *cache.Client, db database.DBAccessor) *Restaurants {
	res := Restaurants{cache: cache, db: db}
	return &res
}

// GetAllRestaurants returns a list of all the Restaurants in database
func (r *Restaurants) GetAllRestaurants() *Restaurants {
	return r
}

// GetRestaurant returns a Restaurant for a given restaurant name
func (r *Restaurants) GetRestaurant(restaurantName string) (*Restaurant, error) {
	for _, res := range r.List {
		if res.Name == restaurantName {
			return &res, nil
		}
	}
	return nil, fmt.Errorf("restaurant not found: %s", restaurantName)
}

// AddRestaurants adds a list of given Restaurants to database. Returns list with ids
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
