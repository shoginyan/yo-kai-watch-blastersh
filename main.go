package main

import (
	"sync"

	"github.com/PretendoNetwork/yo-kai-watch-blasters/nex"
)

var wg sync.WaitGroup

func main() {
	wg.Add(2)

	// TODO - Add gRPC server
	go nex.StartAuthenticationServer()
	go nex.StartSecureServer()

	wg.Wait()
}
