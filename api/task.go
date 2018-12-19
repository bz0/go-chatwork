package api

import fmt    "fmt"
import url    "net/url"

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
	u, err := buildAPIEndpoint("/" + Version + "/rooms/" + rts.roomId + "/tasks")
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

func (rts *RoomTaskService) ExecuteWeak() (result, error){
	reqURL, err := rts.BuildRequestURL()
	fmt.Println(reqURL)
	return RequestJSON(reqURL, rts.token)
}

func (rts *RoomTaskService) Execute() (result, error){
	result, err := rts.ExecuteWeak()
	if err != nil {
		return nil, err
	}

	return result, error
}

