package client

import (
	"net"
)


type Client struct {
	Conn net.Conn
	Room string
}


func CreateClent(conn net.Conn) *Client{
	return &Client{
		Conn: conn,
	}
}


func (c *Client) ChangeRoomForClient(room string) {
	c.Room = room
}


func (c *Client) SendMessage(msg string) {
	c.Conn.Write([]byte(msg + "\n"))
}

