package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

func main() {

	fmt.Printf("\n[%v][INFO] Starting k8s-annotator...", time.Now().Format("2006-01-02 15:04:05"))
	fmt.Printf("\n[%v][INFO] Temp solution for annotating taurus-cloud pods, reach out to manan for issues.", time.Now().Format("2006-01-02 15:04:05"))

	var agentmd AgentMetaData
	agentmd.ShipID = os.Getenv("SHIP_ID")
	agentmd.HarbourID = os.Getenv("HARBOUR_ID")
	agentmd.Ns = os.Getenv("NAMESPACE")
	agentmd.Interval = os.Getenv("INTERVAL")

	fmt.Printf("\n[%v][INFO] Provided:\n ShipID: %s\n, HarbourID: %s\n, Namespace: %s", time.Now().Format("2006-01-02 15:04:05"), agentmd.ShipID, agentmd.HarbourID, agentmd.Ns)
	// Parse interval string to integer (minutes)
	intervalMinutes, err := strconv.Atoi(agentmd.Interval)
	if err != nil {
		fmt.Printf("\n[%v][ERROR] Invalid INTERVAL value: %v\n", time.Now().Format("2006-01-02 15:04:05"), err)
		os.Exit(1)
	}

	ticker := time.NewTicker(intervalMinutes * time.Minute)
	defer ticker.Stop()
	for range ticker.C {
		podUpdateAnnotaion(&agentmd)
	}

}
