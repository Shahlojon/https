package main

import (
	"github.com/Shahlojon/http/pkg/server"
	// "time"
	// "fmt"
	// "io/ioutil"
	"strconv"
	// "strings"
	// "bytes"
	// "io"
	"log"
	"net"
	"os"
)

func main() {
	host := "0.0.0.0"
	port := "9999"
	if err := execute(host, port); err != nil {
		os.Exit(1)
	}
}

func execute(host string, port string) (err error) {
	srv:=server.NewServer(net.JoinHostPort(host, port))
	srv.Register("/", func(req *server.Request) {
		body:="Welcome to our web-site"
		id:=req.QueryParams["id"]
		log.Print(id)
		
		_, err=req.Conn.Write([]byte(
			"HTTP/1.1 200 OK\r\n"+
			"Content-Lenght: "+ strconv.Itoa(len(body))+"\r\n"+
			"Content-Type: text/html\r\n"+
			"Connection: close\r\n"+
			"\r\n"+
			body,
		))
	
		if err!=nil {
			log.Print(err)
		}
	})

	srv.Register("/payment/{id}", func(req *server.Request) {
		id:=req.PathParams["id"]
		log.Print(id)
	})

	srv.Register("/about", func(req *server.Request) {
		body:="About Golang Academy"

		_, err=req.Conn.Write([]byte(
			"HTTP/1.1 200 OK\r\n"+
			"Content-Lenght: "+ strconv.Itoa(len(body))+"\r\n"+
			"Content-Type: text/html\r\n"+
			"Connection: close\r\n"+
			"\r\n"+
			body,
		))
	
		if err!=nil {
			log.Print(err)
		}
	})
	return srv.Start()
	
	// listener, err := net.Listen("tcp", net.JoinHostPort(host, port))
	// if err != nil {
	// 	log.Print(err)
	// 	return err
	// }
	// defer func() {
	// 	if cerr := listener.Close(); cerr != nil {
	// 		if err == nil {
	// 			err = cerr
	// 			return
	// 		}
	// 		log.Print(cerr)
	// 	}
	// }()

	// for {
	// 	conn, err := listener.Accept()
	// 	if err != nil {
	// 		log.Print(err)
	// 		continue
	// 	}

	// 	handle(conn)
	// }
	// return
}

// func handle(conn net.Conn) {
// 	defer func() {
// 	  if cerr := conn.Close(); cerr != nil {
// 		log.Println(cerr)
  
// 	  }
// 	}()
  
// 	buf := make([]byte, 4096)
// 	for {
// 		n, err := conn.Read(buf)
// 		if err == io.EOF {
// 			log.Printf("%s", buf[:n])
// 		}
// 		if err != nil {
// 			log.Println(err)
// 			return
// 		}
	
// 		data := buf[:n]
// 		rLD := []byte{'\r', '\n'}
// 		rLE := bytes.Index(data, rLD)
// 		if rLE == -1 {
// 			log.Printf("Bad Request")
// 			return
// 		}
	
// 		reqLine := string(data[:rLE])
// 		parts := strings.Split(reqLine, " ")
	
// 		if len(parts) != 3 {
// 			log.Println("ErrBadRequest")
// 			return
// 		}
	
// 		path, version := parts[1], parts[2]
	
// 		if version != "HTTP/1.1" {
// 			log.Println("ErrHTTPVersionNotValid")
// 			return
// 		}

// 		if path == "/" {
// 			body, err :=ioutil.ReadFile("static/index.html")
// 			if err != nil {
// 				fmt.Errorf("can't read index.html: %w", err)
// 			}
// 			marker:="{{year}}"
// 			year:=time.Now().Year()
// 			body = bytes.ReplaceAll(body, []byte(marker), []byte(strconv.Itoa(year)))
// 			_, err=conn.Write([]byte(
// 				"HTTP/1.1 200 OK\r\n"+
// 				"Content-Lenght: "+ strconv.Itoa(len(body))+"\r\n"+
// 				"Content-Type: text/html\r\n"+
// 				"Connection: close\r\n"+
// 				"\r\n"+
// 				string(body),
// 			))

// 			if err!=nil {
// 				log.Println(err)
// 			}
// 	    }
// 	}
// }