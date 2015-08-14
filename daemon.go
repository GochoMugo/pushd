package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"net/url"
	"os"
	"time"

	"gopkg.in/tylerb/graceful.v1"
)

// Default Port
const DefaultPort = "9701"

func asPort(port string) string {
	if port == "" {
		port = DefaultPort
	}
	return port
}

// Daemon is a pushbullet daemon
type Daemon struct {
	Port   string
	Pusher *Pusher
	Server *graceful.Server
}

func getToken() string {
	token := os.Getenv("PUSHER_TOKEN")
	return token
}

func write(w http.ResponseWriter, message string) {
	w.Header().Set("content-type", "application/json")
	fmt.Fprintf(w, `{"message":"`+message+`"}`)
}

func (d *Daemon) handlePushRequests(w http.ResponseWriter, r *http.Request) {
	message := r.FormValue("message")
	err := d.Pusher.Send(message)
	if err != nil {
		print(err)
		write(w, "errored")
		return
	}
	write(w, "sent")
}

func (d *Daemon) handleStopRequest(w http.ResponseWriter, r *http.Request) {
	write(w, "acknowledged")
	d.Server.Stop(10 * time.Second)
}

func (d *Daemon) handlePing(w http.ResponseWriter, r *http.Request) {
	write(w, "pong")
}

// Start starts the daemon
func (d *Daemon) Start() {
	d.Port = asPort(d.Port)
	d.Pusher = NewPusher(getToken())
	mux := http.NewServeMux()
	mux.HandleFunc("/", d.handlePushRequests)
	mux.HandleFunc("/stop", d.handleStopRequest)
	mux.HandleFunc("/ping", d.handlePing)

	d.Server = &graceful.Server{
		Timeout: 0,
		ConnState: func(conn net.Conn, state http.ConnState) {
			// conn has a new state
		},
		Server: &http.Server{
			Addr:    ":" + d.Port,
			Handler: mux,
		},
	}
	err := d.Server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}

// StopDaemon stops a daemon server
func StopDaemon(port string) ([]byte, error) {
	port = asPort(port)
	resp, err := http.Get("http://localhost:" + port + "/stop")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	return body, err
}

// CheckDaemonStatus checks the status of the daemon
func CheckDaemonStatus(port string) ([]byte, error) {
	port = asPort(port)
	resp, err := http.Get("http://localhost:" + port + "/ping")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	return body, err
}

// SendNotification sends a notification using the daemon's port
func SendNotification(port string, message string) ([]byte, error) {
	port = asPort(port)
	resp, err := http.Get("http://localhost:" + port + "/?message=" + url.QueryEscape(message) + "")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	return body, err
}
