package server

import (
	"net/url"
	"io"
	"log"
	"bytes"
	"strings"
	"sync"
	"net"
)

type HandlerFunc func(conn net.Conn)

type Server struct {
	addr string
	mu sync.RWMutex
	handlers map[string]HandlerFunc
}

//Request ...
type Request struct {
	Conn        net.Conn
	QueryParams url.Values
	PathParams  map[string]string
	Headers     map[string]string
	Body        []byte
}

func NewServer(addr string) *Server {
	return &Server{addr: addr, handlers: make(map[string]HandlerFunc)}
}

func (s *Server) Register(path string, handler HandlerFunc) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.handlers[path] = handler
}

func (s *Server) Start() (err error) {
	listener, err := net.Listen("tcp", s.addr)
	if err != nil {
		log.Println(err)
		return err
	}
	defer func() {
		if cerr := listener.Close(); cerr != nil {
			err = cerr
			return
		}
	}()
	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}

		go s.handle(conn)

	}
}

func (s *Server) handle(conn net.Conn) {
	defer conn.Close()
  
	buf := make([]byte, 4096)
	for {
	  n, err := conn.Read(buf)
	  if err == io.EOF {
		log.Printf("%s", buf[:n])
	  }
	  if err != nil {
		log.Println(err)
		return
	  }
  
	  data := buf[:n]
	  rLD := []byte{'\r', '\n'}
	  rLE := bytes.Index(data, rLD)
	  if rLE == -1 {
		log.Println("ErrBadRequest")
		return
	  }
  
	  reqLine := string(data[:rLE])
	  parts := strings.Split(reqLine, " ")
  
	  if len(parts) != 3 {
		log.Println("ErrBadRequest")
		return
	  }
  
	  path, version := parts[1], parts[2]
  
	  if version != "HTTP/1.1" {
		log.Println("ErrHTTPVersionNotValid")
		return
	  }
  
	  var handler = func(conn net.Conn) {
		conn.Close()
	  }
	  s.mu.RLock()
	  for i := 0; i < len(s.handlers); i++ {
		if hr, found := s.handlers[path]; found {
		  handler = hr
		  break
		}
	  }
	  s.mu.RUnlock()
	  handler(conn) 
	} 
}