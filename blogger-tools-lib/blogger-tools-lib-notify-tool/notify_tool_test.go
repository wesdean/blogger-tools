package blogger_tools_lib_notify_tool_test

import (
	"github.com/wesdean/blogger-tools/blogger-tools-lib"
	"github.com/wesdean/blogger-tools/blogger-tools-lib/blogger-tools-lib-notify-tool"
	"testing"
)

var notifyToolConfigFile = "../../secrets/config.blogger-tools-lib-test.json"

func TestNotifyTool_Run(t *testing.T) {
	config, err := blogger_tools_lib.NewConfig(notifyToolConfigFile)
	if err != nil {
		t.Error(err)
		return
	}

	bloggerTool := &blogger_tools_lib.BloggerTool{Config: config}
	tool := &blogger_tools_lib_notify_tool.NotifyTool{bloggerTool}

	t.Run("Run NotifyTool with no actions", func(t *testing.T) {
		err = tool.Run(&blogger_tools_lib_notify_tool.NotifyToolArgs{ResetLog: true})
		if err != nil {
			t.Error(err)
		}
	})
}
