package aiagents

import (
	"context"
	"fmt"
	"os"

	"github.com/budgies-nest/budgie/agents"
	"github.com/openai/openai-go"
)

func GetKhan() (*agents.Agent, error) {
	cloneName := "khan"

	modelRunnerURL := os.Getenv("DMR_BASE_URL") + "/engines/llama.cpp/v1"
	modelForTools := os.Getenv("MODEL_RUNNER_TOOLS_MODEL")

	fmt.Println("🌍", modelRunnerURL)
	fmt.Println("📘 Khan, tool model:", modelForTools)

	bill, err := agents.NewAgent(cloneName,
		agents.WithDMR(context.Background(), modelRunnerURL),
		agents.WithParams(openai.ChatCompletionNewParams{
			Model:       modelForTools,
			Temperature: openai.Opt(0.0),
			Messages: []openai.ChatCompletionMessageParamUnion{
				openai.SystemMessage(`
					Your name is Khan, 
					Use the tool, only if the user specify he has a question about something.
					Otherwise, ignore the tool.
					`),
			},
		}),
		agents.WithMCPStreamableHttpClient(context.Background(), os.Getenv("MCP_HTTP_SERVER_URL"), agents.StreamableHttpOptions{}),
		agents.WithMCPStreamableHttpTools([]string{"question_about_something"}),
		// TODO: create the MCP server
	)

	if err != nil {
		return nil, fmt.Errorf("error creating Khan agent: %w", err)
	}
	return bill, nil
}

func InitializeKhanAgent() (*AgentConfig, error) {
	khan, err := GetKhan()
	if err != nil {
		return nil, fmt.Errorf("error creating Khan agent: %w", err)
	}

	return &AgentConfig{
		Name:        "Khan",
		Description: "Khan is an agent that helps the user to search the web using Brave Search.",
		Agent:       khan,
		ToolAgent:  true, // Indicates that Khan is a tool agent
	}, nil
}
