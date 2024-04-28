package dto

type EventRequest struct {
	Name      string `json:"name" form:"name"`
	Date      string `json:"date" form:"date"`
	Detail    string `json:"detail" form:"detail"`
	Organizer string `json:"organizer" form:"organizer"`
	Image     string `json:"image" form:"image"`
}

type EventResponse struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Date      string `json:"date"`
	Detail    string `json:"detail"`
	Organizer string `json:"organizer"`
	Image     string `json:"image"`
}

type EventUpdateRequest struct {
	Name      string `json:"name" form:"name"`
	Date      string `json:"date" form:"date"`
	Detail    string `json:"detail" form:"detail"`
	Organizer string `json:"organizer" form:"organizer"`
	Image     string `json:"image" form:"image"`
}
