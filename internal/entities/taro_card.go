package entities

type TaroCard struct {
	ID             int    `json:"id"`
	DirectMeaning  string `json:"direct_meaning"`
	ReverseMeaning string `json:"reverse_meaning"`
}
