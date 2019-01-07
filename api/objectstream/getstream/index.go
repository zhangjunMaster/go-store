package getstream

import (
	"fmt"
	"io"
	"net/http"
)

type GetStream struct {
	reader io.Reader
}

// 从真是的url获取文件
func newGetStream(url string) (*GetStream, error) {
	r, err := http.Get(url)

	if err != nil {
		return nil, err
	}

	if r.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("dataserver return code %d", r.StatusCode)
	}
	return &GetStream{r.Body}, nil
}

func NewGetStream(server, object string) (*GetStream, error) {
	if server == "" || object == "" {
		return nil, fmt.Errorf("参数错误")
	}
	url := "http://" + server + "/objects/" + object
	return newGetStream(url)
}

func (r *GetStream) Read(p []byte) (n int, err error) {
	return r.reader.Read(p)
}
