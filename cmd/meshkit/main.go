package main

import (
	"context"
	"fmt"
	"os"

	"github.com/meshery/kubectl-cluster-snapshot/internal/utils"
	"github.com/spf13/cobra"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/yaml"

	mesherykube "github.com/layer5io/meshkit/utils/kubernetes"
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
			// Initialize kubeclient
			kubeClient, err := mesherykube.New(nil)
			if err != nil {
				return err
			}

			list, err := kubeClient.KubeClient.CoreV1().Pods(namespace).List(
				context.TODO(),
				metav1.ListOptions{},
			)
			if err != nil {
				return err
			}
			list.APIVersion = "v1"
			list.Kind = "List"
			for i := range list.Items {
				list.Items[i].APIVersion = "v1"
				list.Items[i].Kind = "Pod"
			}

			if outputFile == "" {
				outputFile, err = utils.GenerateUniqueFileNameForSnapshot()
				if err != nil {
					return err
				}
			}

			output, err := yaml.Marshal(list)
			if err != nil {
				return err
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
