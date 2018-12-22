package api

import fmt    "fmt"
import json   "encoding/json"

type GetRoomMember struct {
	AccountID        int    `json:"account_id"`
	Role             string `json:"role"`
	Name             string `json:"name"`
	ChatworkID       string `json:"chatwork_id"`
	OrganizationID   int    `json:"organization_id"`
	OrganizationName string `json:"organization_name"`
	Department       string `json:"department"`
	AvatarImageURL   string `json:"avatar_image_url"`
}

type GetRoomMemberService struct {
	token               string
	roomId              string
}

func (rms *GetRoomMemberService) SetRoomId(roomId string) *GetRoomMemberService {
	rms.roomId = roomId
	return rms
}

func (rms *GetRoomMemberService) BuildRequestURL() (string, error) {
	u, err := buildAPIEndpoint("/rooms/" + rms.roomId + "/members")
	if err != nil {
		return "", err
	}

	return u.String(), nil
}

func NewGetRoomMemberService(token string) *GetRoomMemberService {
	return &GetRoomMemberService{
		token:               token,
	}
}

func (rts *GetRoomMemberService) Execute() []GetRoomMember{
	reqUrl, err := rts.BuildRequestURL()
	result, err := RequestJSON(reqUrl, rts.token)
	var members []GetRoomMember
	err = json.Unmarshal([]byte(result), &members)
	if err != nil {
		fmt.Println("Error at API request:%#v", err)
	}

	return members
}