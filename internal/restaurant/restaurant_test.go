package restaurant

import (
	"encoding/json"
	"testing"

	"github.com/AkankshaNichrelay/TableFinder/internal/cache"
	"github.com/AkankshaNichrelay/TableFinder/internal/database"
	"github.com/stretchr/testify/assert"
)

const res1 = `
{
	"name": "Test1",
	"address": "1123 Magazine St, New Orleans",
	"ratings": 4.5,
	"cuisine": "Italian",
	"open_hours": [
		"Mon-Fri: 10:00AM-6:00PM",
		"Sat-Sun: 11:00AM-10:00PM"
	]
}
`

// MockRestaurant used by tests to create Restautants
func MockRestaurant() *Restaurants {
	cacheNop := &cache.CacheNoop{}
	dbNop := &database.MysqlNoop{}
	cache := cache.New(cacheNop)
	res := New(cache, dbNop)
	return res
}

func TestAddRestaurantsGood(t *testing.T) {
	r := MockRestaurant()
	resList := Restaurants{}
	err := json.Unmarshal([]byte(res1), &resList.List)
	if assert.Error(t, err) {
		return
	}

	result := r.AddRestaurants(resList.List)
	assert.Equal(t, result[0].Name, "Test1", "TestAddRestaurantsGood failed")

}

func TestAddRestaurantsEmpty(t *testing.T) {
	r := MockRestaurant()
	resList := Restaurants{}
	result := r.AddRestaurants(resList.List)
	assert.Nil(t, result, "TestAddRestaurantsEmpty failed")
}

func TestGetRestaurantMiss(t *testing.T) {
	r := MockRestaurant()
	resName := "Test Missing"
	res, err := r.GetRestaurant(resName)
	assert.Error(t, err)
	assert.Nil(t, res, "TestGetRestaurantMiss failed: missing restaurant found")

}

func TestGetRestaurantHit(t *testing.T) {
	r := MockRestaurant()

	resList := Restaurants{}
	err := json.Unmarshal([]byte(res1), &resList.List)
	if assert.Error(t, err) {
		return
	}

	result := r.AddRestaurants(resList.List)
	ok := assert.Nil(t, result, "TestGetRestaurantHit AddRestaurants failed")
	if !ok {
		return
	}
	resName := "test1"
	res, err := r.GetRestaurant(resName)
	assert.NoError(t, err)
	assert.Equal(t, res.Name, resName, "TestGetRestaurantHit failed to Get")

}
