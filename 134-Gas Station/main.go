package main

func canCompleteCircuit(gas []int, cost []int) int {
	var ptr, i, tank int
	var startStation, visitedStations int = -1, 1
	var totalCost, totalGas int

	for i := 0; i < len(gas); i++ {
		totalGas += gas[i]
		totalCost += cost[i]
	}
	if totalCost > totalGas {
		return -1
	}

	for {

		if ptr%len(gas) == startStation {
			if visitedStations > len(gas) {
				return startStation
			} else {
				return -1
			}
		}

		i = ptr % len(gas)

		if visitedStations <= 1 {
			startStation = i
		}

		if tank+gas[i] < cost[i] {
			if startStation+1 >= len(gas) {
				return -1
			}
			ptr = i + 1
			visitedStations = 1
			tank = 0
		} else {
			tank += gas[i]
			tank -= cost[i]
			visitedStations++
			ptr++
		}
	}
}
