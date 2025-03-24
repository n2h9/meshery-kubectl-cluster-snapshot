package main

import (
	"fmt"
	"os"
	"os/exec"
	"time"

	"github.com/spf13/cobra"
)

func main() {
	var namespace string
	var outputFile string
	var meshsyncBinaryPath string

	// Root command for the plugin
	var rootCmd = &cobra.Command{
		Use:   "kubectl-meshery-cluster-snapshot",
		Short: "kubectl krew plugin to capture a cluster snapshot",
		Long:  "",
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := fileExists(meshsyncBinaryPath); err != nil {
				return err
			}

			applyResources()

			fmt.Println("startings meshsync....")

			go func() {

				execCMD := exec.Command(
					meshsyncBinaryPath,
					"--output",
					"file",
					"--outputFile",
					outputFile,
				)
				_, err := execCMD.CombinedOutput()
				if err != nil {
					fmt.Println("Error:", err)
					os.Exit(1)
				}
			}()

			// wait for 8 seconds for now :)
			time.Sleep(time.Second * 8)

			// and now we could stop the execution

			deleteResources()

			fmt.Println("done")

			return nil
		},
	}

	// Add flags
	rootCmd.Flags().StringVarP(&namespace, "namespace", "n", "", "namespace to make a snapshot of")
	rootCmd.Flags().StringVarP(&outputFile, "outputFile", "o", "", "name of a result output file")
	rootCmd.Flags().StringVarP(&meshsyncBinaryPath, "meshsyncBinaryPath", "", "", "path to meshsync binary")

	// Execute the command
	if err := rootCmd.Execute(); err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
}
