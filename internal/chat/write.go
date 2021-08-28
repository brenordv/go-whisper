package chat

import (
	"bufio"
	"github.com/brenordv/go-whisper/internal/handlers"
	"github.com/brenordv/go-whisper/internal/network"
	"os"
	"strings"
	"time"
)

func WriteMessagesLoop(c *Info) {
	reader := bufio.NewReader(os.Stdin)
	var err error
	var msg string

	for {
		msg, err = reader.ReadString('\n')
		isDisconnected := network.IsConnectionClosed(err)
		if isDisconnected {
			c.Quit()
		}
		handlers.PanicOnError(err)

		msg = strings.TrimSpace(msg)
		if msg == "" {
			time.Sleep(500 * time.Millisecond)
			continue
		}

		wasCmd, cmdErr := c.ProcessMsgCmds(msg)
		if wasCmd {
			handlers.PanicOnError(cmdErr)
			continue
		}

		payload, err := c.Session.Partner.Encrypt([]byte(msg))
		handlers.PanicOnError(err)

		err = c.SendBytes(payload)

		isDisconnected = network.IsConnectionClosed(err)
		if isDisconnected {
			c.Quit()
		}
		handlers.PanicOnError(err)
	}
}
