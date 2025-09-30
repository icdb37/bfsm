// Package cmd 漏洞爬取命令
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// rootCmd 根命令
var rootCmd = &cobra.Command{
	Use:   "bfsm",
	Short: "a tool for bfsm",
}

// Execute 执行命令
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Errorf("init cmd failed, error: %w", err)
	}
}
