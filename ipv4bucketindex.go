package ipcounter

import (
	"errors"
	"strconv"
	"strings"
	"sync"
)

var (
	ErrBadIPAddress = errors.New("bad ip address")
)

const (
	bucketSize = 0x100 // all possible combinations for byte
)

type zeroBytesType struct{}

// IPv4BucketIndex implements index for v4 ips, uses buckets of maps
type IPv4BucketIndex struct {
	// buckets used to reduce number of thread locks.
	//
	// Key for bucket consist from last byte of ip address.
	// Used last byte to achieve more equal load to buckets.
	//
	// Key sample:
	// 		'176.32.103.200' => 200 => 0xC8
	buckets     [bucketSize]map[string]zeroBytesType
	bucketsLock [bucketSize]sync.RWMutex
}

func NewIPv4BucketIndex() *IPv4BucketIndex {
	index := &IPv4BucketIndex{}
	for i := 0; i < bucketSize; i++ {
		index.buckets[i] = map[string]zeroBytesType{}
	}
	return index
}

func (i *IPv4BucketIndex) Add(ip string) error {
	k, err := i.bucketKeyByIP(ip)
	if err != nil {
		return err
	}

	i.bucketsLock[k].Lock()

	// built-in maps hash function is faster
	// then converting ip to integer
	i.buckets[k][ip] = zeroBytesType{}

	// better to not use defer here
	// because it would take additional 20ns
	i.bucketsLock[k].Unlock()

	return nil
}

func (i *IPv4BucketIndex) Len() (int64, error) {
	counter := int64(0)
	for itr := range i.buckets {
		i.bucketsLock[itr].RLock()
		counter += int64(len(i.buckets[itr]))
		i.bucketsLock[itr].RUnlock()
	}
	return counter, nil
}

func (*IPv4BucketIndex) bucketKeyByIP(ip string) (int, error) {
	i := strings.LastIndexByte(ip, '.')
	if i < 0 {
		return i, ErrBadIPAddress
	}

	key, err := strconv.Atoi(ip[i+1:])
	if err != nil {
		return -1, ErrBadIPAddress
	}

	return key, nil
}
