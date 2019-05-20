package input

type HitStats struct {
	Parameters
	Hits int `json:"hits"`
}

var stats = make(map[Parameters]int)
var mostPopular HitStats

func RegisterStats(params Parameters) int {
	hits := stats[params] + 1
	stats[params] = hits
	if hits > mostPopular.Hits {
		mostPopular = HitStats{params, hits}
	}
	return hits
}

func MostPopular() HitStats {
	return mostPopular
}
