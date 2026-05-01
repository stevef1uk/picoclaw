package agent

import (
	"testing"

	"github.com/sipeed/picoclaw/pkg/bus"
	"github.com/sipeed/picoclaw/pkg/config"
)

func TestAgentLoop_FreeRideRegistration(t *testing.T) {
	msgBus := bus.NewMessageBus()

	t.Run("freeride enabled", func(t *testing.T) {
		cfg := config.DefaultConfig()
		cfg.Tools.Freeride.Enabled = true
		
		al := NewAgentLoop(cfg, "config.json", msgBus, nil)
		agent := al.GetRegistry().GetDefaultAgent()
		
		if _, ok := agent.Tools.Get("freeride"); !ok {
			t.Errorf("Expected freeride tool to be registered when enabled")
		}
	})

	t.Run("freeride disabled", func(t *testing.T) {
		cfg := config.DefaultConfig()
		cfg.Tools.Freeride.Enabled = false
		
		al := NewAgentLoop(cfg, "config.json", msgBus, nil)
		agent := al.GetRegistry().GetDefaultAgent()
		
		if _, ok := agent.Tools.Get("freeride"); ok {
			t.Errorf("Expected freeride tool NOT to be registered when disabled")
		}
	})

	t.Run("freeride disabled but skills enabled", func(t *testing.T) {
		cfg := config.DefaultConfig()
		cfg.Tools.Freeride.Enabled = false
		cfg.Tools.Skills.Enabled = true
		
		al := NewAgentLoop(cfg, "config.json", msgBus, nil)
		agent := al.GetRegistry().GetDefaultAgent()
		
		if _, ok := agent.Tools.Get("freeride"); ok {
			t.Errorf("Expected freeride tool NOT to be registered even if skills are enabled")
		}
	})
}
