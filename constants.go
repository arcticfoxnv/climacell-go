package climacell

const (
	ClimacellV3 = "https://api.climacell.co/v3"
)

const (
	SI UnitSystem = "si"
	US UnitSystem = "us"
)

const (
	// Weather data layer
	BaroPressure              DataField = "baro_pressure"
	CloudBase                 DataField = "cloud_base"
	CloudCeiling              DataField = "cloud_ceiling"
	CloudCover                DataField = "cloud_cover"
	Dewpoint                  DataField = "dewpoint"
	FeelsLike                 DataField = "feels_like"
	Humidity									DataField = "humidity"
	MoonPhase                 DataField = "moon_phase"
	Precipitation             DataField = "precipitation"
	PrecipitationAccumulation DataField = "precipitation_accumulation"
	PrecipitationProbability  DataField = "precipitation_probability"
	PrecipitationType         DataField = "precipitation_type"
	SatelliteCloud            DataField = "satellite_cloud"
	Sunrise                   DataField = "sunrise"
	Sunset                    DataField = "sunset"
	SurfaceShortwaveRadiation DataField = "surface_shortwave_radiation"
	Temp                      DataField = "temp"
	Visibility                DataField = "visibility"
	WeatherCode               DataField = "weather_code"
	WeatherGroups             DataField = "weather_groups"
	WindDirection             DataField = "wind_direction"
	WindGust                  DataField = "wind_gust"
	WindSpeed                 DataField = "wind_speed"

	// Air Quality
	PM25                  DataField = "pm25"
	PM10                  DataField = "pm10"
	O3                    DataField = "o3"
	NO2                   DataField = "no2"
	CO                    DataField = "co"
	SO2                   DataField = "so2"
	EpaAqi                DataField = "epa_aqi"
	EpaHealthConcern      DataField = "epa_health_concern"
	EpaPrimaryPollutant   DataField = "epa_primary_pollutant"
	ChinaAqi              DataField = "china_aqi"
	ChinaHealthConcern    DataField = "china_health_concern"
	ChinaPrimaryPollutant DataField = "china_primary_pollutant"

	// Pollen

	// Road risk

	// Fire index
	FireIndex DataField = "fire-index"
)
