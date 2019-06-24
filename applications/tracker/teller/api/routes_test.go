package api

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"teller/data"
	"testing"
)

func TestGetCities(t *testing.T) {
	t.Run("Query air quality of all cities.", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/cities", nil)
		response := httptest.NewRecorder()
		GetCities(response, request)
		//got := response.Body.String()
		//fmt.Println(got)

		var air []data.AirQuality
		err := json.NewDecoder(response.Body).Decode(&air)
		if err != nil {
			t.Errorf("error: %s", err)
		}
		if air[0].IndexCityVHash != "175bf71c6370151071d139083b1a95ac3af7148d" {
			t.Errorf("IndexCityVHash expected '175bf71c6370151071d139083b1a95ac3af7148d' got %s#\n", air[0].IndexCityVHash)
		}

		if air[0].IndexCity != "Demo_9999" {
			t.Errorf("IndexCity expected 'Demo_9999' got %s\n", air[0].IndexCity)
		}

		if air[0].AQI != 999 {
			t.Errorf("AQI expected '999' got %d\n", air[0].AQI)
		}

	})
}

func TestGetCity(t *testing.T) {
	t.Run("Query air quality of a city.", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/cities/175bf71c6370151071d139083b1a95ac3af7148d", nil)
		response := httptest.NewRecorder()
		GetCity(response, request)
		//got := response.Body.String()
		//fmt.Println(got)

		var air data.AirQuality
		err := json.NewDecoder(response.Body).Decode(&air)
		if err != nil {
			t.Errorf("error: %s", err)
		}
		if air.IndexCityVHash != "175bf71c6370151071d139083b1a95ac3af7148d" {
			t.Errorf("IndexCityVHash expected '175bf71c6370151071d139083b1a95ac3af7148d' got %s\n", air.IndexCityVHash)
		}

		if air.IndexCity != "Demo_9999" {
			t.Errorf("IndexCity expected 'Demo_9999' got %s\n", air.IndexCity)
		}

		if air.AQI != 999 {
			t.Errorf("AQI expected '999' got %d\n", air.AQI)
		}

	})
}

func TestRecordCity(t *testing.T) {
	t.Run("Query air quality of a city.", func(t *testing.T) {
		air := data.AirQuality{
			IndexCityVHash: "175bf71c6370151071d139083b1a95ac3af7199d",
			IndexCity: "Demo_888",
			StationIndex: 888,
			AQI: 888,
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

		b,_ :=json.MarshalIndent(&air,"","")
		request, _ := http.NewRequest(http.MethodPost, "/cities", bytes.NewReader(b))
		response := httptest.NewRecorder()
		RecordCity(response,request)

		request2, _ := http.NewRequest(http.MethodGet, "/cities", nil)
		response2 := httptest.NewRecorder()
		GetCities(response2, request2)

		var air2 []data.AirQuality
		err := json.NewDecoder(response2.Body).Decode(&air2)
		if err != nil {
			t.Errorf("error: %s", err)
		}

		if len(air2)!=2 {
			t.Errorf("Length of array expected 2 got: %d", len(air2))
		}

	})

}

func TestInMemoryCityStore_RecordAirQuality_ErrData(t *testing.T) {
	t.Run("Query air quality of a city.", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodPost, "/cities", nil)
		response := httptest.NewRecorder()
		RecordCity(response,request)


		if !strings.Contains(response.Body.String(), "is nil") {
			t.Errorf("Response expected with 'nil'")
		}
	})
}

func TestNewRouter(t *testing.T) {
	t.Run("Query air quality of a city.", func(t *testing.T) {
		r := NewRouter()
		p := r.Path("/cities")
		tpl,_ :=p.GetPathTemplate()
		if tpl != "/cities" {
			t.Errorf("Response expected '/cities' got : %s", tpl)
		}
	})
}


