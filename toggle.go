package flags

import (
	"fmt"
	"math"
	"math/rand"
)

// Toggle is a feature flag that is enabled for a configured fraction of
// evaluations.
type Toggle struct {
	// percentage is the probability, between 0.0 and 1.0, that Evaluate
	// returns true. 0.0 never enables the feature and 1.0 always does.
	percentage float64
}

// NewToggle returns a Toggle that is enabled for the given percentage of
// evaluations. percentage must be between 0.0 and 100.0 inclusive; NaN,
// infinities, and out-of-range values yield a nil Toggle and a non-nil error.
func NewToggle(percentage float64) (*Toggle, error) {
	if math.IsNaN(percentage) || percentage < 0 || percentage > 100 {
		return nil, fmt.Errorf("toggle percentage must be between 0.0 and 100.0, got %v", percentage)
	}

	return &Toggle{percentage: percentage}, nil
}

// Evaluate reports whether the feature is enabled for this call. Each call is
// an independent random draw, so the result is not sticky per user or request.
func (rec *Toggle) Evaluate() bool {
	return rec.percentage/100 > rand.Float64()
}
