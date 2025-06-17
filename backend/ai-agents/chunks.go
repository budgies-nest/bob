package aiagents

import (
	"fmt"
	"github.com/budgies-nest/budgie/rag"
)

func GetChunksOfCloneDocuments(pathDocs, cloneName string) ([]string, error) {
	contents, err := rag.GetContentFiles(pathDocs+"/"+cloneName, ".md")
	if err != nil {
		return nil, fmt.Errorf("error getting content files for %s agent: %w", cloneName, err)
	}
	for _, content := range contents {
		fmt.Println("📄", cloneName, "content file:", content)
	}
	chunks := []string{}
	for _, content := range contents {
		chunks = append(chunks, rag.ChunkText(content, 512, 210)...)
	}
	return chunks, nil
}
