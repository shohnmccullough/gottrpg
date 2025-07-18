package tests

import (
	"os"
	"testing"

	. "example.com/gottrpg/v2/log"
	"github.com/sirupsen/logrus"
)

func TestLogging(t *testing.T) {
	NewLogger().InvalidArgument("dice")
}

func TestLoggingToFile(t *testing.T) {
	const debug string = "LOGGING_PATH"
	if err := os.Setenv(debug, "./log.txt"); err != nil {
		t.Fail()
	}

	logger := NewLogger()

	logger.InvalidArgument("dice")
	logger.Log(Log{Id: Info, Message: "%s: %s"}, "status", "nominal")

	if err := os.Unsetenv(debug); err != nil {
		t.Fail()
	}
}

func TestLoggingWithDefaultFields(t *testing.T) {
	const id string = "zDfR54Tszsio09"

	logger := NewLogger().NewLoggerEntry(logrus.Fields{"id": id})

	logger.InvalidArgument("snacks")
}

func TestRemoveLogFile(t *testing.T) {
	if err := os.Remove("./log.txt"); err != nil {
		t.Fail()
	}
}
