package ipcounter_test

import (
	"fmt"
	"math/rand"
	"testing"
	"time"

	. "github.com/walkline/ipcounter"
)

func TestLenForIPv4BucketIndex(t *testing.T) {
	for _, c := range []struct {
		ips         []string
		expectedLen int64
	}{
		{
			[]string{},
			0,
		},
		{
			[]string{"1.1.1.1"},
			1,
		},
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

func TestAddForIPv4BucketIndex(t *testing.T) {
	for _, c := range []struct {
		ip          string
		expectedErr error
	}{
		{
			"1.1.1.1",
			nil,
		},
		{
			"240.240.240.240",
			nil,
		},
		{
			"",
			ErrBadIPAddress,
		},
		{
			"0.0.0.X",
			ErrBadIPAddress,
		},
	} {
		index := NewIPv4BucketIndex()

		err := index.Add(c.ip)
		if err != c.expectedErr {
			t.Errorf("expected err - %+v, actual err - %+v", c.expectedErr, err)
		}
	}
}

func benchmarkAddForIPv4BucketIndex(indexSize int, b *testing.B) {
	rand.Seed(time.Now().UnixNano())

	index := NewIPv4BucketIndex()
	for _, ip := range generateIPs(indexSize) {
		index.Add(ip)
	}

	ip := generateIP()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		index.Add(ip)
	}
}

func BenchmarkAddForIPv4BucketIndex100(b *testing.B)     { benchmarkAddForIPv4BucketIndex(100, b) }
func BenchmarkAddForIPv4BucketIndex10000(b *testing.B)   { benchmarkAddForIPv4BucketIndex(10000, b) }
func BenchmarkAddForIPv4BucketIndex100000(b *testing.B)  { benchmarkAddForIPv4BucketIndex(100000, b) }
func BenchmarkAddForIPv4BucketIndex1000000(b *testing.B) { benchmarkAddForIPv4BucketIndex(1000000, b) }

func randomInt(maxValue int) int {
	min := 1
	max := maxValue
	return rand.Intn(max-min+1) + min
}

func generateIP() string {
	return fmt.Sprintf("%d.%d.%d.%d", randomInt(1), randomInt(255), randomInt(255), randomInt(255))
}

func generateIPs(count int) []string {
	ips := make([]string, count)
	for i := range ips {
		ips[i] = generateIP()
	}
	return ips
}
