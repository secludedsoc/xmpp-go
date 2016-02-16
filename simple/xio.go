package main

import (
	"errors"
	"fmt"
	"trident.li/xmpp-go/xlib"
)

type SimpleXIO struct {
}

func (xio *SimpleXIO) Log(format string, a ...interface{}) {
	fmt.Printf("XMPP:"+format, a...)
}

func (xio *SimpleXIO) Info(msg string) {
	xio.Log("[info] %s\n", msg)
}

func (xio *SimpleXIO) Warn(msg string) {
	xio.Log("[warn] %s\n", msg)
}

func (xio *SimpleXIO) Alert(msg string) {
	xio.Log("[alrt] %s\n", msg)
}

func (xio *SimpleXIO) Critical(msg string) {
	xio.Log("[crit] %s\n", msg)
}

func (xio *SimpleXIO) ReadPassword(msg string) (password string, err error) {
	/* Noop */
	return
}

func (xio *SimpleXIO) SetPrompt(prompt string) {
	/* Noop */
}

func (xio *SimpleXIO) SetPromptEnc(target string, isEncrypted bool) {
	/* Noop */
}

/*
 * Any inbound message we send to our command pipe
 * so that it gets processed by the main Chat goroutine
 */
func (xio *SimpleXIO) Message(timestamp, from, to, channel string, msg []byte, isEncrypted bool, doBell bool) {
	ChatCmds <- cmdRecvMessage{timestamp, from, to, channel, string(msg)}
}

func (xio *SimpleXIO) StatusUpdate(timestamp, from, channel, show, status string, gone bool) {
	xio.Log("[stat] stamp:%s from:%s channel:%s status:>>>%s<<< show:>>>%s<<< gone: %v\n", timestamp, from, channel, status, show, gone)
}

func (xio *SimpleXIO) FormStringForPrinting(s string) string {
	return s
}

func (xio *SimpleXIO) Write(s string) {
	xio.Log("[writ] %s\n", s)
}

func (xio *SimpleXIO) ReadLine() (line string, err error) {
	xio.Log("XMPP is trying to read")
	return "", errors.New("Nothing to read from SimpleXIO")
}

func (xio *SimpleXIO) Destroy() {
	/* noop */
}

func (xio *SimpleXIO) Resize() {
	/* noop */
}

func (xio *SimpleXIO) SetAutoCompleteCallback(f xlib.XIOAutoCompleteCallbackI) {
	/* noop */
}

func NewSimpleXIO() (x *SimpleXIO) {
	return &SimpleXIO{}
}
