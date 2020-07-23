package ipcounter_test

import (
	"testing"

	. "github.com/walkline/ipcounter"
)

func TestLenForIPv4BucketIndex(t *testing.T) {
	for _, c := range []struct {
		ips         []string
		expectedLen int64
	}{
		{
			[]string{"1.1.1.1", "2.2.2.42", "3.3.3.205"},
			3,
		},
		{
			[]string{"1.1.1.1", "2.2.2.42", "1.1.1.1"},
			2,
		},
	} {
		index := NewIPv4BucketIndex()
		for _, ip := range c.ips {
			if err := index.Add(ip); err != nil {
				t.Error(err)
			}
		}

		l, err := index.Len()
		if err != nil {
			t.Error(err)
		}

		if l != c.expectedLen {
			t.Errorf("expected result - %d, actual result - %d", c.expectedLen, l)
		}
	}
}
