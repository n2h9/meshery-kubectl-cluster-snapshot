package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/meshery/kubectl-cluster-snapshot/yamls"
)

var mesheryNamespace = "meshery"

func applyResources() {
	root := []string{"crds.yaml"}
	namespaced := []string{"broker.yaml", "meshsync.yaml"}

	for _, file := range root {
		content, err := yamls.Files.ReadFile(file)
		if err != nil {
			log.Fatal(err)
		}
		applyYAMLString(content, false)
	}

	for _, file := range namespaced {
		content, err := yamls.Files.ReadFile(file)
		if err != nil {
			log.Fatal(err)
		}
		applyYAMLString(content, true)
	}
}

func deleteResources() {
	// TODO
}

func applyYAMLString(yamlContent []byte, namespaced bool) error {
	// Write the YAML string to a temporary file
	tmpFile, err := os.CreateTemp("", "kubectl-meshery-cluster-snapshot-plugin-*.yaml")
	if err != nil {
		return err
	}
	defer os.Remove(tmpFile.Name()) // Clean up the temporary file after use

	// Write YAML content to the temporary file
	_, err = tmpFile.Write(yamlContent)
	if err != nil {
		return err
	}
	tmpFile.Close() // Close the file before using it in kubectl

	// Apply the YAML using kubectl
	parts := []string{"apply", "-f", tmpFile.Name()}
	if namespaced {
		parts = append([]string{"--namespace", mesheryNamespace}, parts...)
	}

	cmd := exec.Command("kubectl", parts...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("error applying YAML: %v", err)
	}

	fmt.Println("YAML applied successfully!")
	return nil
}
