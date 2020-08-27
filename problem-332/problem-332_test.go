package problem_332

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindItinerary(t *testing.T) {
	assert.EqualValues(t, []string{"JFK", "MUC", "LHR", "SFO", "SJC"}, findItinerary([][]string{{"MUC", "LHR"}, {"JFK", "MUC"}, {"SFO", "SJC"}, {"LHR", "SFO"}}))
	assert.EqualValues(t, []string{"JFK", "ATL", "JFK", "SFO", "ATL", "SFO"}, findItinerary([][]string{{"JFK", "SFO"}, {"JFK", "ATL"}, {"SFO", "ATL"}, {"ATL", "JFK"}, {"ATL", "SFO"}}))
}
