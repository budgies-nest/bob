package workflow

import (
	"fmt"
	"net/http"
	"we-are-legion/helpers"

	budgiesHelpers "github.com/budgies-nest/budgie/helpers"

	"github.com/budgies-nest/budgie/agents"
	"github.com/openai/openai-go"
)

func DetectToolCalls(response http.ResponseWriter, flusher http.Flusher, riker *agents.Agent) ([]openai.ChatCompletionMessageToolCall, error) {
	toolCalls, err := riker.ToolsCompletion()
	if err != nil {
		if len(toolCalls) > 0 {
			fmt.Println("😡 Error: ", err.Error())
			helpers.ResponseLabel(response, flusher, "error", "Tool call error detected: "+err.Error())
		} else {
			fmt.Println("🙂 no tool calls detected.")
			helpers.ResponseLabel(response, flusher, "success", "No tool calls detected")
		}
	}
	fmt.Println("🤖 Number of Tool Calls:", len(toolCalls))
	if len(toolCalls) > 0 {
		toolCallsJSON, _ := budgiesHelpers.ToolCallsToJSONString(toolCalls)
		fmt.Println("🤖 Tool Calls:\n", toolCallsJSON)
	}
	return toolCalls, err
}

func DetectMCPToolCalls(response http.ResponseWriter, flusher http.Flusher, khan *agents.Agent) ([]openai.ChatCompletionMessageToolCall, error) {
	mcpTooCalls, err := khan.ToolsCompletion()
	if err != nil {
		if len(mcpTooCalls) > 0 {
			fmt.Println("😡 Error: ", err.Error())
			helpers.ResponseLabel(response, flusher, "error", "MCP Tool call error detected: "+err.Error())
		} else {
			fmt.Println("🙂 no tool calls detected.")
			helpers.ResponseLabel(response, flusher, "success", "No MCP tool calls detected")
		}
	}
	fmt.Println("🤖 Number of MCP Tool Calls:", len(mcpTooCalls))
	if len(mcpTooCalls) > 0 {
		mcpToolCallsJSON, _ := budgiesHelpers.ToolCallsToJSONString(mcpTooCalls)
		fmt.Println("🤖 MCP Tool Calls:\n", mcpToolCallsJSON)
	}
	return mcpTooCalls, err
}
