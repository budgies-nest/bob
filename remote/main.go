package main

import (
	"context"
	"fmt"

	"github.com/budgies-nest/budgie/agents"
	"github.com/budgies-nest/budgie/enums/base"
	"github.com/budgies-nest/budgie/rag"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/openai/openai-go"
)


func main() {

	contents, err := rag.GetContentFiles("/app/docs", ".md")
	if err != nil {
		panic(fmt.Errorf("error getting content files for Bob agent: %w", err))
	}

	chunks := []string{}
	for _, content := range contents {
		chunks = append(chunks, rag.ChunkText(content, 512, 210)...)
	}

	bob, err := agents.NewAgent("Bob",
		agents.WithDMR(context.Background(), base.DockerModelRunnerContainerURL),
		agents.WithEmbeddingParams(
			openai.EmbeddingNewParams{
				Model: "ai/mxbai-embed-large",
			},
		),
		agents.WithRAGMemory(chunks),
		agents.WithMCPStreamableHttpServer(agents.MCPServerConfig{
			Port:     "7070",
			Version:  "v1",
			Name:     "mcp-bob",
			Endpoint: "/mcp",
		}),
	)

	if err != nil {
		panic(err)
	}

	searchInDoc := mcp.NewTool("question_about_pizza",
		mcp.WithDescription(`Find an information about hawaiian pizza.`),
		mcp.WithString("question",
			mcp.Required(),
			mcp.Description("Search question about hawaiian pizza."),
		),
	)

	bob.AddToolToMCPServer(searchInDoc, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()
		question := args["question"].(string)
		similarities, err := bob.RAGMemorySearchSimilaritiesWithText(question, 0.7)
		if err != nil {
			return nil, err
		}
		content := ""
		for _, similarity := range similarities {
			content += similarity + "\n"
		}
		return mcp.NewToolResultText(content), nil
	})

	fmt.Println("MCP Streamable HTTP server Agent is running on port 9090")
	bob.StartMCPHttpServer()

}
