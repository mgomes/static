package static

import (
	"io"
	"io/ioutil"

	"github.com/russross/blackfriday.v2"
)

// Markdown returns a reader transformed to markdown.
func Markdown(r io.Reader) io.ReadCloser {
	pr, pw := io.Pipe()

	go func() {
		b, err := ioutil.ReadAll(r)
		if err != nil {
			pw.CloseWithError(err)
			return
		}

		pw.Write(blackfriday.Run(b))
		pw.Close()
	}()

	return pr
}
