package slack

import (
	"log"
	"os"
	"testing"
)

func TestHandleHaltCallsDisconnectOnClient(t *testing.T) {
	disconnectCalled := false

	fakeClient := &fakeRtmClient{
		DisconnectHandler: func() error {
			disconnectCalled = true
			return nil
		},
	}

	bot := rtmBot{
		rtm:    fakeClient,
		logger: log.New(os.Stdout, "testmusicof-bot: ", log.Lshortfile|log.LstdFlags),
	}

	err := bot.handleHalt()

	if err != nil {
		t.Fatal("Expected no error, got ", err)
	}

	if !disconnectCalled {
		t.Error("Disconnect was not called")
	}
}
