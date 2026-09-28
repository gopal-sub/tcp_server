package client

import (
	"net"
	"time"
)


type Client struct {
	Conn net.Conn
	LastActivity time.Time
	CloseByServer bool
	
}


func CreateClent(conn net.Conn) *Client{
	return &Client{
		Conn: conn,
		LastActivity: time.Now(),
		CloseByServer: false,
	}
}

func (c *Client) SendMessage(msg string) {
	c.Conn.Write([]byte(msg + "\n"))
}

func (c *Client) CloseClientByServer(){
	// c.MU.Lock()
	// defer c.MU.Unlock()
	
	c.CloseByServer = true
}


