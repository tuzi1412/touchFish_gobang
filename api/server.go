package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/tuzi1412/touchFish_gobang/config"
)

type Server struct {
	router *mux.Router
	port   string
}

func NewServer() *Server {
	return &Server{
		router: mux.NewRouter(),
		port:   "22333",
	}
}

func (s *Server) Run() {
	s.routerInit()
	server := http.Server{
		Addr:    ":" + s.port,
		Handler: s.router,
	}
	go server.ListenAndServe()
}

func (s *Server) routerInit() {
	base := s.router.PathPrefix("/touchFish_gobang").Subrouter()
	base.HandleFunc("/", s.processMap).Methods(http.MethodPut)
	base.HandleFunc("/testConnect", s.testConnect).Methods(http.MethodGet)
	base.HandleFunc("/chooseChess", s.chooseChess).Methods(http.MethodPut)
}

func (s *Server) processMap(w http.ResponseWriter, r *http.Request) {
	var msg config.HTTPRsp
	err := json.NewDecoder(r.Body).Decode(&msg)
	if err != nil {
		s.rspError(w, "data error", err)
		return
	}
	go func() { config.MapChan <- msg.Data }()
	s.rspMap(w, msg.Data)
}

func (s *Server) chooseChess(w http.ResponseWriter, r *http.Request) {
	var msg config.HTTPRsp
	err := json.NewDecoder(r.Body).Decode(&msg)
	if err != nil {
		s.rspError(w, "data error", err)
		return
	}
	if config.MyChess == 0 {
		if config.RandomNum > msg.Code {
			config.MyChess = 1
		} else {
			config.MyChess = 2
		}
	}
	var rsp config.HTTPRsp
	rsp.Code = config.RandomNum
	rsp.Message = "success"
	rspbyte, _ := json.Marshal(rsp)
	w.Header().Add("Content-Type", "application/json")
	w.Write(rspbyte)
}

func (s *Server) testConnect(w http.ResponseWriter, r *http.Request) {
	s.rspOk(w)
}

func (s *Server) rspOk(w http.ResponseWriter) {
	var rsp config.HTTPRsp

	rsp.Code = 0
	rsp.Message = "success"
	rspbyte, _ := json.Marshal(rsp)
	w.Header().Add("Content-Type", "application/json")
	w.Write(rspbyte)
}

func (s *Server) rspError(w http.ResponseWriter, msg string, err error) {
	var rsp config.HTTPRsp
	fmt.Println(err)
	rsp.Code = http.StatusInternalServerError
	rsp.Message = msg
	rspbyte, _ := json.Marshal(rsp)
	w.Header().Add("Content-Type", "application/json")
	w.Write(rspbyte)
}

func (s *Server) rspMap(w http.ResponseWriter, msg [15][15]uint8) {
	var rsp config.HTTPRsp
	rsp.Code = 0
	rsp.Message = "success"
	rsp.Data = msg
	rspbyte, _ := json.Marshal(rsp)
	w.Header().Add("Content-Type", "application/json")
	w.Write(rspbyte)
}
