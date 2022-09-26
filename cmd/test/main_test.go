package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFindStartEndAirports(t *testing.T) {
	t.Run("only-2", func(t *testing.T) {
		req := [][]string{{"SFO", "EWR"}}
		resp, err := FindStartEndAirports(req)
		require.NoError(t, err)
		require.Equal(t, resp, []string{"SFO", "EWR"})
	})

	t.Run("only-4", func(t *testing.T) {
		req := [][]string{{"ATL", "EWR"}, {"SFO", "ATL"}}
		resp, err := FindStartEndAirports(req)
		require.NoError(t, err)
		require.Equal(t, resp, []string{"SFO", "EWR"})
	})

	t.Run("only-8", func(t *testing.T) {
		req := [][]string{{"IND", "EWR"}, {"SFO", "ATL"}, {"GSO", "IND"}, {"ATL", "GSO"}}
		resp, err := FindStartEndAirports(req)
		require.NoError(t, err)
		require.Equal(t, resp, []string{"SFO", "EWR"})
	})
}
