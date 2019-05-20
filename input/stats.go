package input

import "sync"

type HitStats struct {
	Parameters
	Hits int `json:"hits"`
}

var stats = make(map[Parameters]int)
var mostPopular HitStats
var mux sync.Mutex

func RegisterHit(params Parameters) int {
	mux.Lock()
	defer mux.Unlock()
	hits := stats[params] + 1
	stats[params] = hits
	if hits > mostPopular.Hits {
		mostPopular = HitStats{params, hits}
	}
	return hits
}

func MostPopular() HitStats {
	mux.Lock()
	defer mux.Unlock()
	return mostPopular
}
