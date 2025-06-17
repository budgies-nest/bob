package aiagents

import "github.com/budgies-nest/budgie/agents"

type AgentConfig struct {
	Name        string        `json:"name"`
	Description string        `json:"description"`
	Agent       *agents.Agent `json:"agent"`
	ToolAgent   bool          `json:"tool_agent,omitempty"` // Indicates if the agent has a tool agent
}
