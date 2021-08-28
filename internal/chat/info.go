package chat

import (
	"fmt"
	"github.com/brenordv/go-whisper/internal/security"
	"net"
	"os"
	"strings"
	"time"
)

type Info struct {
	Conn           net.Conn
	SessionStarted time.Time
	Session     security.Session
}

func (i *Info) printFooter() {
	infoMessage := fmt.Sprintf("disconnected. (duration: %s)", time.Since(i.SessionStarted))
	i.Print(infoMessage, false)
}

func (i *Info) CloseConnection() error {
	defer i.printFooter()
	if i.Conn == nil {
		return nil
	}

	err := i.Conn.Close()
	return err
}

func (i *Info) Send(msg string) error {
	msg = strings.TrimSpace(msg)
	if msg == "" {
		return nil
	}

	return i.SendBytes([]byte(msg))
}

func (i *Info) SendBytes(msg []byte) error {
	_, err := i.Conn.Write(msg)
	return err
}

func (i *Info) Print(msg string, isReceived bool) {
	var prefix string
	if isReceived {
		prefix = "< "
	}
	fmt.Printf("%s%s\n", prefix, msg)
}

func (i *Info) Quit() {
	_ = i.CloseConnection()
	os.Exit(1)
}

func (i Info) ProcessMsgCmds(cmdMsg string) (bool, error) {
	if !strings.HasPrefix(cmdMsg, "/") {
		return false, nil
	}

	cmdParts := strings.Split(cmdMsg, " ")
	cmd := strings.ToLower(cmdParts[0])

	switch cmd {
	case "/quit":
		i.Quit()
		return true, nil

	default:
		return true, fmt.Errorf("command '%s' is not supported", cmd)
	}
}