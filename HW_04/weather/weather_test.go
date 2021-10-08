package weather

import "testing"

func TestWeather(t *testing.T) {
	v := GetWeather("Kyoto")
	if v == "" {
		t.Error("Expected weather, got empty string", v)
	}
}
