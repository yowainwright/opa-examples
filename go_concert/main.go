package main

import (
	"context"
	"fmt"
	"os"

	"github.com/open-policy-agent/opa/rego"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <wristband_color>")
		return
	}

	wristbandColor := os.Args[1]

	ctx := context.Background()

	// Load the policy
	r := rego.New(
		rego.Query("data.concert.privileges"),
		rego.Load([]string{"./policy.rego"}, nil),
	)

	// Prepare the evaluation
	regoPrepared, err := r.PrepareForEval(ctx)
	if err != nil {
		fmt.Println("Error preparing for eval:", err)
		return
	}

	// Define the input with the wristband color
	input := map[string]interface{}{
		"wristband_color": wristbandColor,
	}

	// Evaluate the policy with the input
	results, err := regoPrepared.Eval(ctx, rego.EvalInput(input))
	if err != nil {
		fmt.Println("Error evaluating policy:", err)
		return
	}

	if len(results) == 0 {
		fmt.Println("No results returned.")
		return
	}

	// Extract the privileges
	privileges := results[0].Expressions[0].Value.(map[string]interface{})

	// Print the privileges
	fmt.Printf("Wristband color: %s\n", wristbandColor)
	fmt.Printf("Allow entry: %t\n", privileges["allow_entry"].(bool))
	fmt.Printf("Allow VIP access: %t\n", privileges["allow_vip_access"].(bool))
	fmt.Printf("Allow free food: %t\n", privileges["allow_free_food"].(bool))
}
