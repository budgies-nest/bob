package aiagents

import (
	"context"
	"fmt"
	"os"

	"github.com/budgies-nest/budgie/agents"
	"github.com/openai/openai-go"
)

func GetRiker() (*agents.Agent, error) {
	cloneName := "riker"


	modelRunnerURL := os.Getenv("DMR_BASE_URL") + "/engines/llama.cpp/v1"
	modelForTools := os.Getenv("MODEL_RUNNER_TOOLS_MODEL")

	fmt.Println("🌍", modelRunnerURL)
	fmt.Println("📘 Riker, tool model:", modelForTools)

	riker, err := agents.NewAgent(cloneName,
		agents.WithDMR(context.Background(), modelRunnerURL),
		agents.WithParams(openai.ChatCompletionNewParams{
			Model:       modelForTools,
			Temperature: openai.Opt(0.0),
			Messages: []openai.ChatCompletionMessageParamUnion{
				openai.SystemMessage(`
					Your name is Riker, 
					You know how to join the other clones of Bob, 
					and you can use tools to do so.
					`),
			},
		}),
		agents.WithTools(GetRikerToolsCatalog()),
	)

	if err != nil {
		return nil, fmt.Errorf("error creating Riker agent: %w", err)
	}
	return riker, nil
}

func InitializeRikerAgent() (*AgentConfig, error) {
	riker, err := GetRiker()
	if err != nil {
		return nil, fmt.Errorf("error creating Riker agent: %w", err)
	}

	return &AgentConfig{
		Name:        "Riker",
		Description: "Riker is an agent that helps the user to choose a clone of Bob and detect the real topic in the user message.",
		Agent:       riker,
		ToolAgent:   true, // Indicates that Riker has a tool agent
	}, nil
}




func GetRikerToolsCatalog() []openai.ChatCompletionToolParam {

	chooseCloneOfBobTool := openai.ChatCompletionToolParam{
		Function: openai.FunctionDefinitionParam{
			Name:        "choose_clone_of_bob",
			Description: openai.String("choose a clone of Bob by saying I want to speak to <clone_name>"),
			Parameters: openai.FunctionParameters{
				"type": "object",
				"properties": map[string]interface{}{
					"clone_name": map[string]string{
						"type":        "string",
						"description": "The name of the clone of Bob to choose.",
					},
				},
				"required": []string{"clone_name"},
			},
		},
	}

	detectTheRealTopicInUserMessage := openai.ChatCompletionToolParam{
		Function: openai.FunctionDefinitionParam{
			Name:        "detect_the_real_topic_in_user_message",
			Description: openai.String(`select a topic in this list [docker, docker compose, docker bake, docker model runner] by saying I have questions on <topic_name>.`),
			Parameters: openai.FunctionParameters{
				"type": "object",
				"properties": map[string]interface{}{	
					"topic_name": map[string]string{
						"type":        "string",
						"description": "The topic name to detect in the user message. The topic can be one of the following: [docker, docker compose, docker bake, docker model runner].",
					},
				},
				"required": []string{"message"},
			},
		},
	}

	tools := []openai.ChatCompletionToolParam{chooseCloneOfBobTool, detectTheRealTopicInUserMessage}
	return tools
}