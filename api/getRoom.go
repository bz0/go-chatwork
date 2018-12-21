package api

import fmt    "fmt"
import json   "encoding/json"

type GetRoom struct {
	RoomID         int    `json:"room_id"`
	Name           string `json:"name"`
	Type           string `json:"type"`
	Role           string `json:"role"`
	Sticky         bool   `json:"sticky"`
	UnreadNum      int    `json:"unread_num"`
	MentionNum     int    `json:"mention_num"`
	MytaskNum      int    `json:"mytask_num"`
	MessageNum     int    `json:"message_num"`
	FileNum        int    `json:"file_num"`
	TaskNum        int    `json:"task_num"`
	IconPath       string `json:"icon_path"`
	LastUpdateTime int    `json:"last_update_time"`
}

type GetRoomService struct {
	token               string
}

func (rts *GetRoomService) BuildRequestURL() (string, error) {
	u, err := buildAPIEndpoint("/rooms")
	if err != nil {
		return "", err
	}

	return u.String(), nil
}

func NewGetRoomService(token string) *GetRoomService {
	return &GetRoomService{
		token: token,
	}
}

func (rts *GetRoomService) Execute() []GetRoom{
	reqUrl, err := rts.BuildRequestURL()
	result, err := RequestJSON(reqUrl, rts.token)
	var room []GetRoom
	err = json.Unmarshal([]byte(result), &room)
	if err != nil {
		fmt.Println("Error at API request:%#v", err)
	}

	return room
}