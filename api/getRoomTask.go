package api

import url    "net/url"
import fmt    "fmt"
import json   "encoding/json"

type GetRoomTask struct {
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

type GetRoomTaskService struct {
	token               string
	roomId              string
	accountId           string
	assignedByAccountId string
	status              string
}

func (rts *GetRoomTaskService) SetRoomId(roomId string) *GetRoomTaskService {
	rts.roomId = roomId
	return rts
}

func (rts *GetRoomTaskService) SetAccountId(accountId string) *GetRoomTaskService {
	rts.accountId = accountId
	return rts
}

func (rts *GetRoomTaskService) SetAssignedByAccountId(assignedByAccountId string) *GetRoomTaskService {
	rts.assignedByAccountId = assignedByAccountId
	return rts
}

func (rts *GetRoomTaskService) SetStatus(status string) *GetRoomTaskService {
	rts.status = status
	return rts
}

func (rts *GetRoomTaskService) BuildRequestURL() (*url.URL, error) {
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
		return nil, err
	}
	u.RawQuery = queries.Encode()

	return u, nil
}

func NewGetRoomTaskService(token string) *GetRoomTaskService {
	return &GetRoomTaskService{
		token:               token,
		roomId:              "",
		accountId:           "",
		assignedByAccountId: "",
		status:              "",
	}
}

func (rts *GetRoomTaskService) Execute() []GetRoomTask{
	u, err := rts.BuildRequestURL()
	result, err := RequestJSON(u, "GET", rts.token)
	var tasks []GetRoomTask
	err = json.Unmarshal([]byte(result), &tasks)
	if err != nil {
		fmt.Println("Error at API request:%#v", err)
	}

	return tasks
}