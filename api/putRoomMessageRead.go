package api

import url    "net/url"
import fmt    "fmt"
import json   "encoding/json"

type PutRoomMessageRead struct {
	UnreadNum  int `json:"unread_num"`
	MentionNum int `json:"mention_num"`
}

type PutRoomMessageReadService struct {
	token               string
	roomId              string
	messageId           string
	action              string //read or unread
}

func (prm *PutRoomMessageReadService) SetRoomId(roomId string) *PutRoomMessageReadService {
	prm.roomId = roomId
	return prm
}

func (prm *PutRoomMessageReadService) SetAction(action string) *PutRoomMessageReadService {
	prm.action = action
	return prm
}

func (prm *PutRoomMessageReadService) SetMessageId(messageId string) *PutRoomMessageReadService {
	prm.messageId = messageId
	return prm
}

func (prm *PutRoomMessageReadService) BuildRequestURL() (*url.URL, error) {
	queries := url.Values{}

	queries.Add("message_id", prm.messageId)

	u, err := buildAPIEndpoint("/rooms/" + prm.roomId + "/messages/" + prm.action)
	if err != nil {
		return nil, err
	}
	u.RawQuery = queries.Encode()

	return u, nil
}

func NewPutRoomMessageReadService(token string) *PutRoomMessageReadService {
	return &PutRoomMessageReadService{
		token:              token,
		roomId:             "",
		messageId:          "",
		action:             "", //read or unread
	}
}

func (prm *PutRoomMessageReadService) Execute() PutRoomMessageRead{
	u, err := prm.BuildRequestURL()
	result, err := RequestJSON(u, "PUT", prm.token)
	var messageRead PutRoomMessageRead
	err = json.Unmarshal([]byte(result), &messageRead)
	if err != nil {
		fmt.Println("Error at API request:%#v", err)
	}

	return messageRead
}