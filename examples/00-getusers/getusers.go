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
	// Call the GetOperators function
	users, err := client.GetOperators()
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
	// Print all the users found
	fmt.Println(users)

	// Get the Operator based on the ID
	// Alternatives are:
	// 		client.GetOperatorByEmail(email)
	// 		client.GetOperatorByName(name)
	user, err := client.GetOperator(os.Getenv("BCUSERID"))
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	// Print the user
	fmt.Println(user)
}
