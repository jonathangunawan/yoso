package yoso

type Writer interface {
	Write([]string) error
	GetResultFiles() []string
}

type CSV interface {
	Write(record []string) error
	Flush()
}
