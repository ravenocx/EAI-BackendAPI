package dto

type MemberRequest struct {
	Name   string `json:"name" form:"name"`
	Role   string `json:"role" form:"role"` // divisi
	Status string `json:"status" form:"status"`
	Image  string `json:"image" form:"image"`
}

type MemberResponse struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Role   string `json:"role"`
	Status string `json:"status"`
	Image  string `json:"image"`
}
