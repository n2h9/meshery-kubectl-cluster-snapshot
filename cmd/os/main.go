package main

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/meshery/kubectl-cluster-snapshot/internal/utils"
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
				outputFile, err = utils.GenerateUniqueFileNameForSnapshot()
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
	rootCmd.Flags().StringVarP(&namespace, "namespace", "n", "", "namespace to make a snapshot of")
	rootCmd.Flags().StringVarP(&outputFile, "outputFile", "o", "", "name of a result output file")

	// Execute the command
	if err := rootCmd.Execute(); err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
}
