package testUtil

/*
 * Module Dependencies
 */

import (
	"testing"
)

/*
 * Variables
 */

/*
 * Functions
 */

func Match(t *testing.T, expected interface{}, actual interface{}) {
	if expected != actual {
		t.Errorf("Expected: %v, Actual: %v", expected, actual)
	}
}

func NoError(t *testing.T, err error) {
	if err != nil {
		t.Errorf("Got error: %v", err)
	}
}

func WithError(t *testing.T, err error) {
	if err == nil {
		t.Errorf("Should got error but nil")
	}
}

/*
 * Tests
 */
