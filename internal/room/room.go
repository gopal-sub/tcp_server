package room

import (
	"tcp-chat/internal/client"

)


type Room struct {
	Name string
	Clients map[*client.Client]struct{}
}

func CreateRoom(name string)*Room{
	return &Room{
		Name: name,
		Clients: make(map[*client.Client]struct{}),
	}
}


func (r *Room) AddClientToRoom(client *client.Client){
	r.Clients[client] = struct{}{}
}

func (r *Room) RemoveClientFromRoom(client *client.Client){
	delete(r.Clients, client)
}




