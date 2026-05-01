package config

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v3"
)

func TestBuiltinHookConfig_UnmarshalYAML(t *testing.T) {
	const yamlConfig = `
enabled: true
priority: 10
config:
  allowed_tools:
    cron: true
    exec: false
  requires_approval:
    delete_file: true
`
	var bhc BuiltinHookConfig
	err := yaml.Unmarshal([]byte(yamlConfig), &bhc)
	assert.NoError(t, err)
	assert.True(t, bhc.Enabled)
	assert.Equal(t, 10, bhc.Priority)

	// Verify that Config (RawNode) contains the correct JSON bytes
	var rawData map[string]any
	err = json.Unmarshal(bhc.Config, &rawData)
	assert.NoError(t, err)

	allowedTools, ok := rawData["allowed_tools"].(map[string]any)
	assert.True(t, ok)
	assert.Equal(t, true, allowedTools["cron"])
	assert.Equal(t, false, allowedTools["exec"])

	requiresApproval, ok := rawData["requires_approval"].(map[string]any)
	assert.True(t, ok)
	assert.Equal(t, true, requiresApproval["delete_file"])
}

func TestBuiltinHookConfig_MergeYAML(t *testing.T) {
	// Start with some existing config
	bhc := BuiltinHookConfig{
		Enabled:  false,
		Priority: 5,
		Config:   RawNode(`{"allowed_tools":{"old":true}}`),
	}

	const yamlOverlay = `
enabled: true
config:
  allowed_tools:
    cron: true
`
	err := yaml.Unmarshal([]byte(yamlOverlay), &bhc)
	assert.NoError(t, err)
	assert.True(t, bhc.Enabled)
	assert.Equal(t, 5, bhc.Priority) // Priority should remain 5 as it's not in overlay

	var rawData map[string]any
	err = json.Unmarshal(bhc.Config, &rawData)
	assert.NoError(t, err)

	allowedTools, ok := rawData["allowed_tools"].(map[string]any)
	assert.True(t, ok)
	assert.Equal(t, true, allowedTools["old"])  // Should be preserved
	assert.Equal(t, true, allowedTools["cron"]) // Should be added
}

func TestAgentDefaults_ContextManagerConfig_UnmarshalYAML(t *testing.T) {
	const yamlConfig = `
context_manager: memory
context_manager_config:
  max_tokens: 4096
  strategy: summarize
`
	var defaults AgentDefaults
	err := yaml.Unmarshal([]byte(yamlConfig), &defaults)
	assert.NoError(t, err)
	assert.Equal(t, "memory", defaults.ContextManager)

	var rawData map[string]any
	err = json.Unmarshal(defaults.ContextManagerConfig, &rawData)
	assert.NoError(t, err)
	assert.Equal(t, float64(4096), rawData["max_tokens"]) // JSON unmarshals numbers to float64
	assert.Equal(t, "summarize", rawData["strategy"])
}

func TestToolsConfig_UnmarshalYAML(t *testing.T) {
	const yamlConfig = `
freeride:
  enabled: true
cron:
  enabled: true
  allow_command: true
web:
  enabled: true
  brave:
    enabled: true
    max_results: 10
`
	var tc ToolsConfig
	err := yaml.Unmarshal([]byte(yamlConfig), &tc)
	assert.NoError(t, err)

	// Previously these would have been false because of yaml:"-"
	assert.True(t, tc.Freeride.Enabled, "Freeride should be enabled via YAML")
	assert.True(t, tc.Cron.Enabled, "Cron should be enabled via YAML")
	assert.True(t, tc.Cron.AllowCommand)
	assert.True(t, tc.Web.Enabled)
	assert.True(t, tc.Web.Brave.Enabled)
	assert.Equal(t, 10, tc.Web.Brave.MaxResults)
}
