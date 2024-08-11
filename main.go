package main

import (
	"flag"
	"fmt"

	"github.com/kubearmor_assigment/policyenforcer"
)

var (
	mode = flag.String("mode", "generate", "Mode of the policy enforcer")
)

func main() {

	flag.Parse()

	if *mode == "enforce" {
		err := policyenforcer.PolicyEnforcer()
		if err != nil {
			fmt.Println("Error : %w", err)
			panic(err)
		}
	} else if *mode == "generate" {
		err := policyenforcer.PolicyGenerator()
		if err != nil {
			fmt.Println("Error : %w", err)
			panic(err)
		}
	} else if *mode == "generateWithoutBlank" {
		err := policyenforcer.PolicyGeneratorWithoutBlank()
		if err != nil {
			fmt.Println("Error : %w", err)
			panic(err)
		}
	}

}
