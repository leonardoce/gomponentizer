package gomcoponentizer

import (
	"fmt"
	"io"
	"strings"

	"golang.org/x/net/html"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func spit(node *html.Node, output io.Writer) {
	switch node.Type {
	case html.DocumentNode:
		spit(node.FirstChild, output)

	case html.CommentNode:
		fmt.Fprintf(output, "// %s\n", node.Data)

	case html.ElementNode:
		tagName := cases.Title(language.English, cases.Compact).String(node.Data)
		fmt.Fprintf(output, "html.%s(\n", tagName)
		spitAttrs(node, output)

		for child := node.FirstChild; child != nil; child = child.NextSibling {
			spit(child, output)
		}
		fmt.Fprintf(output, "),\n")

	case html.TextNode:
		text := strings.Trim(node.Data, " \n")
		if len(text) > 0 {
			fmt.Fprintf(output, "gomponents.Text(%s),\n", escapeString(text))
		}

	case html.DoctypeNode:

	default:
		fmt.Printf("unmanaged node type: %v\n", node.Type)
	}
}

func spitAttrs(node *html.Node, output io.Writer) {
	for _, attr := range node.Attr {
		attrName := cases.Title(language.English, cases.Compact).String(attr.Key)

		if strings.Contains(attrName, ":") || strings.Contains(attrName, "@") || strings.Contains(attrName, "-") {
			fmt.Fprintf(output, "gomponents.Attr(%s, %s),\n", escapeString(attr.Key), escapeString(attr.Val))
		} else {
			fmt.Fprintf(output, "html.%s(%s),\n", attrName, escapeString(attr.Val))
		}
	}
}

func escapeString(name string) string {
	return fmt.Sprintf("\"%s\"", strings.ReplaceAll(name, "\"", "\"\""))
}
