package framer

import (
	"bytes"
	"strings"

)


type Framer struct {
	buffer []byte
}

func NewFramer() *Framer {
	return &Framer{}
}


func (f *Framer) Feed(data []byte)[]string{
	// Take whatever bytes TCP just gave us, 
	// add them to our existing bytes, 
	// find every complete \n-terminated message, 
	// return those messages, 
	// and retain anything incomplete.
	messages := []string{}


	f.buffer = append(f.buffer, data...)

	for {
		breakPt := bytes.IndexByte(f.buffer, '|')
		if breakPt == -1 {
			break
		}
		message := strings.TrimSpace(string(f.buffer[:breakPt]))
		messages = append(messages, message)

		f.buffer = f.buffer[breakPt+1:]


	}

	return messages
}




