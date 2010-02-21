package swif

import (
	"http"
	"io"
	"fmt"
	"xml"
	"os"
)


type networkconf struct {
	Address string;
	Port string;
}

type menuconf struct {
	Label string;
	Action string;
}

type config struct {
	Network networkconf;
	Menuentry []menuconf;
}

type Swif struct {
	conf config;
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

func (s *Swif) handleAction(action string, c *http.Conn) {
	switch(action) {
		case "kill":
			io.WriteString(c, "Killing...")
			os.Exit(0)
		default:
			io.WriteString(c, "Unknown action")
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
	for i:=0; i < len(s.conf.Menuentry); i++ {
		fmt.Fprintf(c, "&lt;%s&gt;%s\n",
			s.conf.Menuentry[i].Action, s.conf.Menuentry[i].Label)
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
