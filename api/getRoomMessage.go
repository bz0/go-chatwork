package api

import url    "net/url"
import fmt    "fmt"
import json   "encoding/json"

type GetRoomMessage struct {
	MessageID string `json:"message_id"`
	Account   struct {
		AccountID      int    `json:"account_id"`
		Name           string `json:"name"`
		AvatarImageURL string `json:"avatar_image_url"`
	} `json:"account"`
	Body       string `json:"body"`
	SendTime   int    `json:"send_time"`
	UpdateTime int    `json:"update_time"`
}

type GetRoomMessageService struct {
	token               string
	roomId              string
	force               string
}

func (rms *GetRoomMessageService) SetRoomId(roomId string) *GetRoomMessageService {
	rms.roomId = roomId
	return rms
}

func (rms *GetRoomMessageService) SetForce(force string) *GetRoomMessageService {
	rms.force = force
	return rms
}

func (rms *GetRoomMessageService) BuildRequestURL() (*url.URL, error) {
	queries := url.Values{}

	if (rms.force != ""){
		queries.Add("force", rms.force)
	}

	u, err := buildAPIEndpoint("/rooms/" + rms.roomId + "/messages")
	if err != nil {
		return nil, err
	}
	u.RawQuery = queries.Encode()

	return u, nil
}

func NewGetRoomMessageService(token string) *GetRoomMessageService {
	return &GetRoomMessageService{
		token:               token,
		roomId:              "",
		force:               "",
	}
}

func (rms *GetRoomMessageService) Execute() []GetRoomMessage{
	u, err := rms.BuildRequestURL()
	result, err := RequestJSON(u, "GET", rms.token)
	var messages []GetRoomMessage
	err = json.Unmarshal([]byte(result), &messages)
	if err != nil {
		fmt.Println("Error at API request:%#v", err)
	}

	return messages
}