package main

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"time"

	"github.com/spf13/cobra"
)

func main() {
	var namespace string
	var outputFile string

	// Root command for the plugin
	var rootCmd = &cobra.Command{
		Use:   "kubectl-meshery-cluster-snapshot",
		Short: "kubectl krew plugin to capture a cluster snapshot",
		Long:  "",
		RunE: func(cmd *cobra.Command, args []string) error {
			execCMD := exec.Command("kubectl", "--namespace", namespace, "get", "all", "-o", "yaml") // Get all pods in all namespaces

			output, err := execCMD.CombinedOutput() // Capture both stdout and stderr
			if err != nil {
				return err
			}

			if outputFile == "" {
				outputFile, err = generateUniqueFileNameForSnapshot()
				if err != nil {
					return err
				}
			}

			err = os.WriteFile(outputFile, output, 0644)
			if err != nil {
				return err
			}

			return nil
		},
	}

	// Add flags
	rootCmd.Flags().StringVarP(&namespace, "namespace", "n", "default", "namespace to make a snapshot of")
	rootCmd.Flags().StringVarP(&outputFile, "outputFile", "o", "", "name of a result output file")

	// Execute the command
	if err := rootCmd.Execute(); err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
}

func generateUniqueFileNameForSnapshot() (string, error) {
	// Get the current date and time
	currentTime := time.Now()

	// Format the current date in YYYYMMDD format
	currentDate := currentTime.Format("20060102")

	slug := fmt.Sprintf("meshery-cluster-snapshot-%s", currentDate)

	name := ""
	gotTheName := false
	for i := 0; i < 1024; i++ {
		name = fmt.Sprintf("%s-%02d.yaml", slug, i)
		// Use os.Stat to check if the file exists
		_, err := os.Stat(name)
		if os.IsNotExist(err) {
			gotTheName = true
			break
		}
	}

	if !gotTheName {
		return "", errors.New("no unique name available")
	}
	return name, nil
}
