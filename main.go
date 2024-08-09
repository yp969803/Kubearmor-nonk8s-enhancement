package main

import (
	"fmt"

	"github.com/kubearmor_assigment/policyenforcer"
)

func main() {
	// err := policyenforcer.PolicyGenerator()
	err := policyenforcer.PolicyEnforcer()
	if err != nil {
		fmt.Println("Error : %w", err)
		panic(err)
	}
}