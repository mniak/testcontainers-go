package wait

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestWaitHTTP_TimeoutAccessors(t *testing.T) {
	strategy := ForHTTP("/test")

	strategy.timeout = time.Second * 2
	assert.Equal(t, time.Second*2, strategy.timeout)

	strategy.WithTimeout(time.Second * 3)
	assert.Equal(t, time.Second*3, strategy.timeout)

	// Deprecated
	strategy.WithStartupTimeout(time.Second * 4)
	assert.Equal(t, time.Second*4, strategy.timeout)
}
