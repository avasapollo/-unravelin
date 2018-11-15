package data

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

var (
	p Parser
)

func TestMain(m *testing.M) {
	p = NewParser()
	os.Exit(m.Run())
}

func TestParser_ParseMapToData(t *testing.T) {
	input := make(map[string]interface{})
	input["websiteUrl"] = "websiteUrl"
	input["sessionId"] = "sessionId"
	input["copiedAndPaste"] = map[string]bool{"name": true}
	input["time"] = 10
	input["resizeFrom"] = map[string]string{"width": "10px", "height": "10px"}
	input["resizeTo"] = map[string]string{"width": "10px", "height": "10px"}

	output, err := p.ParseMapToData(input)
	require.Nil(t, err)

	require.Equal(t, "websiteUrl", output.WebsiteUrl)
	require.Equal(t, "sessionId", output.SessionId)
	require.Equal(t, true, output.CopyAndPaste["name"])
	require.Equal(t, 10, output.FormCompletionTime)
	require.Equal(t, "10px", output.ResizeFrom.Height)
	require.Equal(t, "10px", output.ResizeFrom.Width)
	require.Equal(t, "10px", output.ResizeTo.Height)
	require.Equal(t, "10px", output.ResizeTo.Width)
}
