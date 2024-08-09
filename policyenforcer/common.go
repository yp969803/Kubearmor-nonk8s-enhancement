package policyenforcer

import (
	"context"
	"fmt"

	"os"
	"path/filepath"

	"log"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
)

// Funtion to list docker containers
func ListDockerContainers(ctx context.Context) ([]types.Container, error) {
	client, err := client.NewClientWithOpts(client.WithAPIVersionNegotiation())
	if err != nil {
		client.Close()
		log.Println("Error while initializing DockerClient: %w", err)
		panic(err)
	}
	containers, err := client.ContainerList(ctx, container.ListOptions{})
	if err != nil {
		return nil, err
	}
	return containers, nil
}

func FindGoModRoot(startDir string) (string, error) {
	currentDir := startDir
	for {
		if _, err := os.Stat(filepath.Join(currentDir, "go.mod")); os.IsNotExist(err) {
			parentDir := filepath.Dir(currentDir)
			if parentDir == currentDir {
				// Reached the root of the file system
				break
			}
			currentDir = parentDir
		} else {
			return currentDir, nil
		}
	}
	return "", fmt.Errorf("go.mod file not found")
}
