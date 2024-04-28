package dto

type AchievementRequest struct {
	MemberID    string `json:"member_id" form:"member_id"`
	Image       string `json:"image" form:"image"`
	Achievement string `json:"achievement" form:"achievement"`
}

type AchievementResponse struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Image       string `json:"image"`
	Achievement string `json:"achievement"`
}

type AchievementUpdateRequest struct {
	MemberID    string `json:"member_id" form:"member_id"`
	Image       string `json:"image" form:"image"`
	Achievement string `json:"achievement" form:"achievement"`
}
