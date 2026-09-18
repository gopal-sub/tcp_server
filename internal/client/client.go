package client

import (
	"net"
)


type Client struct {
	conn net.Conn
	room string
}


func CreateClent(conn net.Conn) *Client{
	return &Client{
		conn: conn,
	}
}


func (c *Client) ChangeRoomForClient(room string) {
	c.room = room
}

