package main

import (
	"fmt"
	"lanshare/lanshare"
	"os"
)

func main() {
	clientType := lanshare.Receiver
	if len(os.Args) > 1 && os.Args[1] == "s" {
		fmt.Println("sender")
		clientType = lanshare.Sender
	}
	c := lanshare.Client{
		Type: clientType,
	}

	c.Run()
}
