package jobs


import (
	"tcp-chat/internal/client"
	"tcp-chat/internal/protocol"
	"tcp-chat/internal/server"
	"tcp-chat/internal/room"
	"github.com/google/uuid"

)


type Job struct{
	JobId string
	Job []string
	Client *client.Client
	Server *server.Server
}

func NewJob(job []string, client *client.Client, server *server.Server) *Job{
	return &Job{
		JobId: uuid.New().String(),
		Job: job,
		Client: client,
		Server: server,
	}
}

func (j *Job) JobProcessor() {
	for _, message := range j.Job{
		command, err := protocol.Parser(message)
		if err != nil{
			j.Client.Conn.Write([]byte("invalid command\n"))
			continue
		}
			

		switch command.Type{
			case protocol.JOIN:
				
				// add user to room
				roomExists := j.Server.GetRoom(command.Arg)

				// case 1 => room does not exist
				if roomExists == nil{
				// room does not exist create room and add j.Client.Conn to room
					newRoom := room.CreateRoom(command.Arg)
					newRoom.AddClientToRoom(j.Client)
					j.Server.AddRoomToServer(newRoom)
					break
				}
				// case 2 => j.Client join room the j.Client already is in 
				if roomExists.DoesClientExistInRoom(j.Client){
					j.Client.SendMessage("Already in room")
					break
				}


				//case 3 => cleint joins new room
				roomsClientExistsIn := j.Server.RoomsClientExistsIn(j.Client)

				
				// remove j.Client from any existing room
				for _, room_val := range roomsClientExistsIn {
					room_val.RemoveClientFromRoom(j.Client)
				}
				roomExists.AddClientToRoom(j.Client)	
				
				


				
			case protocol.MESSAGE:
				roomsClientExistsIn := j.Server.RoomsClientExistsIn(j.Client)
				if len(roomsClientExistsIn) ==0 {
					j.Client.SendMessage("You are not in any room")
					break
				}

				for _,room := range roomsClientExistsIn{
					room.BroadcastMessageExceptClient(j.Client, command.Arg)
				}
				

			case protocol.QUIT:
				j.Server.RemoveClientFromAllRooms(j.Client)

			case protocol.PONG:
				if j.Client.SentPing {
					j.Client.UpdateLastActivity()
					j.Client.SentPing = false
				}





			
		}


		}
}