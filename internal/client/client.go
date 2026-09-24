package client

import (
	"net"
	"time"
)


type Client struct {
	Conn net.Conn
	LastActivity time.Time
	
}


func CreateClent(conn net.Conn) *Client{
	return &Client{
		Conn: conn,
		LastActivity: time.Now(),
	}
}

func (c *Client) SendMessage(msg string) {
	c.Conn.Write([]byte(msg + "\n"))
}

