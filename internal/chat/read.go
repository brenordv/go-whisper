package chat

import (
	"github.com/brenordv/go-whisper/internal/core"
	"github.com/brenordv/go-whisper/internal/handlers"
	"github.com/brenordv/go-whisper/internal/network"
	"github.com/brenordv/go-whisper/internal/utils"
	"time"
)

func ReadMessagesLoop(i *Info) {
	data := make([]byte, core.MaxMsgSize)

	for {
		length, err := i.Conn.Read(data)
		isDisconnected := network.IsConnectionClosed(err)
		if isDisconnected {
			i.Quit()
		}
		handlers.PanicOnError(err)

		if length == 0 {
			time.Sleep(500 * time.Millisecond)
			continue
		}

		data = data[:length]

		if !core.HandShakeComplete {
			err = i.Session.Partner.BytesToPublicKey(data)
			handlers.PanicOnError(err)
			core.HandShakeComplete = true
			continue
		}

		var msgBytes []byte
		msgBytes, err = i.Session.Mine.Decrypt(data)
		handlers.PanicOnError(err)
		msg := utils.ByteToString(msgBytes, len(msgBytes))

		i.Print(msg, true)
	}
}