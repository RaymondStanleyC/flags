// Package flags provides lightweight feature flag primitives for gating code
// paths at runtime.
//
// It currently offers percentage-based rollouts: a Toggle is configured with a
// probability between 0.0 and 1.0 and, each time it is evaluated, randomly
// decides whether the feature is enabled. This makes it useful for gradual
// rollouts, canary releases, and simple A/B experiments.
package flags

import (
	"math/rand"
)

// Toggle is a feature flag that is enabled for a configured fraction of
// evaluations.
type Toggle struct {
	// percentage is the probability, between 0.0 and 1.0, that evaluate
	// returns true. 0.0 never enables the feature and 1.0 always does.
	percentage float64
}

// evaluate reports whether the feature is enabled for this call. Each call is
// an independent random draw, so the result is not sticky per user or request.
func (rec *Toggle) evaluate() bool {
	return rec.percentage > rand.Float64()
}
