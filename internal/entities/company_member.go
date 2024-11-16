package entities

type CompanyMember struct {
	UserID    int `json:"user_id"`
	CompanyID int `json:"company_id"`
	Position  int `json:"position"`
}

type GetCompanyMembersInfoRequest struct {
	CompanyID int `json:"company_id"`
}

type CompanyMemberInfo struct {
	UserID    int    `json:"user_id"`
	Name      string `json:"name"`
	Surname   string `json:"surname"`
	ThirdName string `json:"third_name"`
	Position  int    `json:"position"`
}
