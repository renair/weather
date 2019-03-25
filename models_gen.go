// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package weather

type Cloud struct {
	Cloudiness float64 `json:"cloudiness"`
	UvIndex    *int    `json:"uvIndex"`
	IsSnow     *bool   `json:"isSnow"`
	IsRain     *bool   `json:"isRain"`
}

type Location struct {
	ID           int     `json:"id"`
	Longitude    float64 `json:"longitude"`
	Latitude     float64 `json:"latitude"`
	LocationName *string `json:"locationName"`
}

type MainValues struct {
	Temperature float64 `json:"temperature"`
	Humidity    float64 `json:"humidity"`
	Brightness  int     `json:"brightness"`
	Pressure    *int    `json:"pressure"`
}

type NewLocation struct {
	Longitude    float64 `json:"longitude"`
	Latitude     float64 `json:"latitude"`
	LocationName *string `json:"locationName"`
}

type WeatherData struct {
	Location Location   `json:"location"`
	Values   MainValues `json:"values"`
	Cloud    Cloud      `json:"cloud"`
	Wind     *Wind      `json:"wind"`
	Date     string     `json:"date"`
}

type Wind struct {
	Speed     float64 `json:"speed"`
	Direction float64 `json:"direction"`
}