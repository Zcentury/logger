package writer

type FileWriter struct {
	Name         string
	AutoDateName bool // 是否自动设置文件名
}

func (w *FileWriter) Write(p []byte) (n int, err error) {
	return 0, nil
}
