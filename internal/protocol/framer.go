package protocol

import "fmt"


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

	f.buffer = append(f.buffer, data...)


	for _, val := range f.buffer {
		fmt.Println(val)
	}


	return []string{"dkbabd"}
}




