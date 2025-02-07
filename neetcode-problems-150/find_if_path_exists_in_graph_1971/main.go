package main

import "log"

func main() {
	edges := [][]int{
		{4, 3}, {1, 4}, {4, 8}, {1, 7}, {6, 4}, {4, 2}, {7, 4}, {4, 0}, {0, 9}, {5, 4},
	}

	log.Println(edges_to_adjacensy_list(edges))
	validPath(10, edges, 5, 9)
}

func validPath(n int, edges [][]int, source int, destination int) bool {
	seen := map[int]struct{}{}
	adjList := edges_to_adjacensy_list(edges)
	return exploreGraph(adjList, source, destination, seen)

}

func exploreGraph(adjList map[int][]int, src, dst int, seen map[int]struct{}) bool {
	if src == dst {
		return true
	}

	if _, exists := seen[src]; !exists {
		seen[src] = struct{}{}
	}

	for _, neighbour := range adjList[src] {
		if _, exists := seen[neighbour]; exists {
			continue
		}

		if exploreGraph(adjList, neighbour, dst, seen) {
			return true
		}
	}

	return false
}

func edges_to_adjacensy_list(edges [][]int) map[int][]int {
	al := map[int][]int{}
	if len(edges) == 0 {
		return al
	}

	for _, edge := range edges {
		a, b := edge[0], edge[1]
		if aval, exists := al[a]; exists {
			aval = append(aval, b)
			al[a] = aval
		} else {
			aval := []int{b}
			al[a] = aval
		}

		if bval, exists := al[b]; exists {
			bval = append(bval, a)
			al[b] = bval
		} else {
			bval := []int{a}
			al[b] = bval
		}
	}

	return al
}
