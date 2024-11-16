package entities

type EmployeeRequest struct {
	Name      string `json:"name" example:"John"`
	ThirdName string `json:"third_name" example:"Petrovich"`
	Surname   string `json:"surname" example:"Doe"`
	BirthDate string `json:"birth_date" example:"1990-01-01"`
}

type CosmogramRequestBody struct {
	Candidate EmployeeRequest   `json:"candidate"`
	Staff     []EmployeeRequest `json:"staff"`
}

type CosmogramResponseBody struct {
	Compatibilities []Compatibility `json:"compatibilities"`
	Cosmogram       Subject2        `json:"cosmogram"`
}

type Compatibility struct {
	CandidateFullName    string `json:"candidate_full_name" example:"Doe John Petrovich"`
	EmployeeFullName     string `json:"employee_full_name" example:"Doe John Petrovich"`
	Communication        int    `json:"communication"`
	CommunicationComment string `json:"communication_comment"`
	Emotions             int    `json:"emotions"`
	EmotionsComment      string `json:"emotions_comment"`
	Work                 int    `json:"work"`
	WorkComment          string `json:"work_comment"`
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
	Data Data `json:"data"`
}

// Структура для данных
type Data struct {
	FirstSubject  Subject2 `json:"first_subject"`
	SecondSubject Subject2 `json:"second_subject"`
}

// Структура для субъекта (например, первого и второго)
type Subject2 struct {
	Name          string  `json:"name"`
	Year          int     `json:"year"`
	Month         int     `json:"month"`
	Day           int     `json:"day"`
	Hour          int     `json:"hour"`
	Minute        int     `json:"minute"`
	City          string  `json:"city"`
	Nation        string  `json:"nation"`
	Lng           float64 `json:"lng"`
	Lat           float64 `json:"lat"`
	Sun           Planet  `json:"sun"`
	Moon          Planet  `json:"moon"`
	Mercury       Planet  `json:"mercury"`
	Venus         Planet  `json:"venus"`
	Mars          Planet  `json:"mars"`
	Jupiter       Planet  `json:"jupiter"`
	Saturn        Planet  `json:"saturn"`
	Uranus        Planet  `json:"uranus"`
	Neptune       Planet  `json:"neptune"`
	Pluto         Planet  `json:"pluto"`
	Chiron        Planet  `json:"chiron"`
	MeanLilith    Planet  `json:"mean_lilith"`
	MeanNode      Planet  `json:"mean_node"`
	MeanSouthNode Planet  `json:"mean_south_node"`
}

// Структура для планеты
type Planet struct {
	Name       string  `json:"name"`
	AbsPos     float64 `json:"abs_pos"`
	Retrograde bool    `json:"retrograde"`
}
