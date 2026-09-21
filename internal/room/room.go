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
func (r *Room) DoesClientExistInRoom(client *client.Client)bool{
	_, exists := r.Clients[client]
	if exists {
		return true
	}
	return  false
}

func (r *Room) BroadcastMessage(msg string){
	for client := range r.Clients{
		client.SendMessage(msg)
	}
}

func (r *Room) BroadcastMessageExceptClient(notclient *client.Client, msg string){
	for client := range r.Clients{
		if client != notclient {
			client.SendMessage(msg)
		}
		
	}
}




