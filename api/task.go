package api

import url    "net/url"
import fmt    "fmt"
import json   "encoding/json"

type Task struct {
	TaskID  int `json:"task_id"`
	Account struct {
		AccountID      int    `json:"account_id"`
		Name           string `json:"name"`
		AvatarImageURL string `json:"avatar_image_url"`
	} `json:"account"`
	AssignedByAccount struct {
		AccountID      int    `json:"account_id"`
		Name           string `json:"name"`
		AvatarImageURL string `json:"avatar_image_url"`
	} `json:"assigned_by_account"`
	MessageID string `json:"message_id"`
	Body      string `json:"body"`
	LimitTime int    `json:"limit_time"`
	Status    string `json:"status"`
}

type RoomTaskService struct {
	token               string
	roomId              string
	accountId           string
	assignedByAccountId string
	status              string
}

func (rts *RoomTaskService) SetRoomId(roomId string) *RoomTaskService {
	rts.roomId = roomId
	return rts
}

func (rts *RoomTaskService) SetAccountId(accountId string) *RoomTaskService {
	rts.accountId = accountId
	return rts
}

func (rts *RoomTaskService) SetAssignedByAccountId(assignedByAccountId string) *RoomTaskService {
	rts.assignedByAccountId = assignedByAccountId
	return rts
}

func (rts *RoomTaskService) SetStatus(status string) *RoomTaskService {
	rts.status = status
	return rts
}

func (rts *RoomTaskService) BuildRequestURL() (string, error) {
	queries := url.Values{}

	if (rts.accountId != ""){
		queries.Add("account_id", rts.accountId)
	}

	if (rts.assignedByAccountId != ""){
		queries.Add("assigned_by_account_id", rts.assignedByAccountId)
	}

	queries.Add("status", rts.status)
	u, err := buildAPIEndpoint("/rooms/" + rts.roomId + "/tasks")
	if err != nil {
		return "", err
	}
	u.RawQuery = queries.Encode()

	return u.String(), nil
}

func NewRoomTaskService(token string) *RoomTaskService {
	return &RoomTaskService{
		token:               token,
		roomId:              "",
		accountId:           "",
		assignedByAccountId: "",
		status:              "",
	}
}

func (rts *RoomTaskService) Execute() []Task{
	reqUrl, err := rts.BuildRequestURL()
	result, err := RequestJSON(reqUrl, rts.token)
	var tasks []Task
	err = json.Unmarshal([]byte(result), &tasks)
	if err != nil {
		fmt.Println("Error at API request:%#v", err)
	}

	return tasks
}