package main

import "log"

func main() {
	log.Println(edges_to_adjacensy_list([][]string{
		{"i", "j"},
		{"k", "i"},
		{"k", "j "},
		{"m", "k"},
		{"k", "l"},
		{"o", "n"},
	}))
}

func edges_to_adjacensy_list(edges [][]string) map[string][]string {
	al := map[string][]string{}
	if len(edges) == 0 {
		return al
	}

	for _, edge := range edges {
		a, b := edge[0], edge[1]
		if aval, exists := al[a]; exists {
			aval = append(aval, b)
			al[a] = aval
		} else {
			aval := []string{b}
			al[a] = aval
		}

		if bval, exists := al[b]; exists {
			bval = append(bval, a)
			al[b] = bval
		} else {
			bval := []string{a}
			al[b] = bval
		}
	}

	return al
}
