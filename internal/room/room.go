package room

import (
	"sync"
	"tcp-chat/internal/client"
)


type Room struct {
	Name string
	Clients map[*client.Client]struct{}
	MU sync.Mutex
}

func CreateRoom(name string)*Room{
	return &Room{
		Name: name,
		Clients: make(map[*client.Client]struct{}),
	}
}

func (r *Room) GetAllClientsFromRoom()[]*client.Client{
	r.MU.Lock()
	defer r.MU.Unlock()

	clients := []*client.Client{}

	for client := range r.Clients{
		clients = append(clients, client)
	}

	return clients
}


func (r *Room) AddClientToRoom(client *client.Client){
	r.MU.Lock()
	defer r.MU.Unlock()

	r.Clients[client] = struct{}{}
}

func (r *Room) RemoveClientFromRoom(client *client.Client){
	r.MU.Lock()
	defer r.MU.Unlock()

	delete(r.Clients, client)
}

func (r *Room) DoesClientExistInRoom(client *client.Client)bool{
	r.MU.Lock()
	defer r.MU.Unlock()


	_, exists := r.Clients[client]
	if exists {
		return true
	}
	return  false
}
//message to clients to room
//do not add locks here

func (r *Room) BroadcastMessage(msg string){

	clients :=r.GetAllClientsFromRoom()

	for _ ,client := range clients{
		client.SendMessage(msg)
	}
}

func (r *Room) BroadcastMessageExceptClient(notclient *client.Client, msg string){
	clients :=r.GetAllClientsFromRoom()

	
	for _ ,client := range clients{
		if client != notclient {
			client.SendMessage(msg)
		}
		
	}
}




