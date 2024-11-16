package entities

type Company struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type CreateCompanyRequest struct {
	Name string `json:"name"`
}

type CreateCompanyResponse struct {
	ID int `json:"id"`
}
