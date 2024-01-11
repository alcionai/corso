package ics

import (
	"bytes"
	"fmt"
	"strings"

	"golang.org/x/net/html"
)

// MS converts content from HTML to text when then have to export it
// to other formats. This is a best effort reproduction of what they do.
func HTMLToText(in string) (string, error) {
	out := ""
	z := html.NewTokenizer(bytes.NewReader([]byte(in)))
	depth := 0
	lastTag := ""
	lastLink := ""

	for {
		tt := z.Next()
		switch tt {
		case html.ErrorToken:
			// TODO(meain): usually EOF, but handle other error
			fmt.Println(out)
			return out, nil
		case html.TextToken:
			if depth > 0 {
				text := string(z.Text())
				text = strings.ReplaceAll(text, "\n", "")
				text = strings.ReplaceAll(text, "  ", " ")

				if len(strings.TrimSpace(text)) == 0 {
					continue
				}

				out += text

				if lastTag == "a" && len(lastLink) > 0 {
					out += "<" + lastLink + ">"
				}
			}
		case html.StartTagToken, html.EndTagToken:
			tn, _ := z.TagName()
			lastTag = string(tn)

			// TODO(meain): what about images?
			switch lastTag {
			case "a":
				if tt == html.StartTagToken {
					link := ""

					for {
						key, val, more := z.TagAttr()
						if string(key) == "href" {
							link = string(val)
							break
						}

						if !more {
							break
						}
					}

					lastLink = link
				} else {
					lastLink = ""
				}
			case "br":
				if tt == html.StartTagToken {
					out += "\n"
				}
			case "div":
				if tt == html.StartTagToken {
					depth++
				} else {
					depth--
				}

				if len(out) > 0 && out[len(out)-1] != '\n' {
					out += "\n"
				}
			}
		}
	}
}
