package main

import "fmt"





func main() {
	
	server, err := NewServer(":3000")
	if err != nil {
		panic(err)
	}
	fmt.Println("server started")
	err = server.Start()

	
	if err != nil {
		fmt.Println("client connection error")
	}


}


