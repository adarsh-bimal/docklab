/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bytes"
	"fmt"

	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

func containerExists(name string) bool {
	cmd := exec.Command("docker", "inspect", name)

	var stderr bytes.Buffer
	cmd.Stderr = &stderr

	err := cmd.Run()
	return err == nil
}

func containerRunning(name string) bool {
	cmd := exec.Command(
		"docker",
		"inspect",
		"-f",
		"{{.State.Running}}",
		name,
	)

	out, err := cmd.Output()
	if err != nil {
		return false
	}

	return string(bytes.TrimSpace(out)) == "true"
}

func startContainer(name string) error {
	cmd := exec.Command("docker", "start", name)
	return cmd.Run()
}

/*
	func createPentestToolkit() error {
		cmd := exec.Command(
			"docker",
			"run",
			"-dit",
			"--name",
			"pentest-toolkit",
			"--network",
			"cyberlab-network",
			"your-kali-image",
		)

		return cmd.Run()
	}
*/
func createContainer(name, image, network string, ports []string) error {
	args := []string{
		"run",
		"-d", // detached
		"--name", name,
		"--network", network,
		"--cap-add", "NET_RAW",
		"--cap-add", "NET_ADMIN",
	}

	// Add port mappings
	for _, port := range ports {
		args = append(args, "-p", port)
	}

	// Add the image name
	args = append(args, image)

	cmd := exec.Command("docker", args...)

	// Hide Docker's output
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}

func ensureContainer(name string, image string, network string, ports []string) error {
	if containerExists(name) {
		if !containerRunning(name) {
			return startContainer(name)
		}
		fmt.Printf("✓ %s already running\n", name)
		return nil
	}

	if containerRunning(name) {
		fmt.Printf("✓ %s already running\n", name)
		return nil
	}

	fmt.Printf("Starting %s...\n", name)

	if err := createContainer(name, image, network, ports); err != nil {
		return err
	}

	return nil

}

func openShell() error {
	fmt.Println("Putting you on the attacker's seat")
	seat := exec.Command(
		"docker",
		"exec",
		"-it",
		"pentest-toolkit",
		"/bin/bash",
	)

	seat.Stdin = os.Stdin
	seat.Stdout = os.Stdout
	seat.Stderr = os.Stderr

	return seat.Run()

}

var networkName string
var driverName string
var shell bool

// upCmd represents the up command
var upCmd = &cobra.Command{
	Use:   "up",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {

		fmt.Println("Checking Docker.....")
		_, err := exec.LookPath("docker")
		if err != nil {
			fmt.Println("Docker is not installed. Please install Docker and try again.")
			return
		}

		// Check if the Docker network exists
		comd := exec.Command("docker", "network", "inspect", networkName)
		var out bytes.Buffer
		comd.Stdout = &out
		err = comd.Run()
		if err != nil {
			fmt.Printf("Docker network '%s' does not exist.\n", networkName)

			fmt.Printf("Preparing to create network '%s' using driver '%s'...\n", networkName, driverName)

			comd = exec.Command("docker", "network", "create", "--driver", driverName, networkName)

			var stderr bytes.Buffer
			comd.Stderr = &stderr

			err = comd.Run()
			if err != nil {
				fmt.Printf("Error creating Docker network: %s\n", stderr.String())
				return
			}
			fmt.Printf("Docker network '%s' created successfully.\n", networkName)

		} else {
			fmt.Printf("Docker network '%s' already exists.\n", networkName)

		}

		ensureContainer(
			"pentest-toolkit",
			"adarshbimal/docklab-toolkit:latest",
			networkName,
			nil,
		)

		ensureContainer(
			"dvwa",
			"vulnerables/web-dvwa",
			networkName,
			[]string{"8080:80"},
		)

		/*ensureContainer(
			"metasploitable",
			"tleemcjr/metasploitable2",
			networkName,
			[]string{"443:443"},
		)*/

		ensureContainer(
			"juice-shop",
			"bkimminich/juice-shop",
			networkName,
			[]string{"3000:3000"},
		)

		ensureContainer(
			"webgoat",
			"webgoat/webgoat",
			networkName,
			[]string{"8081:8080"},
		)
		fmt.Println("-----------------------------------------------------")
		fmt.Println("Attacker:")

		ip_pentest := exec.Command(
			"docker",
			"inspect",
			"-f",
			"{{range.NetworkSettings.Networks}}{{.IPAddress}}{{end}}",
			"pentest-toolkit",
		)
		ipout, err := ip_pentest.Output()
		if err != nil {
			fmt.Printf("Error getting IP address: %v\n", err)
			return
		}

		fmt.Printf("Pentest Toolkit IP: %s\n", string(ipout))

		fmt.Println("Target:")

		ip_dvwa := exec.Command(
			"docker",
			"inspect",
			"-f",
			"{{range.NetworkSettings.Networks}}{{.IPAddress}}{{end}}",
			"dvwa",
		)

		ipout, err = ip_dvwa.Output()
		if err != nil {
			fmt.Printf("Error getting IP address: %s\n", err)
			return
		}
		fmt.Println("DVWA IP Address:" + string(ipout))

		ip_juice := exec.Command(
			"docker",
			"inspect",
			"-f",
			"{{range.NetworkSettings.Networks}}{{.IPAddress}}{{end}}",
			"juice-shop",
		)

		ipout, err = ip_juice.Output()
		if err != nil {
			fmt.Printf("Error getting IP address: %s\n", err)
			return
		}
		fmt.Println("Juice Shop IP Address:" + string(ipout))

		ip_webgoat := exec.Command(
			"docker",
			"inspect",
			"-f",
			"{{range.NetworkSettings.Networks}}{{.IPAddress}}{{end}}",
			"webgoat",
		)

		ipout, err = ip_webgoat.Output()
		if err != nil {
			fmt.Printf("Error getting IP address: %s\n", err)
			return
		}
		fmt.Println("WebGoat IP Address:" + string(ipout))

		if shell {
			if err := openShell(); err != nil {
				fmt.Printf("Error opening shell: %v\n", err)
				return
			}
		}

		/*ip_metasploitable := exec.Command(
			"docker",
			"inspect",
			"-f",
			"{{range.NetworkSettings.Networks}}{{.IPAddress}}{{end}}",
			"metasploitable",
		)*/

		/*`ipout, err = ip_metasploitable.Output()
		if err != nil {
			fmt.Printf("Error getting IP address: %s\n", err)
			return
		}
		fmt.Println("Metasploitable IP Address:" + string(ipout))*/
	},
}

func init() {
	rootCmd.AddCommand(upCmd)

	// Here you will define your flags and configuration settings.
	upCmd.Flags().StringVar(
		&networkName,
		"network",
		"cyberlab",
		"Docker network name",
	)

	upCmd.Flags().StringVar(
		&driverName,
		"driver",
		"bridge",
		"Name of the Docker network driver to check",
	)

	upCmd.Flags().BoolVarP(
		&shell,
		"shell",
		"s",
		false,
		"Open an interactive shell in the pentest toolkit after starting the lab",
	)
	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// upCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// upCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
