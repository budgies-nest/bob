package workflow

import (
	"net/http"
	"we-are-legion/helpers"

	"github.com/budgies-nest/budgie/agents"
	"github.com/openai/openai-go"
)

func ExecuteMCPToolCalls(response http.ResponseWriter, flusher http.Flusher, khan *agents.Agent, mcpTooCalls []openai.ChatCompletionMessageToolCall) ([]string, error) {
	helpers.ResponseLabel(response, flusher, "orange", "Executing MCP tool calls...")
	mcpResults, err := khan.ExecuteMCPStreamableHTTPToolCalls(mcpTooCalls)
	if err != nil {
		helpers.ResponseLabel(response, flusher, "error", "MCP Tool execution failed: "+err.Error())
	} else {
		helpers.ResponseLabel(response, flusher, "success", "MCP Tool calls executed successfully")
	}
	// Reset the messages for Khan
	
	khan.Params.Messages = []openai.ChatCompletionMessageParamUnion{
		openai.SystemMessage(`
					Your name is Khan, 
					Use the tool, only if the user specify he has a question about something.
					Otherwise, ignore the tool.
					`),
	}
	return mcpResults, err
}
