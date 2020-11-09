package server

import (
	"bytes"
	"io"
	"strings"
//	"strconv"
//	"fmt"
	"log"
	"net"
	"sync"
	"net/url"
)

//HandlerFunc for
type HandlerFunc func(req *Request)

//Server for
type Server struct {
	addr string

	mu sync.RWMutex

	handlers map[string]HandlerFunc
}


//Request for
type Request struct {
	Conn net.Conn
	QueryParams url.Values
	PathParams map[string]string
	Headers map[string]string
	Body []byte
}

//NewServer for
func NewServer(addr string) *Server {
	return &Server{addr: addr, handlers: make(map[string]HandlerFunc)}

}

//Register for
func (s *Server) Register(path string, handler HandlerFunc) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.handlers[path] = handler
}


//Start for
func (s *Server) Start() error {
	// TODO: start server on host & port
	listener, err := net.Listen("tcp", s.addr)
	if err != nil {
		log.Print(err)
		return err
	}
	defer func() {
		if cerr := listener.Close(); cerr != nil {

			if err == nil {
				err = cerr
				return
			}
			log.Print(cerr)
		}
	}()
	// TODO: server code

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		
		go s.handle(conn)
		// if err != nil {
		// 	log.Print(err)
		// 	continue
		// }

	}

	

//	return nil

}

func (s *Server) handle(conn net.Conn) {
	var err error
	//mu := s.mu
	pathParameter := map[string]string{}
	headerParameter := map[string]string{}
	defer func() {
		if cerr := conn.Close(); cerr != nil {
			if err == nil {
				err = cerr
				log.Print(err)
			}
			log.Print(err)
		}
	}()

	buf := make([]byte, 4096)

	n, err := conn.Read(buf)
	if err == io.EOF {
		log.Printf("%s", buf[:n])
		log.Print(err)
	}
	if err != nil {
		log.Print(err)
	}
	log.Printf("%s", buf[:n])

	data := buf[:n]
	requestLineDeLim := []byte{'\r', '\n'}
	requestLineEnd := bytes.Index(data, requestLineDeLim)
	if requestLineEnd == -1 {

	}

	dataAfterPathByte := data[requestLineEnd+2:]

	requestLineDeLim2 := []byte{'\r', '\n','\r', '\n'}
	endOfHeader := bytes.Index(dataAfterPathByte, requestLineDeLim2)
	if endOfHeader != -1 {
		dataAfterPathByte = dataAfterPathByte[:endOfHeader]	
	}
	
	bodyByte := dataAfterPathByte[endOfHeader+1:]

	dataAfterPath := string(dataAfterPathByte)
	
	log.Print(dataAfterPath)
	newDataAfterPath := strings.Split(dataAfterPath,"\r\n")
	for _, one := range newDataAfterPath {
		index1 := strings.Index(one, ":")
		if index1 == -1 {
			continue
		}
			key := strings.TrimSpace(one[:index1])
			
			value := strings.TrimSpace(one[index1+1:])
			
			headerParameter[key] = value
			log.Print("key ", key, " value ", value)
	}
	// endOfHeader := bytes.Index(dataAfterPath, requestLineDeLim2)
	// if endOfHeader != -1 {
	// 	dataAfterPath = dataAfterPath[:endOfHeader]	
	// }
	
	// lendataAfterPath := len(dataAfterPath)
	// log.Printf("%s", dataAfterPath)
	// ind := 0
	// for {
		
	// 	if len(dataAfterPath) < 4 {
	// 		break
	// 	}
	// 	requestLineEndNew := bytes.Index(dataAfterPath, requestLineDeLim)
	// 	ind += requestLineEndNew 
	// 	if requestLineEnd == -1 {
	// 		break
	// 	} else {
	// 		requestLineNew := string(dataAfterPath[:requestLineEndNew])
	// 		index1 := strings.Index(requestLineNew, ":")
	// 		key := strings.Trim(requestLineNew[:index1],"\r\n")
	// 		key = strings.Trim(key,"\r\n")
	// 		value := strings.Trim(requestLineNew[index1+1:],"\r\n")
	// 		value = strings.Trim(value,"\r\n")
	// 		headerParameter[key] = value
	// 		log.Print("key ", key, " value ", value)
	// 	}
	// 	if ind + 2 < lendataAfterPath {
	// 	dataAfterPath = dataAfterPath[requestLineEndNew+2:]
	// 	} else {
	// 		break
	// 	}
				
	// }

	requestLine := string(data[:requestLineEnd])
	parts := strings.Split(requestLine, " ")
	if len(parts) != 3 {

	}

	method, path, version := parts[0], parts[1], parts[2]

	if method != "GET" {

	}

	if version != "HTTP/1.1" {

	}

	uri, err := url.ParseRequestURI(path)

	if err != nil {
		log.Print(err)
		return
	}
	//[]string{'{','}'}

	handlerPath := ""
	for j := range s.handlers {
		log.Print("This is j ", j)
		handlerPath = j
	}
	// newHandlerPath1 := ""
	
	// indexToFindFirst := strings.Index(handlerPath, "{")
	// if indexToFindFirst != -1 {
	// 	indexToFindSecond := strings.Index(handlerPath, "}")
	// 	newHandlerPath1 = handlerPath[indexToFindFirst+1:indexToFindSecond] 
	// }
	// log.Print(newHandlerPath1)
	
	// newHandlerPath := strings.SplitAfter(handlerPath, "{")
	// if len(newHandlerPath) > 1 {
	// 	newHandlerPath[0] = strings.TrimRight(newHandlerPath[0], "{")
	// 	for z :=1; z < len(newHandlerPath); z++{
	// 		index := strings.Index(newHandlerPath[z], "}") 
	// 		newHandlerPath[z] = newHandlerPath[z][:index]
	// 	}
		
	// }
	// for t:=0; t < 3; t++{
	// 	strings.Index(handlerPath,"{")

	// } 

//	newHandlerPath := Split(handlerPath,"{","}")
	//log.Print(newHandlerPath)	
	// strings.Index(handlerPath,"{")
	// strings.SplitAfter(handlerPath, "{")
	// strings.
	// handlerPath.Split("{","}")
	log.Print(uri.Path)
	log.Print(uri.Query())
	//QueryParams: uri.Query()	

	urlPath := uri.Path
	//categoryIdValue := ""
	
	indexToFind := strings.Index(handlerPath, "{category}")
	if indexToFind != -1 {
		newURLPath := urlPath[indexToFind:] 
		index := strings.Index(newURLPath, "/")
		categoryIDValue := newURLPath[:index]
		pathParameter["category"] = categoryIDValue
	}
	
	indexToCheck := strings.Index(handlerPath, "{id}")
	if indexToCheck != -1 {

		partsPath := strings.Split(urlPath, "/")
	idPath := ""
	newPath := ""
//	if len(partsPath) == 3 {
		idPath = partsPath[len(partsPath)-1]
//	}
	for i:=0; i < len(partsPath)-1; i++{
		newPath += partsPath[i] + "/"
	}	
	pathParameter["id"] = idPath
		
	} 

	// headerParameter1 := map[string]string{
	// 	"Accept-Encoding": "gzip",
	// 	"Host": "localhost:8000",		
	// 	"User-Agent": "Go-http-client/1.1",
	// }
	
	
	//newPath += "{id}"
	newRequest := &Request{Conn: conn, QueryParams: uri.Query(), PathParams: pathParameter, Headers: headerParameter, Body: bodyByte}
	
	//if path == "/" {
		s.mu.RLock()
		handler, ok := s.handlers[handlerPath]
		//handler := s.handlers["/"]
		s.mu.RUnlock()
		if ok == true {
		//	log.Print("Ok printed ", handler)
			handler(newRequest)
		} else {
//			conn.Close()
			return
		}
	//} 
	// if path == "/about" {
	// 	s.mu.RLock()
	// 	handler, ok := s.handlers["/about"]
		
	// 	s.mu.RUnlock()
	// 	if ok == true {
	// 		handler(conn)	
	// 	}
	// 	//handler(conn)
	// //	log.Print(handler)
	// } 
	// s.mu.RLock()
	// 	handler, ok := s.handlers["/secret"]
		
	// 	s.mu.RUnlock()
	// 	if ok == true {
	// 		handler(conn)	
	// }
	//	handler(conn)
//	log.Print(handler, ok)
	// 	body := "Ok!"
	// //	body, err := ioutil.ReadFile("static/index.html")
	// 	if err != nil {
	// 		return fmt.Errorf("can't read index.html: %w", err)
	// 	}
	// 	_, err = conn.Write([]byte(
	// 		"HTTP/1.1 200 OK\r\n" +
	// 			"Conect-Length: " + strconv.Itoa(len(body)) + "\r\n" +
	// 			"Content-Type: text/html	\r\n" +
	// 			"Connection: close\r\n" +
	// 			"\r\n" +
	// 			string(body),
	// 	))
	// 	if err != nil {
	// 		return err
	// 	}
	
	
	// if path == "/about" {
	// 	s.mu.RLock()
	// 	handler, ok := s.handlers["/about"]
		
	// 	s.mu.RUnlock()
	// 	if ok == true {
	// 		handler(conn)	
	// 	}
	// 	//handler(conn)
	// //	log.Print(handler)
	// }

	
}

//Split for
func Split(str, before, after string) string {
    a := strings.SplitAfterN(str, before, 3)
    b := strings.SplitAfterN(a[len(a)-1], after, 3)
    if 1 == len(b) {
        return b[0]
    }
    return b[0][0:len(b[0])-len(after)]
}