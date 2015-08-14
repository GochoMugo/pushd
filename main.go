package main

import (
	"github.com/GochoMugo/argparse"
	"github.com/GochoMugo/go-out"
)

const pushdName = "pushd"
const pushdDecription = "A pushbullet daemon"
const pushdVersion = "0.0.0"
const pushdHomepage = "https://github.com/GochoMugo/pushd"

func main() {
	program := argparse.New()
	program.Description(pushdName, pushdDecription)
	program.Version(pushdVersion)
	program.Epilog("See " + pushdHomepage + " for feature requests and bug reports")
	program.Command("s", "start", "start daemon", func(a argparse.Args) {
		out.Info("starting daemon")
		d := new(Daemon)
		d.Port = a.AsString("port", "p")
		d.Start()
	})
	program.Command("x", "stop", "stop daemon", func(a argparse.Args) {
		out.Info("stopping daemon")
		res, err := StopDaemon(a.AsString("port", "p"))
		if err != nil {
			out.Error("errored: %s", err)
			return
		}
		out.Info("response: %s", res)
	})
	program.Command("?", "status", "check status", func(a argparse.Args) {
		out.Info("checking daemon status")
		res, err := CheckDaemonStatus(a.AsString("port", "p"))
		if err != nil {
			out.Error("errored: %s", err)
			return
		}
		out.Info("response: %s", res)
	})
	program.Command("n", "notify", "push notification", func(a argparse.Args) {
		out.Info("sending notification")
		res, err := SendNotification(a.AsString("port", "p"), a.AsString("message", "m"))
		if err != nil {
			out.Error("errored: %s", err)
			return
		}
		out.Info("response: %s", res)
	})
	program.Parse()
}
