package room

import (
	"sync"
	"tcp-chat/internal/client"
)


type Room struct {
	Name string
	Clients map[*client.Client]struct{}
	mu sync.Mutex
}

func CreateRoom(name string)*Room{
	return &Room{
		Name: name,
		Clients: make(map[*client.Client]struct{}),
	}
}


func (r *Room) AddClientToRoom(client *client.Client){
	r.mu.Lock()
	defer r.mu.Unlock()

	r.Clients[client] = struct{}{}
}

func (r *Room) RemoveClientFromRoom(client *client.Client){
	r.mu.Lock()
	defer r.mu.Unlock()
	delete(r.Clients, client)
}
func (r *Room) DoesClientExistInRoom(client *client.Client)bool{
	r.mu.Lock()
	defer r.mu.Unlock()

	_, exists := r.Clients[client]
	if exists {
		return true
	}
	return  false
}

func (r *Room) BroadcastMessage(msg string){
	r.mu.Lock()
	defer r.mu.Unlock()

	for client := range r.Clients{
		client.SendMessage(msg)
	}
}

func (r *Room) BroadcastMessageExceptClient(notclient *client.Client, msg string){
	r.mu.Lock()
	defer r.mu.Unlock()

	for client := range r.Clients{
		if client != notclient {
			client.SendMessage(msg)
		}
		
	}
}




