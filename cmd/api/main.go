package main

import (
	"fmt"
	"github.com/go-chi/chi"
	"github.com/ryamgoh/go-api/internal/handlers"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func main() {
	log.SetReportCaller(true)
	var r *chi.Mux = chi.NewRouter()
	handlers.Handler(r)

	fmt.Println("Starting GO API Service")
	fmt.Println(`                                                       
  ▄▄█▀▀▀█▄█  ▄▄█▀▀██▄           ██     ▀███▀▀▀██▄▀████▀
▄██▀     ▀█▄██▀    ▀██▄        ▄██▄      ██   ▀██▄ ██  
██▀       ▀██▀      ▀██       ▄█▀██▄     ██   ▄██  ██  
██         ██        ██      ▄█  ▀██     ███████   ██  
██▄    ▀█████▄      ▄██      ████████    ██        ██  
▀██▄     ██▀██▄    ▄██▀     █▀      ██   ██        ██  
  ▀▀███████  ▀▀████▀▀     ▄███▄   ▄████▄████▄    ▄████▄`)

	err := http.ListenAndServe("localhost:8000", r)
	if err != nil {
		log.Error(err)
	}
}
