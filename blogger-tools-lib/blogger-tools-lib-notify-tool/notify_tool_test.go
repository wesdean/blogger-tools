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
		results, err := tool.Run(&blogger_tools_lib_notify_tool.NotifyToolArgs{ResetLog: true})
		if err != nil {
			t.Error(err)
		}

		if results.Success != true {
			t.Errorf("expected true, got %v", results.Success)
			return
		}
	})

	t.Run("Run NotifyTool with BlogUpdatedAction", func(t *testing.T) {
		toolOptions := &blogger_tools_lib_notify_tool.NotifyToolArgs{
			ResetLog: true,
			Actions: &blogger_tools_lib_notify_tool.NotifyToolActions{
				&blogger_tools_lib_notify_tool.ActionBlogUpdatedOptions{},
			},
		}
		results, err := tool.Run(toolOptions)
		if err != nil {
			t.Error(err)
		}

		if results.Success != true {
			t.Errorf("expected true, got %v", results.Success)
			return
		}

		if results.Actions.BlogUpdated != true {
			t.Errorf("expected true, got %v", results.Actions.BlogUpdated)
			return
		}
	})
}
