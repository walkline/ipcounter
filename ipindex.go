package ipcounter

type IPIndex interface {
	Add(ip string) error
	Len() (int64, error)
}
