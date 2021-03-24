package telex

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCombine(t *testing.T) {
	tests := []struct {
		name  string
		base  string
		extra string
		want  string
	}{
		{
			name:  "á",
			base:  "a",
			extra: "s",
			want:  "á",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := combine(tc.base, tc.extra)
			assert.Equal(t, tc.want, got)
		})
	}
}
