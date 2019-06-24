package api

import (
	"teller/data"
	"testing"
)

func TestInMemoryCityStore_RecordAirQuality(t *testing.T) {

	t.Run("Record air quality into store.", func(t *testing.T) {

		air := data.AirQuality{
			IndexCityVHash: "175bf71c6370151071d139083b1a95ac3af7148d",
			IndexCity: "Demo_9999",
			StationIndex: 9999,
			AQI: 999,
			City: "Demo",
			CityCN: "Demo",
			Latitude: "99.9999",
			Longitude: "99.9999",
			Co: "99.9999",
			H: "99.9999",
			No2: "99.9999",
			P: "99.9999",
			Pm10: "99.9999",
			Pm25: "99.9999",
			So2: "99.9999",
			T: "99.9999",
			W: "99.9999",
			S: "2019-12-30 21:00:00",
			TZ: "+08:00",
			V: 999999,
		}
		in.RecordAirQuality(air)

		data := in.CitiesAirQuality()
		if len(data) !=1 {
			t.Errorf("Length of array expected '1' got %d\n", air.AQI)
		}
		if data[0].IndexCityVHash !="175bf71c6370151071d139083b1a95ac3af7148d" {
			t.Errorf("IndexCityVHash expected '175bf71c6370151071d139083b1a95ac3af7148d' got %d\n", air.AQI)
		}

	})
}

func TestInMemoryCityStore_CityAirQuality(t *testing.T) {
	t.Run("Query air quality by city index.", func(t *testing.T) {
		air := in.CityAirQuality("175bf71c6370151071d139083b1a95ac3af7148d")

		if air.IndexCity != "Demo_9999" {
			t.Errorf("IndexCity expected 'Demo_9999' got %s\n", air.IndexCity)
		}

		if air.AQI != 999 {
			t.Errorf("AQI expected '999' got %d\n", air.AQI)
		}
	})
}

func TestInMemoryCityStore_CitiesAirQuality(t *testing.T) {
	t.Run("Query air quality of all cities.", func(t *testing.T) {
		air := in.CitiesAirQuality()
		if len(air)!=1 {
			t.Errorf("Length of city air quality expected '1' got %d\n", len(air))
		}

		if air[0].IndexCity != "Demo_9999" {
			t.Errorf("IndexCity of first element in array expected 'Demo_9999' got %s\n", air[0].IndexCity)
		}

		if air[0].AQI != 999 {
			t.Errorf("AQI of first element in array expected '999' got %d\n", air[0].AQI)
		}
	})
}

