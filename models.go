package main

import "time"

// LifeIndexValueType represents a life index value with index and description
type LifeIndexValueType struct {
	Index float64 `json:"index"`
	Desc  string  `json:"desc"`
}

// LifeIndexDailyValueType represents a daily life index value
type LifeIndexDailyValueType struct {
	Date  string `json:"date"`
	Index string `json:"index"`
	Desc  string `json:"desc"`
}

// AQIValue holds AQI values for CHN and USA
type AQIValue struct {
	CHN float64 `json:"chn"`
	USA float64 `json:"usa"`
}

type HourlyDataPoint[T any] struct {
	Datetime string `json:"datetime"`
	Value    T      `json:"value"`
}

// HourlyPrecipitationDataPoint includes both value and probability
type HourlyPrecipitationDataPoint struct {
	Datetime    string  `json:"datetime"`
	Value       float64 `json:"value"`
	Probability int     `json:"probability"`
}

type HourlyWindDataPoint struct {
	Datetime  string  `json:"datetime"`
	Speed     float64 `json:"speed"`
	Direction float64 `json:"direction"`
}

type DailyRangeData struct {
	Date string  `json:"date"`
	Min  float64 `json:"min"`
	Max  float64 `json:"max"`
	Avg  float64 `json:"avg"`
}

// DailyPrecipitationData includes probability field
type DailyPrecipitationData struct {
	Date        string  `json:"date"`
	Min         float64 `json:"min"`
	Max         float64 `json:"max"`
	Avg         float64 `json:"avg"`
	Probability int     `json:"probability"`
}

type DailyWindData struct {
	Date string `json:"date"`
	Max  struct {
		Speed     float64 `json:"speed"`
		Direction float64 `json:"direction"`
	} `json:"max"`
	Min struct {
		Speed     float64 `json:"speed"`
		Direction float64 `json:"direction"`
	} `json:"min"`
	Avg struct {
		Speed     float64 `json:"speed"`
		Direction float64 `json:"direction"`
	} `json:"avg"`
}

type DailyStringIndex struct {
	Date  string `json:"date"`
	Value string `json:"value"`
}

// DailyAstro represents sunrise and sunset times
type DailyAstro struct {
	Date    string `json:"date"`
	Sunrise struct {
		Time string `json:"time"`
	} `json:"sunrise"`
	Sunset struct {
		Time string `json:"time"`
	} `json:"sunset"`
}

// DailyAQIData represents daily AQI with min, max, avg
type DailyAQIData struct {
	Date string   `json:"date"`
	Max  AQIValue `json:"max"`
	Avg  AQIValue `json:"avg"`
	Min  AQIValue `json:"min"`
}

type AirQualityRealtimeType struct {
	PM25        float64  `json:"pm25"`
	PM10        float64  `json:"pm10"`
	O3          float64  `json:"o3"`
	SO2         float64  `json:"so2"`
	NO2         float64  `json:"no2"`
	CO          float64  `json:"co"`
	AQI         AQIValue `json:"aqi"`
	Description struct {
		CHN string `json:"chn"`
		USA string `json:"usa"`
	} `json:"description"`
}

// CaiyunAPIResponse represents the structure of the Caiyun API response
type CaiyunAPIResponse struct {
	Status     string    `json:"status"`
	ErrorMsg   string    `json:"error,omitempty"`
	APIVersion string    `json:"api_version"`
	APIStatus  string    `json:"api_status"`
	Lang       string    `json:"lang"`
	Unit       string    `json:"unit"`
	TZShift    int       `json:"tzshift"`
	Timezone   string    `json:"timezone"`
	ServerTime int64     `json:"server_time"`
	Location   []float64 `json:"location"`
	Result     struct {
		Alert struct {
			Status  string `json:"status"`
			Content []struct {
				Province      string    `json:"province"`
				City          string    `json:"city"`
				County        string    `json:"county"`
				Adcode        string    `json:"adcode"`
				RegionID      string    `json:"regionId"`
				Title         string    `json:"title"`
				Code          string    `json:"code"`
				Description   string    `json:"description"`
				Status        string    `json:"status"`
				Pubtimestamp  int64     `json:"pubtimestamp"`
				Source        string    `json:"source"`
				AlertID       string    `json:"alertId"`
				Location      string    `json:"location"`
				RequestStatus string    `json:"request_status"`
				Latlon        []float64 `json:"latlon"`
			} `json:"content"`
			Adcodes []struct {
				Adcode int    `json:"adcode"`
				Name   string `json:"name"`
			} `json:"adcodes"`
		} `json:"alert"`
		Realtime struct {
			Status              string  `json:"status"`
			Temperature         float64 `json:"temperature"`
			ApparentTemperature float64 `json:"apparent_temperature"`
			Humidity            float64 `json:"humidity"`
			Cloudrate           float64 `json:"cloudrate"`
			Skycon              string  `json:"skycon"`
			Visibility          float64 `json:"visibility"`
			Dswrf               float64 `json:"dswrf"`
			Wind                struct {
				Speed     float64 `json:"speed"`
				Direction float64 `json:"direction"`
			} `json:"wind"`
			Pressure      float64 `json:"pressure"`
			Precipitation struct {
				Local struct {
					Status     string  `json:"status"`
					Datasource string  `json:"datasource"`
					Intensity  float64 `json:"intensity"`
				} `json:"local"`
				Nearest struct {
					Status    string  `json:"status"`
					Distance  float64 `json:"distance"`
					Intensity float64 `json:"intensity"`
				} `json:"nearest"`
			} `json:"precipitation"`
			AirQuality AirQualityRealtimeType        `json:"air_quality"`
			LifeIndex  map[string]LifeIndexValueType `json:"life_index"`
		} `json:"realtime"`
		Hourly struct {
			Status              string                         `json:"status"`
			Description         string                         `json:"description"`
			Precipitation       []HourlyPrecipitationDataPoint `json:"precipitation"`
			Temperature         []HourlyDataPoint[float64]     `json:"temperature"`
			ApparentTemperature []HourlyDataPoint[float64]     `json:"apparent_temperature"`
			Wind                []HourlyWindDataPoint          `json:"wind"`
			Humidity            []HourlyDataPoint[float64]     `json:"humidity"`
			Cloudrate           []HourlyDataPoint[float64]     `json:"cloudrate"`
			Pressure            []HourlyDataPoint[float64]     `json:"pressure"`
			Visibility          []HourlyDataPoint[float64]     `json:"visibility"`
			Dswrf               []HourlyDataPoint[float64]     `json:"dswrf"`
			Skycon              []HourlyDataPoint[string]      `json:"skycon"`
			AirQuality          struct {
				AQI  []HourlyDataPoint[AQIValue] `json:"aqi"`
				PM25 []HourlyDataPoint[float64]  `json:"pm25"`
			} `json:"air_quality"`
		} `json:"hourly"`
		Daily struct {
			Status              string                   `json:"status"`
			Astro               []DailyAstro             `json:"astro"`
			Precipitation08h20h []DailyPrecipitationData `json:"precipitation_08h_20h"`
			Precipitation20h32h []DailyPrecipitationData `json:"precipitation_20h_32h"`
			Precipitation       []DailyPrecipitationData `json:"precipitation"`
			Temperature         []DailyRangeData         `json:"temperature"`
			Temperature08h20h   []DailyRangeData         `json:"temperature_08h_20h"`
			Temperature20h32h   []DailyRangeData         `json:"temperature_20h_32h"`
			Wind                []DailyWindData          `json:"wind"`
			Wind08h20h          []DailyWindData          `json:"wind_08h_20h"`
			Wind20h32h          []DailyWindData          `json:"wind_20h_32h"`
			Humidity            []DailyRangeData         `json:"humidity"`
			Cloudrate           []DailyRangeData         `json:"cloudrate"`
			Pressure            []DailyRangeData         `json:"pressure"`
			Visibility          []DailyRangeData         `json:"visibility"`
			Dswrf               []DailyRangeData         `json:"dswrf"`
			AirQuality          struct {
				AQI  []DailyAQIData   `json:"aqi"`
				PM25 []DailyRangeData `json:"pm25"`
			} `json:"air_quality"`
			Skycon       []DailyStringIndex                   `json:"skycon"`
			Skycon08h20h []DailyStringIndex                   `json:"skycon_08h_20h"`
			Skycon20h32h []DailyStringIndex                   `json:"skycon_20h_32h"`
			LifeIndex    map[string][]LifeIndexDailyValueType `json:"life_index"`
		} `json:"daily"`
		Primary          int    `json:"primary"`
		ForecastKeypoint string `json:"forecast_keypoint"`
	} `json:"result"`
}

// LightWeatherResponse represents essential weather information
type LightWeatherResponse struct {
	Location    LocationInfo    `json:"location"`
	Alerts      []WeatherAlert  `json:"alerts"`
	Current     CurrentWeather  `json:"current"`
	Hourly      []HourlyWeather `json:"hourly"`
	Daily       []DailyWeather  `json:"daily"`
	Summary     WeatherSummary  `json:"summary"`
	LastUpdated time.Time       `json:"last_updated"`
}

// LocationInfo represents location details
type LocationInfo struct {
	Coordinates []float64 `json:"coordinates"`
	Region      string    `json:"region"`
	City        string    `json:"city"`
	Timezone    string    `json:"timezone"`
}

// WeatherAlert represents weather warnings and alerts
type WeatherAlert struct {
	Title       string    `json:"title"`
	Level       string    `json:"level"` // e.g., "蓝色预警", "Ⅳ级/一般"
	Description string    `json:"description"`
	Location    string    `json:"location"`
	PublishedAt time.Time `json:"published_at"`
	Source      string    `json:"source"`
}

// CurrentWeather represents current weather conditions
type CurrentWeather struct {
	Temperature         float64           `json:"temperature"`
	ApparentTemperature float64           `json:"apparent_temperature"`
	Condition           string            `json:"condition"`      // skycon
	ConditionText       string            `json:"condition_text"` // human readable
	Humidity            float64           `json:"humidity"`
	Wind                WindInfo          `json:"wind"`
	Pressure            float64           `json:"pressure"`
	Visibility          float64           `json:"visibility"`
	Precipitation       PrecipitationInfo `json:"precipitation"`
	AirQuality          AirQualityInfo    `json:"air_quality"`
	LifeIndices         map[string]string `json:"life_indices"` // index desc only
}

// HourlyWeather represents hourly forecast
type HourlyWeather struct {
	Time                time.Time `json:"time"`
	Temperature         float64   `json:"temperature"`
	ApparentTemperature float64   `json:"apparent_temperature"`
	Condition           string    `json:"condition"`
	PrecipitationMM     float64   `json:"precipitation_mm"`
	PrecipitationProb   int       `json:"precipitation_probability"`
	WindSpeed           float64   `json:"wind_speed"`
	Humidity            float64   `json:"humidity"`
}

// DailyWeather represents daily forecast
type DailyWeather struct {
	Date              time.Time         `json:"date"`
	TemperatureMin    float64           `json:"temperature_min"`
	TemperatureMax    float64           `json:"temperature_max"`
	Condition         string            `json:"condition"`
	ConditionDay      string            `json:"condition_day"`
	ConditionNight    string            `json:"condition_night"`
	PrecipitationMM   float64           `json:"precipitation_mm"`
	PrecipitationProb int               `json:"precipitation_probability"`
	Wind              WindInfo          `json:"wind"`
	Sunrise           string            `json:"sunrise"`
	Sunset            string            `json:"sunset"`
	AirQuality        AirQualityInfo    `json:"air_quality"`
	LifeIndices       map[string]string `json:"life_indices"`
}

// WindInfo represents wind information
type WindInfo struct {
	Speed     float64 `json:"speed"`
	Direction float64 `json:"direction"`
	Level     string  `json:"level"` // wind level description
}

// PrecipitationInfo represents precipitation details
type PrecipitationInfo struct {
	Intensity float64 `json:"intensity"`
	Status    string  `json:"status"`
	Nearby    bool    `json:"nearby"`
}

// AirQualityInfo represents air quality information
type AirQualityInfo struct {
	AQI         int    `json:"aqi"`
	Level       string `json:"level"` // e.g., "优", "良"
	PM25        int    `json:"pm25"`
	PrimaryPoll string `json:"primary_pollutant"`
}

// WeatherSummary represents natural language summaries
type WeatherSummary struct {
	Current  string `json:"current"`  // current weather description
	Hourly   string `json:"hourly"`   // hourly forecast description
	Forecast string `json:"forecast"` // forecast keypoint
}

// ConvertToLightModel converts the full API response to a lightweight model
func ConvertToLightModel(full *CaiyunAPIResponse) *LightWeatherResponse {
	if full == nil || full.Status != "ok" {
		return nil
	}

	light := &LightWeatherResponse{
		LastUpdated: time.Unix(full.ServerTime, 0),
		Location: LocationInfo{
			Coordinates: full.Location,
			Timezone:    full.Timezone,
		},
		Summary: WeatherSummary{
			Forecast: full.Result.ForecastKeypoint,
		},
	}

	// Extract location info from alerts if available
	if len(full.Result.Alert.Content) > 0 {
		alert := full.Result.Alert.Content[0]
		light.Location.Region = alert.Province
		light.Location.City = alert.City
	}

	// Convert alerts
	for _, alert := range full.Result.Alert.Content {
		light.Alerts = append(light.Alerts, WeatherAlert{
			Title:       alert.Title,
			Level:       extractAlertLevel(alert.Title),
			Description: alert.Description,
			Location:    alert.Location,
			PublishedAt: time.Unix(alert.Pubtimestamp, 0),
			Source:      alert.Source,
		})
	}

	// Convert current weather
	rt := full.Result.Realtime
	light.Current = CurrentWeather{
		Temperature:         rt.Temperature,
		ApparentTemperature: rt.ApparentTemperature,
		Condition:           rt.Skycon,
		ConditionText:       translateSkycon(rt.Skycon),
		Humidity:            rt.Humidity,
		Wind: WindInfo{
			Speed:     rt.Wind.Speed,
			Direction: rt.Wind.Direction,
			Level:     getWindLevel(rt.Wind.Speed),
		},
		Pressure:   rt.Pressure,
		Visibility: rt.Visibility,
		Precipitation: PrecipitationInfo{
			Intensity: rt.Precipitation.Local.Intensity,
			Status:    rt.Precipitation.Local.Status,
			Nearby:    rt.Precipitation.Nearest.Distance < 5.0, // within 5km
		},
		AirQuality: AirQualityInfo{
			AQI:         int(rt.AirQuality.AQI.CHN),
			Level:       rt.AirQuality.Description.CHN,
			PM25:        int(rt.AirQuality.PM25),
			PrimaryPoll: getPrimaryPollutant(rt.AirQuality),
		},
		LifeIndices: convertLifeIndices(rt.LifeIndex),
	}

	// Convert hourly data (next 24 hours)
	hourly := full.Result.Hourly
	light.Summary.Hourly = hourly.Description

	maxHours := min(24, len(hourly.Temperature))
	for i := 0; i < maxHours; i++ {
		t, _ := time.Parse(time.RFC3339, hourly.Temperature[i].Datetime)

		precipProb := 0
		precipMM := 0.0
		if i < len(hourly.Precipitation) {
			precipProb = hourly.Precipitation[i].Probability
			precipMM = hourly.Precipitation[i].Value
		}

		windSpeed := 0.0
		if i < len(hourly.Wind) {
			windSpeed = hourly.Wind[i].Speed
		}

		humidity := 0.0
		if i < len(hourly.Humidity) {
			humidity = hourly.Humidity[i].Value
		}

		apparentTemp := hourly.Temperature[i].Value
		if i < len(hourly.ApparentTemperature) {
			apparentTemp = hourly.ApparentTemperature[i].Value
		}

		condition := ""
		if i < len(hourly.Skycon) {
			condition = hourly.Skycon[i].Value
		}

		light.Hourly = append(light.Hourly, HourlyWeather{
			Time:                t,
			Temperature:         hourly.Temperature[i].Value,
			ApparentTemperature: apparentTemp,
			Condition:           condition,
			PrecipitationMM:     precipMM,
			PrecipitationProb:   precipProb,
			WindSpeed:           windSpeed,
			Humidity:            humidity,
		})
	}

	// Convert daily data
	daily := full.Result.Daily
	maxDays := min(7, len(daily.Temperature))
	for i := 0; i < maxDays; i++ {
		date, _ := time.Parse("2006-01-02T15:04-07:00", daily.Temperature[i].Date)

		precipProb := 0
		precipMM := 0.0
		if i < len(daily.Precipitation) {
			precipProb = daily.Precipitation[i].Probability
			precipMM = daily.Precipitation[i].Max
		}

		wind := WindInfo{}
		if i < len(daily.Wind) {
			wind = WindInfo{
				Speed:     daily.Wind[i].Max.Speed,
				Direction: daily.Wind[i].Max.Direction,
				Level:     getWindLevel(daily.Wind[i].Max.Speed),
			}
		}

		sunrise, sunset := "", ""
		if i < len(daily.Astro) {
			sunrise = daily.Astro[i].Sunrise.Time
			sunset = daily.Astro[i].Sunset.Time
		}

		condition := ""
		conditionDay := ""
		conditionNight := ""
		if i < len(daily.Skycon) {
			condition = daily.Skycon[i].Value
		}
		if i < len(daily.Skycon08h20h) {
			conditionDay = daily.Skycon08h20h[i].Value
		}
		if i < len(daily.Skycon20h32h) {
			conditionNight = daily.Skycon20h32h[i].Value
		}

		aqi := AirQualityInfo{}
		if i < len(daily.AirQuality.AQI) {
			aqi = AirQualityInfo{
				AQI:   int(daily.AirQuality.AQI[i].Avg.CHN),
				Level: getAQILevel(int(daily.AirQuality.AQI[i].Avg.CHN)),
			}
		}

		light.Daily = append(light.Daily, DailyWeather{
			Date:              date,
			TemperatureMin:    daily.Temperature[i].Min,
			TemperatureMax:    daily.Temperature[i].Max,
			Condition:         condition,
			ConditionDay:      conditionDay,
			ConditionNight:    conditionNight,
			PrecipitationMM:   precipMM,
			PrecipitationProb: precipProb,
			Wind:              wind,
			Sunrise:           sunrise,
			Sunset:            sunset,
			AirQuality:        aqi,
			LifeIndices:       convertDailyLifeIndices(daily.LifeIndex, i),
		})
	}

	return light
}

// Helper functions
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func extractAlertLevel(title string) string {
	// Extract alert level from title like "台风蓝色预警[Ⅳ级/一般]"
	if len(title) > 0 {
		// Simple extraction - could be improved with regex
		if contains := func(s, substr string) bool {
			return len(s) >= len(substr) && s[len(s)-len(substr):] == substr
		}; contains(title, "]") {
			start := -1
			for i := len(title) - 1; i >= 0; i-- {
				if title[i] == '[' {
					start = i
					break
				}
			}
			if start != -1 {
				return title[start+1 : len(title)-1]
			}
		}
	}
	return "一般"
}

func translateSkycon(skycon string) string {
	translations := map[string]string{
		"CLEAR_DAY":           "晴天",
		"CLEAR_NIGHT":         "晴夜",
		"PARTLY_CLOUDY_DAY":   "多云",
		"PARTLY_CLOUDY_NIGHT": "多云",
		"CLOUDY":              "阴天",
		"LIGHT_RAIN":          "小雨",
		"MODERATE_RAIN":       "中雨",
		"HEAVY_RAIN":          "大雨",
		"STORM_RAIN":          "暴雨",
		"WIND":                "大风",
	}
	if text, ok := translations[skycon]; ok {
		return text
	}
	return skycon
}

func getWindLevel(speed float64) string {
	// Convert m/s to wind level
	if speed < 0.3 {
		return "无风"
	} else if speed < 1.6 {
		return "软风"
	} else if speed < 3.4 {
		return "轻风"
	} else if speed < 5.5 {
		return "微风"
	} else if speed < 8.0 {
		return "和风"
	} else if speed < 10.8 {
		return "清风"
	} else if speed < 13.9 {
		return "强风"
	} else if speed < 17.2 {
		return "疾风"
	} else {
		return "大风"
	}
}

func getPrimaryPollutant(aq any) string {
	// This would need actual implementation based on pollutant levels
	return "PM2.5" // placeholder
}

func getAQILevel(aqi int) string {
	if aqi <= 50 {
		return "优"
	} else if aqi <= 100 {
		return "良"
	} else if aqi <= 150 {
		return "轻度污染"
	} else if aqi <= 200 {
		return "中度污染"
	} else if aqi <= 300 {
		return "重度污染"
	}
	return "严重污染"
}

func convertLifeIndices(indices map[string]LifeIndexValueType) map[string]string {
	result := make(map[string]string)
	for key, value := range indices {
		result[key] = value.Desc
	}
	return result
}

func convertDailyLifeIndices(indices map[string][]LifeIndexDailyValueType, dayIndex int) map[string]string {
	result := make(map[string]string)
	for key, values := range indices {
		if dayIndex < len(values) {
			result[key] = values[dayIndex].Desc
		}
	}
	return result
}
