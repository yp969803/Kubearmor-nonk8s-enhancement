package types

import (
	tb "github.com/kubearmor/KubeArmor/KubeArmor/types"
)

type KubearmorContainerPolicy struct {
	ApiVersion string `json:"apiVersion"`
	Kind       string `json:"kind"`
	tb.K8sKubeArmorPolicy
}
