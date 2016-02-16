package main

import (
	"fmt"
	"trident.li/xmpp-go/xlib"
)

/* A very simple XMPP client */

var ChatCmds chan interface{}

type cmdRecvMessage struct {
	timestamp string
	from      string
	to        string
	channel   string
	msg       string
}

type cmdSendMessage struct {
	from string
	to   string
	msg  string
}

type cmdJoinMUC struct {
	to string
}

func ReceiveMsg(s *xlib.Session, m cmdRecvMessage) {
	fmt.Printf("[mesg] stamp:%s from:%s, to:%s channel:%s msg:%s\n", m.timestamp, m.from, m.to, m.channel, m.msg)
}

func main() {
	/* Settings */
	user := "you@example.com"
	pass := "passw@rd"
	muc := "channel@conference.example.com"
	logfile := "/tmp/xmpp-go.log"

	/* Command Channel (to be used so other goroutines can affect state too, but async) */
	ChatCmds = make(chan interface{}, 100)
	xio := NewSimpleXIO()

	config := &xlib.Config{Account: user, Password: pass, RawLogFile: logfile}

	lgr := xlib.NewLineLogger(xio)

	s, err := xlib.Connect(xio, config, lgr, nil)
	if err != nil {
		xio.Alert("Failed to connect: " + err.Error())
		return
	}

	s.SignalPresence("")

	go s.Handle()

	/* Join a channel */
	ChatCmds <- cmdJoinMUC{muc}

	running := true

	for running {
		select {
		case cmd, ok := <-ChatCmds:
			if !ok {
				xio.Warn("Exiting because command channel closed")
				running = false
				break
			}

			s.LastAction()

			switch cmd := cmd.(type) {
			case cmdRecvMessage:
				ReceiveMsg(s, cmd)
				break

			case cmdSendMessage:
				s.Msg(cmd.to, cmd.msg, nil)
				break

			case cmdJoinMUC:
				s.JoinMUC(cmd.to, "", "")
				break
			}
		}

	}
}
