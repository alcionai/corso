package ics

import (
	"bytes"
	"regexp"
	"strings"

	"golang.org/x/net/html"
)

// removeTrailingWhiltesapce removes trailing whitespace from every
// line in a string.
func removeTrailingWhiltesapce(in string) string {
	out := ""
	lines := strings.Split(in, "\n")

	for _, line := range lines {
		out += strings.TrimRight(line, " ") + "\n"
	}

	return strings.TrimSpace(out)
}

// MS converts content from HTML to text when then have to export it
// to other formats. This is a best effort reproduction of what they do.
func HTMLToText(in string) (string, error) {
	in = strings.ReplaceAll(in, "\r\n", "\n")
	z := html.NewTokenizer(bytes.NewReader([]byte(in)))
	depth := 0
	out := ""
	lastTag := ""
	lastLink := ""

	for {
		tt := z.Next()

		switch tt {
		case html.ErrorToken:
			// TODO(meain): usually EOF, but handle other error
			return removeTrailingWhiltesapce(out), nil
		case html.TextToken:
			if depth > 0 {
				origText := string(z.Text())
				if (origText == " " || origText == "\n") && len(out) > 0 &&
					(out[len(out)-1] != '\n' && out[len(out)-1] != ' ') {
					out += " "
				}

				origText = strings.ReplaceAll(origText, "\n", " ")
				text := regexp.MustCompile(`\s+`).ReplaceAllString(origText, " ")

				if len(strings.TrimSpace(text)) == 0 {
					continue
				}

				if origText[0] == ' ' && len(out) > 0 && out[len(out)-1] != ' ' {
					out += " "
				}

				out += strings.TrimSpace(text)

				if origText[len(origText)-1] == ' ' && len(out) > 0 && out[len(out)-1] != ' ' {
					out += " "
				}
			}
		case html.StartTagToken, html.EndTagToken:
			tn, _ := z.TagName()
			lastTag = string(tn)

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
					if len(lastLink) > 0 {
						out += "<" + lastLink + ">"
						lastLink = ""
					}
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
			case "img":
				if tt == html.StartTagToken {
					cid := ""

					for {
						key, val, more := z.TagAttr()
						if string(key) == "src" {
							cid = string(val)
							break
						}

						if !more {
							break
						}
					}

					out += "[" + cid + "]"
				}
			}
		}
	}
}
