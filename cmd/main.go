package main

import (
//	"bytes"
//	"fmt"
//	"io"
//	"io/ioutil"
//	"log"
	"net"
	"os"
//	"strconv"
//	"strings"
	"github.com/Iftikhor99/server/pkg/server"
)

func main() {

	host := "0.0.0.0"
	port := "9999"

	if err := execute(host, port); err != nil {

		os.Exit(1)
	}
}

func execute(host string, port string) (err error) {

	
	srv := server.NewServer(net.JoinHostPort(host, port))

	srv.Register("api/payments", func(req *server.Request) {
		// id := req.PathParams["id"]
		// ctID := req.PathParams["category"]
		// log.Print("ID printed ", id, ctID)
		
	})
	


	// srv.Register("/", func(conn net.Conn) {
	// 	body := "Welcome to our web-site"

	// 	_, err = conn.Write([]byte(
	// 		"HTTP/1.1 200 OK\r\n" +
	// 			"Content-Length: " + strconv.Itoa(len(body)) + "“\r\n" +
	// 			"Content-Type: text/html\r\n" +
	// 			"Connection: close\r\n" +
	// 			"\r\n" +
	// 			body,
	// 	))
	// 	if err != nil {
	// 		log.Print(err)
	// 	}
	// })
	// srv.Register("/about", func(conn net.Conn) {
	// 	body := "About Golang Academy"

	// 	_, err = conn.Write([]byte(
	// 		"HTTP/1.1 200 OK\r\n" +
	// 			"Content-Length: “ + strconv.Itoa(len(body)) + “\r\n" +
	// 			"Content-Type: text/html\r\n" +
	// 			"Connection: close\r\n" +
	// 			"\r\n" +
	// 			body,
	// 	))
	// 	if err != nil {
	// 		log.Print(err)
	// 	}
	// })

	errS := srv.Start()
	// log.Print("We now here")
	// conn, err := net.Dial("tcp", "127.0.0.1:9999") 
    // if err != nil { 
    //     log.Println(err) 
        
    // } 
    // defer conn.Close() 

	return errS


}

// func handle(conn net.Conn) (err error) {
// 	defer func() {
// 		if cerr := conn.Close(); cerr != nil {
// 			if err == nil {
// 				err = cerr
// 				return
// 			}
// 			log.Print(err)
// 		}
// 	}()

// 	buf := make([]byte, 4096)

// 	n, err := conn.Read(buf)
// 	if err == io.EOF {
// 		log.Printf("%s", buf[:n])
// 		return nil
// 	}
// 	if err != nil {
// 		return err
// 	}
// 	log.Printf("%s", buf[:n])

// 	data := buf[:n]
// 	requestLineDeLim := []byte{'\r', '\n'}
// 	requestLineEnd := bytes.Index(data, requestLineDeLim)
// 	if requestLineEnd == -1 {

// 	}

// 	requestLine := string(data[:requestLineEnd])
// 	parts := strings.Split(requestLine, " ")
// 	if len(parts) != 3 {

// 	}

// 	method, path, version := parts[0], parts[1], parts[2]

// 	if method != "GET" {

// 	}

// 	if version != "HTTP/1.1" {

// 	}

// 	if path == "/" {
// 		//	body := "Ok!"
// 		body, err := ioutil.ReadFile("static/index.html")
// 		if err != nil {
// 			return fmt.Errorf("can't read index.html: %w", err)
// 		}
// 		_, err = conn.Write([]byte(
// 			"HTTP/1.1 200 OK\r\n" +
// 				"Conect-Length: " + strconv.Itoa(len(body)) + "\r\n" +
// 				"Content-Type: text/html	\r\n" +
// 				"Connection: close\r\n" +
// 				"\r\n" +
// 				string(body),
// 		))
// 		if err != nil {
// 			return err
// 		}
// 	}

// 	return nil
// }
