package interval_test

import (
	"testing"

	"github.com/panagiotisptr/cov-diff/files"
	"github.com/panagiotisptr/cov-diff/interval"
)

func TestUnionIgnoresMalformedBoundsFromEmptyFunction(t *testing.T) {
	functionIntervals, err := files.GetIntervalsFromFile([]string{
		"package fixture",
		"",
		"func empty() {",
		"}",
	}, false)
	if err != nil {
		t.Fatalf("get function intervals: %v", err)
	}
	got := interval.Union(
		[]interval.Interval{{Start: 1, End: 4}},
		functionIntervals,
	)
	if len(got) != 0 {
		t.Fatalf("expected no valid intersections, got %#v", got)
	}
	if whitespace := interval.TotalWhitespace(got, []string{"", "", "", ""}); whitespace != 0 {
		t.Fatalf("expected no whitespace, got %d", whitespace)
	}
}
