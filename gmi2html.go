// Copyright (c) Alisdair MacLeod <copying@alisdairmacleod.co.uk>
//
// Permission to use, copy, modify, and/or distribute this software for any
// purpose with or without fee is hereby granted.
//
// THE SOFTWARE IS PROVIDED "AS IS" AND THE AUTHOR DISCLAIMS ALL WARRANTIES WITH
// REGARD TO THIS SOFTWARE INCLUDING ALL IMPLIED WARRANTIES OF MERCHANTABILITY
// AND FITNESS. IN NO EVENT SHALL THE AUTHOR BE LIABLE FOR ANY SPECIAL, DIRECT,
// INDIRECT, OR CONSEQUENTIAL DAMAGES OR ANY DAMAGES WHATSOEVER RESULTING FROM
// LOSS OF USE, DATA OR PROFITS, WHETHER IN AN ACTION OF CONTRACT, NEGLIGENCE OR
// OTHER TORTIOUS ACTION, ARISING OUT OF OR IN CONNECTION WITH THE USE OR
// PERFORMANCE OF THIS SOFTWARE.

package main

import (
	"bufio"
	"fmt"
	"html"
	"log"
	"os"
	"strings"
)

func main() {
	var preformatted, list, blockquote bool

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		line := scanner.Text()
		if blockquote && !strings.HasPrefix(line, ">") {
			fmt.Println("</blockquote>")
			blockquote = false
		}
		if list && !strings.HasPrefix(line, "* ") {
			fmt.Println("</ul>")
			list = false
		}
		switch {
		case line == "```":
			if preformatted {
				fmt.Println("</pre>")
			} else {
				fmt.Println("<pre>")
			}
			preformatted = !preformatted
		case preformatted:
			fmt.Println(html.EscapeString(line))
		case strings.HasPrefix(line, "* "):
			if !list {
				fmt.Println("<ul>")
				list = true
			}
			fmt.Printf("<li>%s</li>\n", strings.TrimPrefix(line, "* "))
		case strings.HasPrefix(line, ">"):
			if !blockquote {
				fmt.Println("<blockquote>")
				blockquote = true
			}
			fmt.Printf("<p>%s</p>\n", strings.TrimPrefix(line, ">"))
		case strings.HasPrefix(line, "###"):
			fmt.Printf("<h3>%s</h3>\n", strings.TrimSpace(strings.TrimPrefix(line, "###")))
		case strings.HasPrefix(line, "##"):
			fmt.Printf("<h2>%s</h2>\n", strings.TrimSpace(strings.TrimPrefix(line, "##")))
		case strings.HasPrefix(line, "#"):
			fmt.Printf("<h1>%s</h1>\n", strings.TrimSpace(strings.TrimPrefix(line, "#")))
		case strings.HasPrefix(line, "=>"):
			line = strings.TrimSpace(strings.TrimPrefix(line, "=>"))
			fields := strings.Fields(line)
			if len(fields) == 0 {
				break
			}
			link := fields[0]
			text := fields[0]
			if len(fields) > 1 {
				text = strings.TrimSpace(strings.TrimPrefix(line, link))
			}
			fmt.Printf("<p><a href=%q>%s</a></p>\n", link, text)
		case line == "":
			fmt.Println("<br />")
		default:
			fmt.Printf("<p>%s</p>\n", line)
		}
	}
	if err := scanner.Err(); err != nil {
		log.Println(err)
	}
}
