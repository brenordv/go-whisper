package main

import (
	"errors"
	"fmt"
	"github.com/brenordv/go-whisper/internal/chat"
	"github.com/brenordv/go-whisper/internal/core"
	"github.com/brenordv/go-whisper/internal/handlers"
	"github.com/brenordv/go-whisper/internal/security"
	"github.com/brenordv/go-whisper/internal/utils"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	var chatInfo *chat.Info
	var err error
	fmt.Printf("go.Whisper | v%s\n", core.Version)
	action := strings.ToLower(os.Args[1])

	privateKey, publicKey, err := security.GenerateKeyPair(core.EncKeySize)
	handlers.PanicOnError(err)

	myKeys := security.Key{
		Public:  publicKey,
		Private: privateKey,
	}

	if action == "wait" {
		waitPort := getWaitPort()
		chatInfo, err = chat.WaitForPartner(waitPort, myKeys)
		handlers.PanicOnError(err)
	}

	if action == "connect" {
		remoteHost := getRemoteHost()
		chatInfo, err = chat.ConnectToPartner(remoteHost, myKeys)
		handlers.PanicOnError(err)
	}

	defer func(chatInfo *chat.Info) {
		_ = chatInfo.CloseConnection()
	}(chatInfo)

	utils.ClearScreen()

	chatInfo.Session.Mine = myKeys
	chatInfo.Print(core.ConnectedMsg, false)

	go chat.ReadMessagesLoop(chatInfo)
	go chat.WriteMessagesLoop(chatInfo)

	handlers.WatchForUserInterruption(nil)
}

func getRemoteHost() string {
	remoteHost := os.Args[2]
	return remoteHost
}

func getWaitPort() int {
	if len(os.Args) != 3 {
		log.Panic(errors.New("wrong usage. correct: 'wait <port>'"))
	}
	waitPort, err := strconv.Atoi(os.Args[2])

	handlers.PanicOnError(err)
	return waitPort
}
