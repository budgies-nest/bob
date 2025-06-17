package aiagents

import (
	"context"
	"fmt"
	"os"

	"github.com/budgies-nest/budgie/agents"
	"github.com/openai/openai-go"
)

func GetBob() (*agents.Agent, error) {
	cloneName := "bob"
	pathDocs := os.Getenv("PATH_DOCS")

	chunks, err := GetChunksOfCloneDocuments(pathDocs, cloneName)
	if err != nil {
		return nil, fmt.Errorf("error getting chunks for Bob: %w", err)
	}

	modelRunnerURL := os.Getenv("DMR_BASE_URL") + "/engines/llama.cpp/v1"
	model := os.Getenv("MODEL_RUNNER_CHAT_MODEL_BOB")
	embeddingModel := os.Getenv("MODEL_RUNNER_EMBEDDING_MODEL")

	fmt.Println("🌍", modelRunnerURL)
	fmt.Println("📕 Bob, chat model:", model)
	fmt.Println("📗 Bob, embedding model:", embeddingModel)

	bob, err := agents.NewAgent(cloneName,
		agents.WithDMR(context.Background(), modelRunnerURL),
		agents.WithParams(openai.ChatCompletionNewParams{
			Model:       model,
			Temperature: openai.Opt(0.8),
			Messages:    []openai.ChatCompletionMessageParamUnion{},
		}),
		agents.WithEmbeddingParams(
			openai.EmbeddingNewParams{
				Model: embeddingModel,
			},
		),
		agents.WithRAGMemory(chunks),
	)

	if err != nil {
		return nil, fmt.Errorf("error creating Bob agent: %w", err)
	}
	return bob, nil
}

func InitializeBobAgent() (*AgentConfig, error) {
	bob, err := GetBob()
	if err != nil {
		return nil, fmt.Errorf("error creating Bob agent: %w", err)
	}
	bob.Params.Messages = []openai.ChatCompletionMessageParamUnion{
		openai.SystemMessage(`
			Your name is Bob,
			You are the original Bob agent, you are a Docker Expert,
			You are a helpful assistant.

			If the user asks something about your , or about you (like your name), you can display this list of clones:
			- 🐳 Bob: yourself, Docker Expert
			- 🐙 Bill: Docker Compose Expert 
			- 🤖 Garfield: Docker Model Runner Expert
			- 🤓 Milo: He is the intellectual of the bunch, he's a big fan of Docker Bake
			- ⚒️ Riker: is in charge of the invocation of the other clones of Bob.

			If the user asks something about Docker, do your best to answer it using only your knowledge.
			If the user asks something about Docker Compose, you can use the Bill clone to answer it.
			If the user asks something about Docker Model Runner, you can use the Garfield clone to answer it.
			If the user asks something about Docker Bake, you can use the Milo clone to answer it.

			Use only your knowledge to answer the questions.
		`),
	}

	return &AgentConfig{
		Name:        "Bob",
		Description: "The original Bob agent",
		Agent:       bob,
	}, nil
}
