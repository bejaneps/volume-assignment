package calculate_test

import (
	"testing"

	"github.com/bejaneps/volume-assignment/internal/service/calculate"
	"github.com/stretchr/testify/require"
)

func TestFindStartEndAirports(t *testing.T) {
	svc := calculate.New()

	t.Run("only-2", func(t *testing.T) {
		airports := [][]string{{"SFO", "EWR"}}
		resp, err := svc.FindStartEndAirports(airports)
		require.NoError(t, err)
		require.Equal(t, resp, []string{"SFO", "EWR"})
	})

	t.Run("only-4", func(t *testing.T) {
		airports := [][]string{{"ATL", "EWR"}, {"SFO", "ATL"}}
		resp, err := svc.FindStartEndAirports(airports)
		require.NoError(t, err)
		require.Equal(t, resp, []string{"SFO", "EWR"})
	})

	t.Run("only-8", func(t *testing.T) {
		airports := [][]string{{"IND", "EWR"}, {"SFO", "ATL"}, {"GSO", "IND"}, {"ATL", "GSO"}}
		resp, err := svc.FindStartEndAirports(airports)
		require.NoError(t, err)
		require.Equal(t, resp, []string{"SFO", "EWR"})
	})
}
