package climacell

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestRealtimeRequestToQuery(t *testing.T) {
	defaults := RequestDefaults{
		UnitSystem: SI,
	}

	req := &RealtimeRequest{
		Fields: DataFieldList{
			Temp,
			FeelsLike,
		},
		Latitude:   40.7128,
		Longitude:  -74.0059,
		LocationId: "test_location",
	}

	values := req.ToQuery(&defaults)
	assert.Equal(t, "40.7128", values.Get("lat"))
	assert.Equal(t, "-74.0059", values.Get("lon"))
	assert.Equal(t, "test_location", values.Get("location_id"))
	assert.Equal(t, "si", values.Get("unit_system"))
	assert.Equal(t, "temp,feels_like", values.Get("fields"))

	req.UnitSystem = US
	values = req.ToQuery(&defaults)
	assert.Equal(t, "us", values.Get("unit_system"))
}

func TestRealtimeWeather(t *testing.T) {
	h := http.HandlerFunc(testingMockResponseHandler(t, "testdata/realtime.json"))

	httpClient, teardown := testingHTTPClient(h)
	defer teardown()

	cli := NewClient("abc123", SetHTTPClient(httpClient))

	req := &RealtimeRequest{
		Fields: DataFieldList{
			Dewpoint,
			FeelsLike,
		},
		Latitude:  40.7128,
		Longitude: -74.0059,
	}
	data, err := cli.RealtimeWeather(req)

	assert.Nil(t, err)
	assert.True(t, data.BaroPressure.Unset())
	assert.Equal(t, 20.25, data.FeelsLike.Value)
	assert.Equal(t, "C", data.FeelsLike.Units)
	assert.Equal(t, 37.0, data.EpaAqi.Value)
}
