package lib

import (
	"bytes"
	"io"
)

type IoParser struct {
	Raw	io.Reader
	buf []byte
}

func NewIoParser(raw io.Reader) *IoParser {
	return &IoParser{
		Raw: raw,
		buf: make([]byte,0),
	}
}

func (i *IoParser) ReadLine(delimiter ...string) (string, error) {
	temp := make([]byte, 4096)
	for {
		idx := bytes.Index(i.buf, []byte("\r\n"))
		if idx > -1 {
			line := string(i.buf[:idx])
			i.buf = i.buf[idx + 2:]
			return line, nil
		}
		n, err := i.Raw.Read(temp)
		i.buf = append(i.buf, temp[:n]...)		

		if err != nil {
			return "", err
		}
	}
}

func (i *IoParser) ReadBytes(length int) ([]byte, error) {
	temp := make([]byte, 4096)
	for {
		if len(i.buf) >= length {
			body := i.buf[:length]
			i.buf = i.buf[length:]
			return body, nil
		}
		n, err := i.Raw.Read(temp)
		i.buf = append(i.buf, temp[:n]...)
		if err != nil {
			return nil, err
		}
	}
}