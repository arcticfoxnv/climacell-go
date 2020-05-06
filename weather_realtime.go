package climacell

import (
	"fmt"
	"net/url"
	"strings"
)

type RealtimeRequest struct {
	Fields     DataFieldList
	Latitude   float64
	Longitude  float64
	LocationId string
	UnitSystem UnitSystem
}

type RealtimeResponse struct {
	Latitude  float64 `json:"lat"`
	Longitude float64 `json:"lon"`

	// Weather data layer
	BaroPressure              DataPoint `json:"baro_pressure,omitempty"`
	CloudBase                 DataPoint `json:"cloud_base,omitempty"`
	CloudCeiling              DataPoint `json:"cloud_ceiling,omitempty"`
	CloudCover                DataPoint `json:"cloud_cover,omitempty"`
	Dewpoint                  DataPoint `json:"dewpoint,omitempty"`
	FeelsLike                 DataPoint `json:"feels_like,omitempty"`
	Humidity                  DataPoint `json:"humidity,omitempty"`
	MoonPhase                 DataPoint `json:"moon_phase,omitempty"`
	Precipitation             DataPoint `json:"precipitation,omitempty"`
	PrecipitationAccumulation DataPoint `json:"precipitation_accumulation,omitempty"`
	PrecipitationProbability  DataPoint `json:"precipitation_probability,omitempty"`
	PrecipitationType         DataPoint `json:"precipitation_type,omitempty"`
	SatelliteCloud            DataPoint `json:"satellite_cloud,omitempty"`
	Sunrise                   DataPoint `json:"sunrise,omitempty"`
	Sunset                    DataPoint `json:"sunset,omitempty"`
	SurfaceShortwaveRadiation DataPoint `json:"surface_shortwave_radiation,omitempty"`
	Temp                      DataPoint `json:"temp,omitempty"`
	Visibility                DataPoint `json:"visibility,omitempty"`
	WeatherCode               DataPoint `json:"weather_code,omitempty"`
	WeatherGroups             DataPoint `json:"weather_groups,omitempty"`
	WindDirection             DataPoint `json:"wind_direction,omitempty"`
	WindGust                  DataPoint `json:"wind_gust,omitempty"`
	WindSpeed                 DataPoint `json:"wind_speed,omitempty"`

	// Air Quality
	PM25                  DataPoint `json:"pm25,omitempty"`
	PM10                  DataPoint `json:"pm10,omitempty"`
	O3                    DataPoint `json:"o3,omitempty"`
	NO2                   DataPoint `json:"no2,omitempty"`
	CO                    DataPoint `json:"co,omitempty"`
	SO2                   DataPoint `json:"so2,omitempty"`
	EpaAqi                DataPoint `json:"epa_aqi,omitempty"`
	EpaHealthConcern      DataPoint `json:"epa_health_concern,omitempty"`
	EpaPrimaryPollutant   DataPoint `json:"epa_primary_pollutant,omitempty"`
	ChinaAqi              DataPoint `json:"china_aqi,omitempty"`
	ChinaHealthConcern    DataPoint `json:"china_health_concern,omitempty"`
	ChinaPrimaryPollutant DataPoint `json:"china_primary_pollutant,omitempty"`
}

func (r *RealtimeRequest) ToQuery(defaults *RequestDefaults) *url.Values {
	q := &url.Values{}

	q.Add("lat", fmt.Sprintf("%g", r.Latitude))
	q.Add("lon", fmt.Sprintf("%g", r.Longitude))

	if r.LocationId != "" {
		q.Add("location_id", r.LocationId)
	}

	if r.UnitSystem != "" {
		q.Add("unit_system", r.UnitSystem.String())
	} else if defaults.UnitSystem != "" {
		q.Add("unit_system", defaults.UnitSystem.String())
	}

	if len(r.Fields) > 0 {
		q.Add("fields", strings.Join(r.Fields.Strings(), ","))
	}

	return q
}

func (c *Client) RealtimeWeather(request *RealtimeRequest) (*RealtimeResponse, error) {
	endpoint := "weather/realtime"
	req, _ := c.newGetRequest("v3", endpoint, request)

	data := new(RealtimeResponse)
	if err := c.do(req, data); err != nil {
		return nil, err
	}

	return data, nil
}
