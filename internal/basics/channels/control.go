package channels

import (
	"fmt"
	"time"
)

type Server struct {
	quitch chan struct{}
	msgch  chan string
}

func (s *Server) Start() {
	fmt.Println("Starting server")
	s.Loop() //blocking call
}

func (s *Server) Stop() {
	println("Stopping server")
	close(s.quitch)
}

func (s *Server) Send(msg string) {
	s.msgch <- msg
}

func (s *Server) Loop() {
	for {
		select {
		case msg := <-s.msgch:
			s.HandleMessage(msg)
		case <-s.quitch:
			return
		}
	}
}

func (s *Server) HandleMessage(msg string) {
	fmt.Printf("We have a message: %v", msg)
}

func newServer() *Server {
	return &Server{
		quitch: make(chan struct{}),
		msgch:  make(chan string, 128),
	}
}

func RunControlledChannels() {
	server := newServer()

	go func() {
		time.Sleep(time.Second * 5)
		server.Send("Hello")
		server.Stop()
	}()

	server.Start()
}
