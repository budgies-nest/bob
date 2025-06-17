package workflow

import "we-are-legion/ai-agents"

func InitializeAgents() map[string]*aiagents.AgentConfig {
	// create a map of agents
	agentsCatalog := map[string]*aiagents.AgentConfig{
		"bob": func() *aiagents.AgentConfig {
			cfg, err := aiagents.InitializeBobAgent()
			if err != nil {
				panic("Error initializing Bob agent: " + err.Error())
			}
			return cfg
		}(),
		"bill": func() *aiagents.AgentConfig {
			cfg, err := aiagents.InitializeBillAgent()
			if err != nil {
				panic("Error initializing Bill agent: " + err.Error())
			}
			return cfg
		}(),
		"milo": func() *aiagents.AgentConfig {
			cfg, err := aiagents.InitializeMiloAgent()
			if err != nil {
				panic("Error initializing Milo agent: " + err.Error())
			}
			return cfg
		}(),
		"garfield": func() *aiagents.AgentConfig {
			cfg, err := aiagents.InitializeGarfieldAgent()
			if err != nil {
				panic("Error initializing Garfield agent: " + err.Error())
			}
			return cfg
		}(),
		"riker": func() *aiagents.AgentConfig {
			cfg, err := aiagents.InitializeRikerAgent()
			if err != nil {
				panic("Error initializing Riker agent: " + err.Error())
			}
			return cfg
		}(),
		"khan": func() *aiagents.AgentConfig {
			cfg, err := aiagents.InitializeKhanAgent()
			if err != nil {
				panic("Error initializing Khan agent: " + err.Error())
			}
			return cfg
		}(),
	}
	return agentsCatalog

}
