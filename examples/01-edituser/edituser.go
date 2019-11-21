package main

import (
	"fmt"
	"github.com/hlfstr/boldchat"
	"os"
)

/* In this example, a client is created and the Operators are retrieved
   via an API Call
*/
func main() {
	// Print the Version, not required
	fmt.Println(boldchat.Version())
	// Create the client
	// 		Requires the ID, SettingID, API Key, and BoldChat Endpoint
	client := boldchat.New(
		os.Getenv("BCID"),
		os.Getenv("BCSETTING"),
		os.Getenv("BCKEY"),
		boldchat.US,
	)

	// Create an Operator Object to edit
	// Requires the LoginID to be filled
	user := &boldchat.Operator{
		LoginID: os.Getenv("BCUSERID"),
	}

	fmt.Println("Pushing: ", user)

	// Use the EditOperator API Call.
	// Pass in the 'user' object created earlier
	err := client.EditOperator(user)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
}
