package protocol

import (
	"strings"
)

// JOIN <username>
// SEND <message>
// QUIT
type CommandType string;

const (
	JOIN CommandType = "JOIN"
	MESSAGE CommandType = "MESSAGE"
	QUIT CommandType = "QUIT"
	PING CommandType = "PING" // this is for the timout imp
	PONG CommandType = "PONG"
)

type Command struct {
	Type CommandType
	Arg string
}


func Parser(message string) (*Command, error){
	command := &Command{}
	if message == ""{
		return nil, NoCommandProvided
	}
	firstSpace := strings.Index(message, " ")
	

	//accept message as A B where A = string and B = string
	// B can be = ""
	if firstSpace == -1{
		//single string with no space
		command.Type = CommandType(message)
		
	}else{
		command.Type = CommandType(message[:firstSpace])
		command.Arg = message[firstSpace+1:]
	}



	validCommand := false

	if command.Type == QUIT && command.Arg == ""{
		validCommand = true
	}else if command.Type == JOIN && 
			command.Arg != "" && 
			!strings.ContainsAny(command.Arg, " \n\t\r"){
		validCommand = true
	}else if command.Type == MESSAGE{
		validCommand = true
	}
	
	if validCommand{
		return command, nil
	}


	return command, InvalidCommand
}

