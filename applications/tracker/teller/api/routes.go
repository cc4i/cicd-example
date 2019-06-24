package api

import (
	"github.com/gorilla/mux"
	"net/http"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route


func NewRouter() *mux.Router {

	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.HandlerFunc)
	}

	return router
}


var routes = Routes{

	Route{
		"AllCitiesAirQuality",
		"GET",
		"/cities",
		GetCities,
	},
	Route{
		"CityAirQuality",
		"GET",
		"/cities/{index_city_v_hash}",
		GetCity,
	},
	Route{
		"AirQualityFeed",
		"POST",
		"/cities/{city}",
		RecordCity,
	},
}
