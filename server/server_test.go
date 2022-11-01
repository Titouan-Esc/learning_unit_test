package server

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

type Tests struct {
	name          string
	server        *httptest.Server
	response      *Weather
	expectedError error
}

func TestGetWeather(t *testing.T) {
	dataTest := []Tests{
		{
			name: "basic-request",
			server: httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
				writer.WriteHeader(http.StatusOK)
				writer.Write([]byte(`{ "city" : "Bordeaux", "forecast" : "sunny"}`))
			})),
			response: &Weather{
				City:     "Bordeaux",
				Forecast: "sunny",
			},
			expectedError: nil,
		},
	}

	for _, value := range dataTest {
		t.Run(value.name, func(t *testing.T) {
			defer value.server.Close()

			resp, err := GetWeather(value.server.URL)

			if !reflect.DeepEqual(resp, value.response) {
				t.Errorf("FAILED: expected %v, got %v\n", value.response, resp)
			}
			if !errors.Is(err, value.expectedError) {
				t.Errorf("Expected error FAILED: expected %v, got %v\n", value.expectedError, err)
			}
		})
	}
}
