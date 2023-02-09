package gomcoponentizer

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

func parse(fileName string) (*html.Node, error) {
	stream, err := os.Open(fileName) // #nosec
	if err != nil {
		return nil, fmt.Errorf("while parsing HTML: %w", err)
	}

	return html.Parse(stream)
}
