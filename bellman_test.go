package bellman

import (
	"fmt"
	"testing"
)

func TestSearch(t *testing.T) {

	rt := Edges{
		{"A", "E", -3},
		{"B", "C", 7},
		{"C", "F", -5},
		{"E", "B", 1},
		{"E", "G", 6},
		{"E", "H", -3},
		{"F", "B", 8},
		{"G", "D", 2},
		{"H", "G", 1},
		{"S", "A", 2},
	}

	tb, err := rt.Search("S")
	if err != nil {
		t.Error(err)
	}

	distance := map[string]int{
		"S": 0, "A": 2, "B": 0,
		"C": 7, "D": -1, "E": -1,
		"F": 2, "G": -3, "H": -4,
	}
	// list of predecessors
	pred := map[string]string{
		"S": " ", "A": "S", "B": "E",
		"C": "B", "D": "G", "E": "A",
		"F": "C", "G": "H", "H": "E",
	}

	for b, n := range tb {
		if distance[b] != n.Distance {
			t.Errorf("route %q distance got %d want %d", b, n.Distance, distance[b])
		}
		if pred[b] != n.From {
			t.Errorf("route %q predecessor got %q want %q", b, n.From, pred[b])
		}
	}
}

type testcase struct {
	src, dest string
	path      string
}

var tests = []*testcase{
	&testcase{"San Luis Obispo", "Los Angeles", "San Luis Obispo 106,Santa Barbara 95,"},
	&testcase{"Santa Barbara", "Las Vegas", "Santa Barbara 95,Los Angeles 65,San Bernardino 73,Barstow 62,Baker 92,"},
	&testcase{"San Diego", "Los Angeles", "San Diego 121,"},
}

func TestShortestPath(t *testing.T) {

	for _, test := range tests {

		rt := Edges{
			{"San Luis Obispo", "Bakersfield", 117},
			{"Bakersfield", "Mojave", 65},
			{"Mojave", "Barstow", 70},
			{"Barstow", "Baker", 62},
			{"Baker", "Las Vegas", 92},
			{"San Luis Obispo", "Santa Barbara", 106},
			{"San Luis Obispo", "Santa Barbara", 113},
			{"Santa Barbara", "Los Angeles", 95},
			{"Bakersfield", "Wheeler Ridge", 24},
			{"Wheeler Ridge", "Los Angeles", 88},
			{"Mojave", "Los Angeles", 94},
			{"Los Angeles", "San Bernardino", 65},
			{"San Bernardino", "Barstow", 73},
			{"Los Angeles", "San Diego", 121},
			{"San Bernardino", "San Diego", 103},
		}
		for _, r := range rt {
			// Add reversed edges.
			rt = append(rt, &Edge{From: r.To, To: r.From, Distance: r.Distance})
		}

		paths, err := rt.Search(test.src)
		if err != nil {
			t.Fatal(err)
		}

		all, err := paths.ShortestPath(test.src, test.dest)
		if err != nil {
			t.Fatal(err)
		}

		var prev int
		var s string
		for i := len(all) - 1; i >= 0; i-- {
			s += fmt.Sprintf("%s %d,", all[i].From, all[i].Distance-prev)
			prev = all[i].Distance
		}

		if test.path != s {
			t.Errorf("shortest path does not match, got %q want %q", s, test.path)
		}
	}
}
