package chat

import (
	"fmt"
	"github.com/brenordv/go-whisper/internal/network"
	"github.com/brenordv/go-whisper/internal/security"
	"net"
	"time"
)

func WaitForPartner(waitPort int, sec security.Key) (*Info, error) {
	extIp, err := network.GetExternalIp()
	if err != nil {
		return nil, err
	}
	fmt.Printf("waiting for connection. external ip: %s:%d\n", extIp, waitPort)

	listenTo := fmt.Sprintf(":%d", waitPort)
	listener, err := net.Listen("tcp4", listenTo)
	if err != nil {
		return nil, err
	}

	defer func(listener net.Listener) {
		_ = listener.Close()
	}(listener)

	conn, sessionStart := network.WaitForConnection(listener)

	chatInfo := Info{
		Conn:           conn,
		SessionStarted: sessionStart,
		Session: security.Session{
			Mine: sec,
		},
	}
	return sendPublicKey(sec, chatInfo)
}

func ConnectToPartner(remoteHost string, sec security.Key) (*Info, error) {
	conn, err := net.Dial("tcp", remoteHost)
	if err != nil {
		return nil, err
	}

	chatInfo := Info{
		Conn:           conn,
		SessionStarted: time.Now(),
		Session: security.Session{
			Mine: sec,
		},
	}

	return sendPublicKey(sec, chatInfo)
}

//todo: think of a better method name
func sendPublicKey(sec security.Key, chatInfo Info) (*Info, error) {
	var err error
	var pubKeyBytes []byte
	pubKeyBytes, err = sec.PublicKeyToBytes()
	if err != nil {
		return nil, err
	}

	err = chatInfo.SendBytes(pubKeyBytes)
	if err != nil {
		return nil, err
	}

	return &chatInfo, nil
}
