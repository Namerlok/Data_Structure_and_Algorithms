func canCompleteCircuit(gas []int, cost []int) int {
	startStation := 0
	totalBalance := 0
	currentTank := 0
	for i := 0; i < len(gas); i++ {
		totalBalance += gas[i] - cost[i]
		currentTank += gas[i] - cost[i]
		if currentTank < 0 {
			startStation = (i + 1) % len(gas)
			currentTank = 0
		}
	}

	if totalBalance < 0 {
		return -1
	} else {
		return startStation
	}
}