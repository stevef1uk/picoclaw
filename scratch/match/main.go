package main

import (
	"fmt"
	"strings"
)

func main() {
	tool := "mcp_hdn-server_weather"
	w := "hdn-server"
	match := strings.HasPrefix(tool, "mcp_"+w+"_") ||
		strings.HasPrefix(tool, "tool_"+w+"_") ||
		strings.HasPrefix(tool, w+"_")
	fmt.Printf("Match: %v\n", match)
}
