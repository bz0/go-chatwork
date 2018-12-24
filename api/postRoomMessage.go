package api

import url    "net/url"
import fmt    "fmt"
import json   "encoding/json"

type PostRoomMessage struct {
	MessageId  string `json:"message_id"`
}

type PostRoomMessageService struct {
	token               string
	roomId              string
	body                string
	selfUnread          string
}

func (prm *PostRoomMessageService) SetRoomId(roomId string) *PostRoomMessageService {
	prm.roomId = roomId
	return prm
}

func (prm *PostRoomMessageService) SetBody(body string) *PostRoomMessageService {
	prm.body = body
	return prm
}

func (prm *PostRoomMessageService) SetSelfUnread(selfUnread string) *PostRoomMessageService {
	prm.selfUnread = selfUnread
	return prm
}

func (prm *PostRoomMessageService) BuildRequestURL() (*url.URL, error) {
	queries := url.Values{}

	queries.Add("body", prm.body)

	if (prm.selfUnread != ""){
		queries.Add("self_unread", prm.selfUnread)
	}

	u, err := buildAPIEndpoint("/rooms/" + prm.roomId + "/messages")
	if err != nil {
		return nil, err
	}
	u.RawQuery = queries.Encode()

	return u, nil
}

func NewPostRoomMessageService(token string) *PostRoomMessageService {
	return &PostRoomMessageService{
		token:               token,
		body:                "",
		selfUnread:          "",
	}
}

func (rts *PostRoomMessageService) Execute() PostRoomMessage{
	u, err := rts.BuildRequestURL()
	result, err := RequestJSON(u, "POST", rts.token)
	var messages PostRoomMessage
	err = json.Unmarshal([]byte(result), &messages)
	if err != nil {
		fmt.Println("Error at API request:%#v", err)
	}

	return messages
}