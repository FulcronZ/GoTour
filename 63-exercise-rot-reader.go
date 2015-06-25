package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (reader rot13Reader) Read(p []byte) (int, error) {
	for {
		n, err := reader.r.Read(p)
		if err == nil {
			for i := 0; i < n; i++ {
				switch {
				case p[i] >= 'a' && p[i] <= 'z':
					p[i] += 13
					if p[i] > 'z' {	
						p[i] = p[i] - 'z' + 'a'	- 1
					}	
				case p[i] >= 'A' && p[i] <= 'Z':
					p[i] += 13
					if p[i] > 'Z' {
						p[i] = p[i] - 'Z' + 'A' - 1
					}	
				}
			}
			return n, nil
		}
		
		return 0, io.EOF
	}
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, r)	
}