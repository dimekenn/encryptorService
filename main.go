package main

import (
	"context"
	"encryptService/server"
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main()  {
	ctx := context.Background()

	errChan := make(chan error, 1)

	go func(){
		sigCh := make(chan os.Signal)
		signal.Notify(sigCh, syscall.SIGTERM, syscall.SIGINT)
		errChan <- fmt.Errorf("%s", <-sigCh)
	}()

	go server.StartHTTPServer(ctx, errChan)

	fmt.Printf("Terminated: %s", <-errChan)
}
