package app

import (
	"encoding/json"
	"lectionhttpserver/cmd/bank/app/dto"
	"lectionhttpserver/pkg/card"
	"log"
	"net/http"
)

type Server struct {
	cardSvc *card.Service
	mux *http.ServeMux
}

func NewServer(cardSvc *card.Service, mux *http.ServeMux) *Server {
	return &Server{cardSvc: cardSvc, mux: mux}
}

func (s *Server) Init() {
	s.mux.HandleFunc("/getCards", s.getCards)
	s.mux.HandleFunc("/addCard", s.addCard)
	s.mux.HandleFunc("/editCard", s.editCard)
	s.mux.HandleFunc("/removeCard", s.removeCard)
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.mux.ServeHTTP(w, r)
}

func (s *Server) getCards(w http.ResponseWriter, r *http.Request) {
	cards := s.cardSvc.All(r.Context())
	dtos := make([]*dto.CardDTO, len(cards))
	for i, c := range cards {
		dtos[i] = &dto.CardDTO{
			Id:     c.Id,
			Number: c.Number,
		}
	}

	// TODO: вынести в отдельную функцию
	respBody, err := json.Marshal(dtos)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	// по умолчанию статус 200 Ok
	_, err = w.Write(respBody)
	if err != nil {
		log.Println(err)
	}
}

func (s *Server) addCard(w http.ResponseWriter, r *http.Request) {
	panic("implement me")
}

func (s *Server) editCard(w http.ResponseWriter, r *http.Request) {
	panic("implement me")
}

func (s *Server) removeCard(w http.ResponseWriter, r *http.Request) {
	panic("implement me")
}

