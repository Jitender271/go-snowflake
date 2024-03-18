package main

import "fmt"

func main() {

	// Create a new Node with a Node number of 1
	node, err := NewSnowflakeNode(2)
	if err != nil {
		fmt.Println(err)
		return
	}

	// GenerateID a snowflake ID.
	id := node.GenerateID()

	// Print out the ID in a few different ways.
	fmt.Printf("Int64  ID: %d\n", id)
	fmt.Printf("String ID: %s\n", id)
	fmt.Printf("Base2  ID: %s\n", id.Base2())
	fmt.Printf("Base64 ID: %s\n", id.Base64())

	// Print out the ID's timestamp
	fmt.Printf("ID Time  : %d\n", id.Time())

	// Print out the ID's node number
	fmt.Printf("ID Node  : %d\n", id.Node())

	// Print out the ID's sequence number
	fmt.Printf("ID Step  : %d\n", id.Step())

	// GenerateID and print, all in one.
	fmt.Printf("ID       : %d\n", node.GenerateID().Int64())
}
