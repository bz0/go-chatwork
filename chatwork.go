package chatwork

import api "./api"

const (
	Domain  = "https://api.chatwork.com"
	Version = "v2"
)

type Client struct {
	// services
	GetRoomTask *api.GetRoomTaskService
	GetRoomMember *api.GetRoomMemberService
	GetRoom     *api.GetRoomService
}

func New(token string) *Client {
	c := &Client{
	}

	c.GetRoomTask   = api.NewGetRoomTaskService(token)
	c.GetRoomMember = api.NewGetRoomMemberService(token)
	c.GetRoom       = api.NewGetRoomService(token)
	return c
}