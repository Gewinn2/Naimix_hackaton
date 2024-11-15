package entities

type User struct {
	ID          int    `json:"id" example:"1"`
	Name        string `json:"name" example:"John"`
	ThirdName   string `json:"third_name" example:"Petrovich"`
	Surname     string `json:"surname" example:"Doe"`
	Email       string `json:"email" example:"john.doe@example.com"`
	Password    string `json:"password" example:"P@ssw0rd123"`
	PhoneNumber string `json:"phone_number" example:"+1234567890"`
	BirthDate   string `json:"birth_date" example:"1990-01-01"`
	Role        string `json:"role" example:"user"`
}

type CreateUserRequest struct {
	Name        string `json:"name" example:"John"`
	ThirdName   string `json:"third_name" example:"Middle"`
	Surname     string `json:"surname" example:"Doe"`
	Email       string `json:"email" example:"john.doe@example.com"`
	Password    string `json:"password" example:"P@ssw0rd123"`
	PhoneNumber string `json:"phone_number" example:"+1234567890"`
	BirthDate   string `json:"birth_date" example:"1990-01-01"`
	Role        string `json:"role" example:"HR manager"`
}

type CreateUserResponse struct {
	AccessToken  string `json:"access_token" example:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJNYXBDbGFpbXMiOnsiRXhwaXJlc0F0IjoxNzMxNzE1ODM0LCJJc3N1ZWRBciI6MTczMTY5NDIzNH0sInVzZXJfaWQiOjJ9.4PZY_8iDHCeACmPs5woqtowvCrrByTB4Gadr6oLdlZg"`
	RefreshToken string `json:"refresh_token" example:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJNYXBDbGFpbXMiOnsiRXhwaXJlc0F0IjoxNzM0MzAwNjM0LCJJc3N1ZWRBciI6MTczMTY5NDIzNH0sInVzZXJfaWQiOjJ9.o9i-1GBcLtUR0SpDNiBuHfzQWHBt465bhxcC_X2WqBY"`
	ID           int    `json:"id" example:"1"`
	Email        string `json:"email" example:"john.doe@example.com"`
}

type LoginUserRequest struct {
	PhoneNumber string `json:"phone_number" example:"+1234567890"`
	Password    string `json:"password" example:"P@ssw0rd123"`
	Email       string `json:"email" example:"john.doe@example.com"`
}

type LoginUserResponse struct {
	AccessToken  string `json:"access_token" example:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJNYXBDbGFpbXMiOnsiRXhwaXJlc0F0IjoxNzMxNzE1ODM0LCJJc3N1ZWRBciI6MTczMTY5NDIzNH0sInVzZXJfaWQiOjJ9.4PZY_8iDHCeACmPs5woqtowvCrrByTB4Gadr6oLdlZg"`
	RefreshToken string `json:"refresh_token" example:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJNYXBDbGFpbXMiOnsiRXhwaXJlc0F0IjoxNzM0MzAwNjM0LCJJc3N1ZWRBciI6MTczMTY5NDIzNH0sInVzZXJfaWQiOjJ9.o9i-1GBcLtUR0SpDNiBuHfzQWHBt465bhxcC_X2WqBY"`
	ID           int    `json:"id" example:"1"`
	Email        string `json:"email" example:"john.doe@example.com"`
}

type ErrorResponse struct {
	Error string `json:"error" example:"Invalid data provided"`
}
