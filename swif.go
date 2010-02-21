package main

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

var conf config

func TestServer(c *http.Conn, req *http.Request) {
	c.WriteHeader(http.StatusOK)
	io.WriteString(c, "<html><body><pre>")
	fmt.Fprintf(c, "%s:%s", conf.Network.Address, conf.Network.Port)
	io.WriteString(c, "</pre></body></html>")
}

func main() {
	f, err := os.Open("swif.conf.xml", os.O_RDONLY, 0)
	if err != nil {
		panic("Open\n", err.String())
	}
	err = xml.Unmarshal(f, &conf)
	if err != nil {
		panic("Unmarshal\n", err.String())
	}
	f.Close()

	http.Handle("/", http.HandlerFunc(TestServer))
	err = http.ListenAndServe(conf.Network.Address+":"+conf.Network.Port, nil)
	if err != nil {
		panic("ListenAndServe\n", err.String())
	}

}
