package objectstream

import (
	"io"
	"net/http"
)

type PutStream struct {
	writer *io.PipeWriter
	c      chan error
}

// server指的是选中的server节点，是ip地址
func NewPutStream(server, object string) *PutStream {
	reader, writer := io.Pipe()
	c := make(chan error)
	go func() {
		request, err := http.NewRequest("PUT", "http://"+server+"/objects/"+object, reader)
		if err != nil {
			c <- err
		}
		client := http.Client{}
		res, err := client.Do(request)
		if err != nil && res.StatusCode != http.StatusOK {
			c <- err
		}

	}()

	return &PutStream{writer, c}
}

// PutStream实现Write方法，就实现了io.Writer接口
func (ps *PutStream) Write(p []byte) (n int, err error) {
	return ps.writer.Write(p)
}

// PutStream实现Close方法，就实现了io.WriterClose接口

func (ps *PutStream) Close() error {
	return ps.writer.Close()
}
