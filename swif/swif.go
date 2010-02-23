package swif

import (
	"http"
	"io"
	"fmt"
	"xml"
	"os"
)


type networkconf struct {
	Address string
	Port string
}

type menuconf struct {
	Label string
	Action string
}

type config struct {
	Network networkconf
	Menuentry []menuconf
}

type Swif struct {
	conf config
	handler map[string] SwifHandler	
}

func (s *Swif) RegisterHandler(name string, r SwifHandler) {
	s.handler[name] = r 
}

type SwifHandler interface {
	HandleAction (c *http.Conn)
}

type HandlerFunc func(c *http.Conn)

func (f HandlerFunc) HandleAction(c *http.Conn) {
	f(c)
}

func KillHandler(c *http.Conn) {
	fmt.Fprintf(c, "Killing now...\n") 
	os.Exit(0)
}

func (s *Swif) handleError(code int, c *http.Conn) {
	switch(code) {
		case http.StatusBadRequest:
			c.WriteHeader(code)
			io.WriteString(c, "Invalid Request")
		default:
			c.WriteHeader(http.StatusBadRequest)
			io.WriteString(c, "Invalid Request")
	}
}


func NewSwif() *Swif {
	s := new(Swif)
	s.handler = make(map[string] SwifHandler)
	s.RegisterHandler("kill", HandlerFunc(KillHandler))
	return s
}

func (s *Swif) handleAction(actionname string, c *http.Conn) {
	h, ok := s.handler[actionname]
	if !ok {
			io.WriteString(c, "Unknown action")
	} else {
		h.HandleAction(c)
	}
}

func (s *Swif) ServeHTTP(c *http.Conn, req *http.Request) {
	err := req.ParseForm()

	if err != nil {
		s.handleError(http.StatusBadRequest, c)
		return
	}

	action, isAction := req.Form["action"]
	if isAction {
		s.handleAction(action[0], c)
		return
	}

	s.printMenu(c)
}

func (s *Swif) printMenu(c *http.Conn) {
	fmt.Fprintf(c, "%d\n", len(s.handler))
	for i:=0; i < len(s.handler); i++ {
//		fmt.Fprintf(c, "&lt;%s&gt;%s\n",
//			s.handler[i], "bla")
	}
}

func (s *Swif) ReadConfig(r io.Reader) os.Error {
	err := xml.Unmarshal(r, &s.conf)
	return err
}

func (s *Swif) Start() os.Error {
	http.Handle("/", s)
	err := http.ListenAndServe(s.conf.Network.Address+":"+s.conf.Network.Port, nil)
	return err
}
