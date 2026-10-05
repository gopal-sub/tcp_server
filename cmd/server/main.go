package main

import (
	"errors"
	"fmt"
	"tcp-chat/internal/server"
)





func main() {
	
	svr, err := server.NewServer(":3000")
	if err != nil {
		panic(err)
	}
	fmt.Println("server started")
	err = svr.Start()

	if errors.Is(err, server.ServerShutdown){
		svr.Waitgrp.Wait()
		fmt.Println("shutdown complete")
		return
	}
	
	if err != nil {
		fmt.Println("client connection error")
	}


}


