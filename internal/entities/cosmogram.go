package entities

type CosmogramRequestBody struct {
	Candidate User   `json:"candidate"`
	Staff     []User `json:"staff"`
}

type Subject struct {
	Name       string  `json:"name"`
	Year       int     `json:"year"`
	Month      int     `json:"month"`
	Day        int     `json:"day"`
	Hour       int     `json:"hour"`
	Minute     int     `json:"minute"`
	Longitude  float64 `json:"longitude"`
	Latitude   float64 `json:"latitude"`
	City       string  `json:"city"`
	Nation     string  `json:"nation"`
	Timezone   string  `json:"timezone"`
	ZodiacType string  `json:"zodiac_type"`
}

type CosmogramApiRequestBody struct {
	FirstSubject  Subject `json:"first_subject"`
	SecondSubject Subject `json:"second_subject"`
}

type CosmogramApiResponseBody struct {
	Status           string   `json:"status"`
	Score            int      `json:"score"`
	ScoreDescription string   `json:"score_description"`
	IsDestinySign    bool     `json:"is_destiny_sign"`
	Aspects          []Aspect `json:"aspects"`
	Data             Data     `json:"data"`
}

// Структура для аспектов
type Aspect struct {
	P1Name string  `json:"p1_name"`
	P2Name string  `json:"p2_name"`
	Aspect string  `json:"aspect"`
	Orbit  float64 `json:"orbit"`
}

// Структура для данных
type Data struct {
	FirstSubject  Subject2 `json:"first_subject"`
	SecondSubject Subject2 `json:"second_subject"`
}

// Структура для субъекта (например, первого и второго)
type Subject2 struct {
	Name                      string  `json:"name"`
	Year                      int     `json:"year"`
	Month                     int     `json:"month"`
	Day                       int     `json:"day"`
	Hour                      int     `json:"hour"`
	Minute                    int     `json:"minute"`
	City                      string  `json:"city"`
	Nation                    string  `json:"nation"`
	Lng                       float64 `json:"lng"`
	Lat                       float64 `json:"lat"`
	TZStr                     string  `json:"tz_str"`
	ZodiacType                string  `json:"zodiac_type"`
	SiderealMode              *string `json:"sidereal_mode"`
	HousesSystemIdentifier    string  `json:"houses_system_identifier"`
	HousesSystemName          string  `json:"houses_system_name"`
	PerspectiveType           string  `json:"perspective_type"`
	IsoFormattedLocalDatetime string  `json:"iso_formatted_local_datetime"`
	IsoFormattedUTCDateTime   string  `json:"iso_formatted_utc_datetime"`
	JulianDay                 float64 `json:"julian_day"`
	UTCtime                   float64 `json:"utc_time"`
	LocalTime                 float64 `json:"local_time"`
	Sun                       Planet  `json:"sun"`
	Moon                      Planet  `json:"moon"`
	Mercury                   Planet  `json:"mercury"`
	Venus                     Planet  `json:"venus"`
	Mars                      Planet  `json:"mars"`
	Jupiter                   Planet  `json:"jupiter"`
	Saturn                    Planet  `json:"saturn"`
	Uranus                    Planet  `json:"uranus"`
	Neptune                   Planet  `json:"neptune"`
	Pluto                     Planet  `json:"pluto"`
	Chiron                    Planet  `json:"chiron"`
	MeanLilith                Planet  `json:"mean_lilith"`
	MeanNode                  Planet  `json:"mean_node"`
	MeanSouthNode             Planet  `json:"mean_south_node"`
}

// Структура для планеты
type Planet struct {
	Name       string  `json:"name"`
	Quality    string  `json:"quality"`
	Element    string  `json:"element"`
	Sign       string  `json:"sign"`
	SignNum    int     `json:"sign_num"`
	Position   float64 `json:"position"`
	AbsPos     float64 `json:"abs_pos"`
	Emoji      string  `json:"emoji"`
	PointType  string  `json:"point_type"`
	House      string  `json:"house"`
	Retrograde bool    `json:"retrograde"`
}
