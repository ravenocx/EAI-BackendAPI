package controllers

import "github.com/ravenocx/EAI-BackendAPI/internal/services"

type MemberController struct {
	memberService services.MemberService
}

func NewMemberController(memberService services.MemberService) *MemberController{
	return &MemberController{memberService: memberService}
}