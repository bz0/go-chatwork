package chatwork

import api "./api"

const (
	Domain  = "https://api.chatwork.com"
	Version = "v2"
)

type Client struct {
	// services
	RoomTask *api.RoomTaskService
}

func New(token string) *Client {
	c := &Client{
	}

	c.RoomTask = api.NewRoomTaskService(token)
	return c
}