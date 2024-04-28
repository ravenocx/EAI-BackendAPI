package dto

type ProductRequest struct {
	Product       string `json:"product" form:"product"`
	Description   string `json:"description" form:"description"`
	Category      string `json:"category" form:"category"`
	OnDevelopment string `json:"on_development" form:"on_development"`
	Image         string `json:"image" form:"image"`
}

type ProductResponse struct {
	ID            string `jsoin:"id"`
	Product       string `json:"product"`
	Description   string `json:"description"`
	Category      string `json:"category"`
	OnDevelopment string `json:"on_development"`
	Image         string `json:"image"`
}

type ProductUpdateRequest struct {
	Product       string `json:"product" form:"product"`
	Description   string `json:"description" form:"description"`
	Category      string `json:"category" form:"category"`
	OnDevelopment string `json:"on_development" form:"on_development"`
	Image         string `json:"image" form:"image"`
}
