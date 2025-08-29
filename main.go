package main

import (
	"fmt"
	"os"
	"time"
)

func main() {

	fmt.Printf("\n[%v][INFO] Starting k8s-annotator...", time.Now().Format("2006-01-02 15:04:05"))
	fmt.Printf("\n[%v][INFO] Temp solution for annotating taurus-cloud pods, reach out to manan for issues.", time.Now().Format("2006-01-02 15:04:05"))

	var agentmd AgentMetaData
	agentmd.ShipID = os.Getenv("SHIP_ID")
	agentmd.HarbourID = os.Getenv("HARBOUR_ID")
	agentmd.Ns = os.Getenv("NAMESPACE")

	fmt.Printf("\n[%v][INFO] Provided:\n ShipID: %s,\n HarbourID: %s,\n Namespace: %s\n", time.Now().Format("2006-01-02 15:04:05"), agentmd.ShipID, agentmd.HarbourID, agentmd.Ns)
	// Parse interval string to integer (minutes)
	cs := &ClientSet{}
	cs.getClientSet()

	ticker := time.NewTicker(5 * time.Minute)
	defer ticker.Stop()
	for range ticker.C {
		podUpdateAnnotaion(&agentmd, cs)
	}

}
