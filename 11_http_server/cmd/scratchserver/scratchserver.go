package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
)

type ResponseDTO struct {
	Status string `json:"status"`
}

type handler struct {}

func (h *handler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	resp := &ResponseDTO{Status: "ok"}
	respBody, err := json.Marshal(resp)
	if err != nil {
		log.Println(err)
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	writer.Header().Add("Content-Type", "application/json")
	// по умолчанию статус 200 Ok
	_, err = writer.Write(respBody)
	if err != nil {
		log.Println(err)
	}
}

func main() {
	server := &http.Server{
		Addr: "0.0.0.0:9999",
		Handler: &handler{},
	}
	err := server.ListenAndServe()
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
}
