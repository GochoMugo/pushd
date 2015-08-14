package main

import (
	"os"
	"os/user"

	"github.com/xconstruct/go-pushbullet"
)

// Pusher is a clas handling pushing notifications
type Pusher struct {
	Client  *pushbullet.Client
	Devices []*pushbullet.Device
	User    *pushbullet.User
	Name    string
}

// NewPusher creates a pushbullet client
func NewPusher(token string) *Pusher {
	p := Pusher{}
	hostname, err := os.Hostname()
	currentuser, err := user.Current()
	p.Name = currentuser.Username + "@" + hostname
	p.Client = pushbullet.New(token)
	devices, err := p.Client.Devices()
	if err != nil {
		panic(err)
	}
	p.Devices = devices
	user, err := p.Client.Me()
	if err != nil {
		panic(err)
	}
	p.User = user
	return &p
}

// Send a message to all of your devices
func (p *Pusher) Send(message string) error {
	err := p.Client.PushNote("", "(Pushd Alert) "+p.Name, message)
	if err != nil {
		return err
	}
	return nil
}
