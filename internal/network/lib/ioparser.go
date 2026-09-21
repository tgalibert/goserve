package lib

import (
	"io"
	"strings"
)

type IoParser struct {
	Raw	io.Reader
}

func NewIoParser(raw io.Reader) *IoParser {
	return &IoParser{
		Raw: raw,
	}
}

func (i *IoParser) ReadString(delimiter ...string) []string {
	buffer := make([]byte, 1024)
	var sb strings.Builder

	for {
		n, err :=i.Raw.Read(buffer)

		if n > 0 {
			sb.Write(buffer[:n])
		}
		content := sb.String()

		if strings.Contains(content, "\r\n\r\n") {                                                                                                          
			break                                                                                                                                               
		}                                                                                                                                                   
                                                                                                                                                         
		if err != nil {
			if err == io.EOF {
				break
			}
		}

	}
	if len(delimiter) > 0 && delimiter[0] != "" {
		return strings.SplitN(sb.String(), delimiter[0], -1)
	} else {
		return []string{sb.String()}
	}
}