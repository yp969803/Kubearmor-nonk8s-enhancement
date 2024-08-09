package policyenforcer

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	tp "github.com/kubearmor/KubeArmor/KubeArmor/types"
	pb "github.com/kubearmor/KubeArmor/protobuf"

	"google.golang.org/grpc"

	"sigs.k8s.io/yaml"
)

type PolicyOptions struct {
	GRPC string
}

func sendPolicyOverGRPC(o PolicyOptions, policyEventData []byte) error {
	gRPC := ""

	if o.GRPC != "" {
		gRPC = o.GRPC
	} else {
		if val, ok := os.LookupEnv("KUBEARMOR_SERVICE"); ok {
			gRPC = val
		} else {
			gRPC = "localhost:32767"
		}
	}

	conn, err := grpc.Dial(gRPC, grpc.WithInsecure())
	if err != nil {
		return err
	}

	client := pb.NewPolicyServiceClient(conn)

	req := pb.Policy{
		Policy: policyEventData,
	}

	resp, err := client.ContainerPolicy(context.Background(), &req)
	if err != nil {
		return fmt.Errorf("failed to send policy")
	}
	fmt.Printf("Policy %s \n", resp.Status)
	return nil

}

func PolicyEnforcer() error {
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

	var containerPolicy tp.K8sKubeArmorPolicy

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

		policyEvent := tp.K8sKubeArmorPolicyEvent{
			Type:   "ADDED",
			Object: containerPolicy,
		}
		policyEventData, err := json.Marshal(policyEvent)
		if err != nil {
			return err
		}
		o := PolicyOptions{}

		if err = sendPolicyOverGRPC(o, policyEventData); err != nil {
			return err
		}
	}

	return nil
}
