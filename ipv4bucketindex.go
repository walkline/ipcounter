package ipcounter

type IPv4BucketIndex struct {
}

func NewIPv4BucketIndex() *IPv4BucketIndex {
	return &IPv4BucketIndex{}
}

func (*IPv4BucketIndex) Add(ip string) error {
	return nil
}

func (*IPv4BucketIndex) Len() (int64, error) {
	return 0, nil
}
