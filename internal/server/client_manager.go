package server

import (
	"fmt"
	"tcp-chat/internal/client"
	"time"
)


func (s *Server)CreateGlobalClientList(){
	s.Client = make(map[*client.Client]struct{})
}
func (s *Server)AddToGlobalClientList(client *client.Client){
	s.mu.Lock()

	defer s.mu.Unlock()
	s.Client[client] = struct{}{}
}

func (s *Server)RemoveClientByServer(client *client.Client){
	
	s.RemoveClientFromAllRooms(client)
	delete(s.Client, client)
	client.DeactivateClient()
	client.Conn.Close()
}


func (s *Server) TimeoutChecker(TimeoutFrequency time.Duration){


	ticker := time.NewTicker(TimeoutFrequency)
	for {
		<-ticker.C
		for c, _ := range s.Client{
			if time.Since(c.LastActivity) > TimeoutFrequency {
				fmt.Println("client removed")
				s.RemoveClientByServer(c)
			}
		}
	}
}

func (s *Server) PingClients(TimeoutFrequency time.Duration){


	freq := TimeoutFrequency /3

	ticker := time.NewTicker(freq)

	for {
		<-ticker.C
		s.mu.Lock()
		for c, _ := range s.Client{
			c.SendMessage("PING")
			c.SentPing = true
		}
		s.mu.Unlock()
	}

}