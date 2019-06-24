package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"teller/data"
)

type CityStore interface {
	CitiesAirQuality() []data.AirQuality
	CityAirQuality(index string) data.AirQuality
	RecordAirQuality(air data.AirQuality)
}

type InMemoryCityStore struct {
	store map[string]data.AirQuality
}

func (i *InMemoryCityStore) CitiesAirQuality() []data.AirQuality {
	values := []data.AirQuality{}
	for _, value := range i.store {
		values = append(values, value)
	}
	return values
}

func (i *InMemoryCityStore) RecordAirQuality(air data.AirQuality ) {
	i.store[air.IndexCityVHash]=air
}


func (i *InMemoryCityStore) CityAirQuality(index string) data.AirQuality {
	return i.store[index]
}


var in = &InMemoryCityStore { store: map[string]data.AirQuality{} }

func GetCities(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(in.CitiesAirQuality())
}

func GetCity(w http.ResponseWriter, r *http.Request) {
	index := r.URL.Path[len("/cities/"):]
	json.NewEncoder(w).Encode(in.CityAirQuality(index))

}

func RecordCity(w http.ResponseWriter, r *http.Request) {
	var air data.AirQuality
	if r.Body != nil {
		err := json.NewDecoder(r.Body).Decode(&air)
		if err != nil {
			log.Println(err)
			fmt.Fprintf(w, `{"status":"error", "description":"%s"}`, err)
			return
		}
	} else {
		fmt.Fprintf(w, `{"status":"error", "description":"Body is nil"}`)
		return
	}

	in.RecordAirQuality(air)
}

