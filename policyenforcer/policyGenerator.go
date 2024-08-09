package policyenforcer

import (
	"context"
	"encoding/json"
	"fmt"

	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/kubearmor_assigment/types"
	"sigs.k8s.io/yaml"
)


//  Generate policy in outpolicy folder
func PolicyGenerator() error {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	containers, err := ListDockerContainers(ctx)
	if err != nil {
		return err
	}

	startDir, err := os.Getwd()
	if err != nil {
		fmt.Println("Error:", err)
		return err
	}
	rootDir, err := FindGoModRoot(startDir)
	if err != nil {
		fmt.Println("Error:", err)
	}

	var containerPolicy types.KubearmorContainerPolicy

	path := filepath.Join(rootDir, "example-policy", "kubearmor_containerpolicy.yaml")
	policy, err := os.ReadFile(filepath.Clean(path))

	if err != nil {
		return err
	}

	js, err := yaml.YAMLToJSON([]byte(policy))
	if err != nil {
		return err
	}
	err = json.Unmarshal(js, &containerPolicy)

	if err != nil {
		return err
	}

	for _, container := range containers {
		newContainerPolicy := containerPolicy
		containerName := strings.TrimPrefix(container.Names[0], "/")

		newContainerPolicy.Spec.Selector.MatchLabels["kubearmor.io/container.name"] = containerName
		jsonData, err := json.Marshal(newContainerPolicy)
		if err != nil {
			return err
		}
		yaml, err := yaml.JSONToYAML(jsonData)
		if err != nil {
			return err
		}

		fileName := filepath.Join(rootDir, "outpolicy", containerName+"-policy.yaml")

		err = os.WriteFile(fileName, yaml, 0644)
		if err != nil {
			return err
		}

	}

	return nil
}
