package server

import (
	"context"
	"encryptService/config"
	"encryptService/handler"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"time"
)
//server
func StartHTTPServer(ctx context.Context, errCh chan <- error)  {
	fmt.Println("listening")

	router := mux.NewRouter()

	cfg, err := config.LoadConfiguration("./config.json")
	if err != nil{
		errCh <- err
	}

	handler2 := handler.StringHandler

	router.HandleFunc("/encryptor", handler2).Methods("POST")

	srv := &http.Server{
		Addr: cfg.Port,
		Handler: router,
		WriteTimeout: 15*time.Second,
		ReadTimeout: 15*time.Second,
	}

	err = srv.ListenAndServe()
	if err != nil{
		fmt.Println("listen and serve error")
	}
}
