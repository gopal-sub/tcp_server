package client

import (
	"errors"
	"net"
	"time"
)


type Client struct {
	Conn net.Conn
	LastActivity time.Time
	Active bool
	SentPing bool
	
}


func CreateClient(conn net.Conn) *Client{
	return &Client{
		Conn: conn,
		LastActivity: time.Now(),
		Active: true,
	}
}

func (c *Client) SendMessage(msg string) {
	msgBytes := []byte(msg + "\n")
	n, err := c.Conn.Write(msgBytes)
	if errors.Is(err, net.ErrClosed){
		return
	}
	if err != nil {
		msgBytes = msgBytes[n:]
		for range 3 {
			
			n, err := c.Conn.Write(msgBytes)
			if err == nil{
				break
			}
			msgBytes = msgBytes[n:]
		
		}

	}

	
}


func (c *Client) DeactivateClient(){
	// c.MU.Lock()
	// defer c.MU.Unlock()
	
	c.Active = false
}


func (c *Client) UpdateLastActivity(){
	c.LastActivity = time.Now()
}

