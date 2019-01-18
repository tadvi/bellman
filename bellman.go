// Bellman-Ford algorithm for a DAG. If you want to use it for undirected
// graph then simply add pair of routes into both directions. Make sure to
// add two routes of the same value otherwise it will form a cycle.
package bellman

import (
	"errors"
)

const inf = int(^uint(0) >> 1)

var ErrNegativeCycle = errors.New("bellman: graph contains negative cycle")
var ErrNoPath = errors.New("bellman: no path to vertex")

type Edges []*Edge

// Pair of vertices.
type Edge struct {
	From, To string
	Distance int
}

type Vertices map[string]*Vertex

// Vertex with predecessor information.
type Vertex struct {
	Distance int
	From     string // predecessor
}

// ShortestPath finds shortest path from source to destination after calling Search.
func (paths Vertices) ShortestPath(src, dest string) ([]*Vertex, error) {
	if len(paths) == 0 {
		// Did you forget calling Search?
		return nil, ErrNoPath
	}

	var all []*Vertex
	pred := dest

	for pred != src {
		if v := paths[pred]; v != nil {
			pred = v.From
			all = append(all, v)
		} else {
			return nil, ErrNoPath
		}
	}
	return all, nil
}

// AddEdge for search.
func (rt *Edges) AddEdge(from, to string, distance int) {
	*rt = append(*rt, &Edge{From: from, To: to, Distance: distance})
}

// Search for single source shortest path.
func (rt Edges) Search(start string) (Vertices, error) {
	// Resulting table constains vertex name to Vertex struct mapping.
	// Use v.From predecessor to trace the path back.
	tb := make(Vertices)
	for _, p := range rt {
		tb[p.From] = &Vertex{inf, " "}
		tb[p.To] = &Vertex{inf, " "}
	}
	tb[start].Distance = 0

	// As many iterations as there are nodes.
	for i := 0; i < len(tb); i++ {
		var changed bool

		// Iterate over pairs.
		for _, pair := range rt {
			n := tb[pair.From]
			if n.Distance != inf {
				if tb[pair.To].Distance > n.Distance+pair.Distance {
					tb[pair.To].Distance = n.Distance + pair.Distance
					tb[pair.To].From = pair.From
					changed = true
				}
			}
		}

		if !changed {
			// No more changes made. We done.
			break
		}
	}

	for _, pair := range rt {
		if tb[pair.From].Distance != inf &&
			tb[pair.To].Distance != inf &&
			tb[pair.From].Distance+pair.Distance < tb[pair.To].Distance {
			return nil, ErrNegativeCycle
		}
	}
	return tb, nil
}
