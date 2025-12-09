package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

type JunctionBox struct {
	x, y, z float64
}

func (jb1 JunctionBox) String() string {
	return fmt.Sprintf("(%.0f,%.0f,%.0f)", jb1.x, jb1.y, jb1.z)
}

type Pair struct {
	x, y     JunctionBox
	distance float64
}

func (p Pair) String() string {
	return fmt.Sprintf("%s <-> %s : %.2f", p.x, p.y, p.distance)
}

func (jb1 JunctionBox) Distance(jb2 JunctionBox) float64 {
	return math.Sqrt(math.Pow(jb2.x-jb1.x, 2) + math.Pow(jb2.y-jb1.y, 2) + math.Pow(jb2.z-jb1.z, 2))
}

func getJunctionBoxes() []JunctionBox {
	file, err := os.Open("./input/day8.txt")
	if err != nil {
		panic("You are a complete and utter failure")
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	junctionBoxes := []JunctionBox{}
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), ",")
		x, _ := strconv.ParseFloat(line[0], 64)
		y, _ := strconv.ParseFloat(line[1], 64)
		z, _ := strconv.ParseFloat(line[2], 64)
		jb := JunctionBox{x: x, y: y, z: z}
		junctionBoxes = append(junctionBoxes, jb)
	}

	return junctionBoxes
}

func Part1Part2(jbs []JunctionBox) {
	maxConns := 1000
	if len(jbs) < 200 {
		maxConns = 10
	}
	distances := map[JunctionBox]map[JunctionBox]float64{}
	// initialize distances matrix
	for i, x := range jbs {
		distances[x] = make(map[JunctionBox]float64)
		for j, y := range jbs {
			if i != j && j > i {
				distances[x][y] = x.Distance(y)
			}
		}
	}

	// create a pairing of all distances
	pairs := []Pair{}
	for x, m := range distances {
		for y, d := range m {
			pairs = append(pairs, Pair{x: x, y: y, distance: d})
		}
	}

	// sort those guys
	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].distance < pairs[j].distance
	})

	circuits := map[JunctionBox]int{}
	circuitLists := [][]JunctionBox{}
	i := 0
	connections := 0
	totalCircuits := 0
	var lastP Pair
	for _, p := range pairs {

		_, xInCircuit := circuits[p.x]
		_, yInCircuit := circuits[p.y]

		if !xInCircuit && !yInCircuit {
			// If neither are in a circuit, assign them to their new circuit
			connections++
			circuits[p.x] = i
			circuits[p.y] = i
			circuitLists = append(circuitLists, []JunctionBox{p.x, p.y})
			i++
			totalCircuits++
		} else if xInCircuit && !yInCircuit {
			// if x is in circuit and y isn't, then put y in that circuit
			connections++
			circuits[p.y] = circuits[p.x]
			circuitLists[circuits[p.x]] = append(circuitLists[circuits[p.x]], p.y)
		} else if yInCircuit && !xInCircuit {
			// if y is in circuit and x isn't, then put x in that circuit
			connections++
			circuits[p.x] = circuits[p.y]
			circuitLists[circuits[p.y]] = append(circuitLists[circuits[p.y]], p.x)
		} else if xInCircuit && yInCircuit && circuits[p.x] != circuits[p.y] {
			// both are in different circuits, connect the circuits!
			connections++
			// Merge the circuits
			xID := circuits[p.x]
			yID := circuits[p.y]

			// move all yID circuit members to xID circuits
			for _, jb := range circuitLists[yID] {
				circuits[jb] = xID
				circuitLists[xID] = append(circuitLists[xID], jb)
			}
			// remove yID circuit
			circuitLists[yID] = []JunctionBox{}
			totalCircuits--
		} else {
			// both are in the same circuit already, do nothing
			connections++ // I didn't understand counting this..... but it is the only way it works
			continue
		}
		if totalCircuits == 1 {
			lastP = p
		}

		// print the number of circuits with thier lengths
		// fmt.Printf("Circuits: %d\n", totalCircuits)
		// // print the circuites
		// for id, list := range circuitLists {
		// 	fmt.Printf("Circuit %d[%d]: %v\n", id, len(list), list)
		// }
		// fmt.Printf("Connections made: %d\n", connections)
		// fmt.Println("-----")

		// Part1
		if connections == maxConns {
			copyList := make([][]JunctionBox, len(circuitLists))
			copy(copyList, circuitLists)

			sort.Slice(copyList, func(i, j int) bool {
				return len(copyList[i]) > len(copyList[j])
			})
			part1 := len(copyList[0]) * len(copyList[1]) * len(copyList[2])
			fmt.Println(part1)
		}
	}

	// Part2
	part2 := int(lastP.x.x) * int(lastP.y.x)
	fmt.Println(part2)
}

func main() {
	jbs := getJunctionBoxes()
	Part1Part2(jbs)
}
