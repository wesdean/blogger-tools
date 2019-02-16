package blogger_tools_lib_test

import (
	"github.com/wesdean/blogger-tools/blogger-tools-lib"
	"testing"
)

var notifyToolLogFile = "../logs/notify_tool.log"

func TestNotifyTool_Run(t *testing.T) {
	var args = &blogger_tools_lib.NotifyToolArgs{
		Config:      getConfig(),
		LogFilePath: notifyToolLogFile,
	}

	var tool = blogger_tools_lib.NewNotifyTool()
	err := tool.Run(args)
	if err != nil {
		t.Error(err)
	}
}
