package main

func main() {}

type request = [][]string

func FindStartEndAirports(req request) ([]string, error) {
	if len(req) == 1 {
		return req[0], nil
	}

	var airportsCount = make(map[string]struct{
		count int
		left bool
	})
	for _, airports := range req {
		airportsCount[airports[0]] = struct{count int; left bool}{
			count: airportsCount[airports[0]].count+1,
			left: true,
		}
		airportsCount[airports[1]] = struct{count int; left bool}{
			count: airportsCount[airports[1]].count+1,
			left: false,
		}
	}

	startEnd := make([]string, 2)
	for airport, countLeft := range airportsCount {
		if countLeft.count == 1 && countLeft.left {
			startEnd[0] = airport
		}
		if countLeft.count == 1 && !countLeft.left {
			startEnd[1] = airport
		}
	}

	return startEnd, nil
}

// go through the list, if airport is there only 1 and if it's [0] then it's start
// go through the list, if airport is there only 1 and if it's [1] then it's end