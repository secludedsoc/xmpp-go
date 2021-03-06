package xlib

type XIOAutoCompleteCallbackI func(line string, pos int, key rune) (string, int, bool)

type XIO interface {
	Info(msg string)
	Warn(msg string)
	Debug(format string, a ...interface{})
	Alert(msg string)
	Critical(msg string)
	ReadPassword(msg string) (password string, err error)
	SetPrompt(prompt string)
	SetPromptEnc(target string, isEncrypted bool)
	Message(timestamp, from, fromres, to string, msg []byte, isEncrypted bool, doBell bool, delayed bool)
	StatusUpdate(timestamp, from, to, show string, status string, gone bool)
	FormStringForPrinting(s string) string
	Write(s string)
	ReadLine() (line string, err error)
	SetAutoCompleteCallback(f XIOAutoCompleteCallbackI)
	Resize()
	Destroy()
}
