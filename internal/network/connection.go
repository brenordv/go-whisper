package network

import (
	"io"
	"net"
	"strings"
	"time"
)

func IsConnectionClosed(e error) bool {
	if e == nil {
		return false
	}

	if e == io.EOF {
		return true
	}

	errMsg := strings.ToLower(e.Error())
	if !strings.Contains(errMsg, "wsarecv") {
		return false
	}

	opErr, ok := e.(*net.OpError)
	if !ok || (opErr.Timeout() && opErr.Temporary()) {
		return false
	}

	return true
}

func WaitForConnection(listener net.Listener) (net.Conn, time.Time) {
	var err error
	var conn net.Conn

	for {
		if conn, err = listener.Accept(); err == nil {
			return conn, time.Now()
		}
	}
}