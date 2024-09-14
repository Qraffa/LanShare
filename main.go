package main

import "lanshare/lanshare"

func main() {
	c := lanshare.Client{
		Type: lanshare.Receiver,
	}

	c.Run()
}
