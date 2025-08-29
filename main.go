package main

import (
	"os"
	"time"
)

func main() {

	var agentmd AgentMetaData
	agentmd.ShipID = os.Getenv("SHIP_ID")
	agentmd.HarbourID = os.Getenv("HARBOUR_ID")
	agentmd.Ns = os.Getenv("NAMESPACE")
	// I need this operations to be performed every 5 minutes
	ticker := time.NewTicker(3 * time.Minute)
	defer ticker.Stop()
	for range ticker.C {
		podUpdateAnnotaion(&agentmd)
	}

}
