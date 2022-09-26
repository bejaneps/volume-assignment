package calculate

// Service represents a business layer of Airport route calculator
type Service struct{}

// New constructs Airport route calculator service
func New() *Service {
	return &Service{}
}

type airportInfo struct {
	count  int
	isStart bool
}

// FindStartEndAirports accepts a list of sorted or unsorted airports
// and returns a start and end airports
func (s *Service) FindStartEndAirports(airports [][]string) ([]string, error) {
	if len(airports) == 1 {
		return airports[0], nil
	}

	// iterate over airports list and increase count of each airport seen
	// and also mark if airport is starting point or not
	airportsinfo := make(map[string]airportInfo)
	for _, startEnd := range airports {
		airportsinfo[startEnd[0]] = airportInfo{
			count:  airportsinfo[startEnd[0]].count + 1,
			isStart: true,
		}
		airportsinfo[startEnd[1]] = airportInfo{
			count: airportsinfo[startEnd[1]].count + 1,
		}
	}

	// iterate over list of airports with their counts and start end info
	// if airport is seen only once then check if it's either start or end
	startEnd := make([]string, 2)
	for airport, info := range airportsinfo {
		if info.count == 1 && info.isStart {
			startEnd[0] = airport
		}
		if info.count == 1 && !info.isStart {
			startEnd[1] = airport
		}
	}

	return startEnd, nil
}
