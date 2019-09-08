package main

import (
	"context"
	"flag"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"
	"github.com/twsouza/tax-packet-shipping/app"
	"github.com/twsouza/tax-packet-shipping/common"
)

func main() {
	log.Println("Starting app")

	r := mux.NewRouter()
	r.HandleFunc("/tax", app.CalcProdTax).Methods("POST")
	r.HandleFunc("/track", app.TrackProd).Methods("POST")
	r.HandleFunc("/track/{id}", app.GetTrack).Methods("GET")

	config, err := common.GetConfig()
	if err != nil {
		log.Println("Err db.connect->common.GetConfig", err)
	}

	var wait time.Duration
	flag.DurationVar(&wait, "graceful-timeout", time.Second*config.ServerParams.GracefulTimeout, "the duration for which the server gracefully wait for existing connections to finish")
	flag.Parse()

	srv := &http.Server{
		Addr:         config.ServerParams.Addr,
		WriteTimeout: time.Second * config.ServerParams.WriteTimeout,
		ReadTimeout:  time.Second * config.ServerParams.ReadTimeout,
		IdleTimeout:  time.Second * config.ServerParams.IdleTimeout,
		Handler:      r,
	}

	log.Println("App is now listening...")
	// Run our server in a goroutine so that it doesn't block.
	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Println("err on server", err)
		}
	}()

	c := make(chan os.Signal, 1)
	// We'll accept graceful shutdowns when quit via SIGINT (Ctrl+C)
	// SIGKILL, SIGQUIT or SIGTERM (Ctrl+/) will not be caught.
	signal.Notify(c, os.Interrupt)

	// Block until we receive our signal.
	<-c

	// Create a deadline to wait for.
	ctx, cancel := context.WithTimeout(context.Background(), wait)
	defer cancel()
	// Doesn't block if no connections, but will otherwise wait
	// until the timeout deadline.
	srv.Shutdown(ctx)
	log.Println("shutting down")
	os.Exit(0)
}
