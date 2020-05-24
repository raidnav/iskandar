package helper

import (
	"github.com/go-playground/assert/v2"
	"testing"
	"time"
)

func TestToMillis(t *testing.T) {
	testedTime := time.Date(2020, 5, 26, 1, 0, 0, 0, time.UTC)
	assert.Equal(t, TimeToMillis(testedTime), int64(1590454800000))
}
