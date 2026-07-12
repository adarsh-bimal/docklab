package cmd

import (
	"fmt"
	"os/exec"

	"github.com/spf13/cobra"
)

// downCmd represents the down command
var downCmd = &cobra.Command{
	Use:   "down",
	Short: "Stop and remove the pentest-toolkit container",
	Long:  `This command stops and removes the pentest-toolkit container.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Stopping Docklab.......")
		containers := []string{"pentest-toolkit", "dvwa", "juice-shop", "webgoat"}
		for _, container := range containers {
			if err := removeContainer(container); err != nil {
				fmt.Printf("✗ Failed to remove %s\n", container)
			} else {
				fmt.Println("✓ Removed ", container)
			}
		}
		if err := removeNetwork(networkName); err != nil {
			fmt.Printf("✗ Failed to remove network '%s'\n", networkName)
		} else {
			fmt.Println("✓ Removed network ", networkName)
		}

		fmt.Println("✓ CyberLab stopped.")
	},
}

func removeContainer(name string) error {
	cmd := exec.Command("docker", "rm", "-f", name)
	return cmd.Run()
}

func removeNetwork(name string) error {
	cmd := exec.Command("docker", "network", "rm", name)
	return cmd.Run()
}

func init() {
	rootCmd.AddCommand(downCmd)
}
