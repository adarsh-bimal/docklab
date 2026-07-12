package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var shellCmd = &cobra.Command{
	Use:   "shell",
	Short: "Open an interactive shell in the Pentest Toolkit",
	Long:  "Drops you into the Pentest Toolkit container.",

	Run: func(cmd *cobra.Command, args []string) {
		if err := openShell(); err != nil {
			fmt.Println("Error:", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(shellCmd)
}
