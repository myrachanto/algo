package main

import (
	"fmt"
	"strings"
	"time"
)

type resource struct {
	name   string
	value  int
	weight float64
	volume float64
}

func decodeGenes(candidate string, resources []resource, geneSet string) map[resource]int {
	resourceCounts := make(map[resource]int, len(candidate)/2)
	for i := 0; i < len(candidate); i++ {
		chromosome := candidate[i : i+2]
		resourceId := scale(strings.Index(geneSet, chromosome[0:1]), len(geneSet), len(resources))
		resourceCount := strings.Index(geneSet, chromosome[1:2])
		resource := resources[resourceId]
		resourceCounts[resource] = resourceCounts[resource] + resourceCount
	}
	return resourceCounts
}
func scale(value, currentMax, newMax int) int {
	return value * newMax / currentMax
}

// The fitness function will use the decoded resource counts and the resource constraints to determine an appropriate fitness for the candidate.
func getFitness(resourceCounts map[resource]int, maxWeight float64, maxVolume float64) int {
	weight, volume, value := 0.0, 0.0, 0
	for resource, count := range resourceCounts {
		weight += resource.volume * float64(count)
		volume += resource.volume * float64(count)
		value += resource.value * count
	}
	if weight > maxWeight || volume > maxVolume {
		return -value
	}
	return int(value)
}

// Next is a display function that tells us which resources and how many of each to take and their value (fitness).
func display(resourceCounts map[resource]int, fitnes int, elapsed time.Duration) {
	label := ""
	for resource, count := range resourceCounts {
		label += fmt.Sprint(count, " ", resource.name, "")
	}
	fmt.Println(fitnes, "\t", label, "\t", elapsed)
}
func krilo() {
	// Name 	Value (each) 	Weight 	Volume (each)
	// Bark 	3000 			0.3 	.025
	// Herb 	1800 			0.2 	.015
	// Root 	2500 			2.0 	.002
	// max weight = 25
	// max vol 0.25
	// Our goal is to maximize the value of the contents of the knapsack without exceeding either of the weight or volume limits.
	// Let’s think about how we would solve this problem by hand. We want to maximize the value within the constraints. So we want a high ratio of value to weight and value to volume. And we want as many of those in the bag as we can get. When we can’t stuff any more of the top item into the bag, we fill in the remaining space with the next most valuable item, and so on. This process is known as hill climbing.
	// resources := []resource{
	// 	{name: "Bark", value: 3000, weight: 0.3, volume: .025},
	// 	{name: "Herb", value: 1800, weight: 0.2, volume: .015},
	// 	{name: "Root", value: 2500, weight: 2.0, volume: .002},
	// }

	// const maxWeight = 25.0
	// const maxVolume = .25

	// geneSet := "0123456789ABCDEFGH"

	// calc := func(candidate string) int {
	//     decoded := decodeGenes(candidate, resources, geneSet)
	//     return getFitness(decoded, maxWeight, maxVolume)
	// }

	// start := time.Now()

	// disp := func(candidate string) {
	//     decoded := decodeGenes(candidate, resources, geneSet)
	//     fitness := getFitness(decoded, maxWeight, maxVolume)
	//     display(decoded, fitness, time.Since(start))
	// }

	// var solver = new(genetic.Solver)
	// solver.MaxSecondsToRunWithoutImprovement = .1
	// solver.MaxRoundsWithoutImprovement = 2

	// var best = solver.GetBestUsingHillClimbing(calc, disp, geneSet, 10, 2, math.MaxInt32)

	fmt.Println("\nFinal:")
	// disp(best)
}
