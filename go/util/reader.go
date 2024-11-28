package util

import (
	"bufio"
	"bytes"
	"fmt"
	"iter"
	"os"
	"strings"
)

func ReadText(text string, delim byte) iter.Seq[string] {
	return func(yield func(string) bool) {
		split := strings.Split(text, string(delim))
		for _, elem := range split {
			if !yield(elem) {
				return
			}
		}
	}
}

func ReadFile(path string, delim byte) iter.Seq[string] {
	return func(yield func(string) bool) {
		f, err := os.Open(path)

		if err != nil {
			panic(err)
		}

		defer f.Close()

		scanner := bufio.NewScanner(f)

		split := func(data []byte, atEOF bool) (advance int, token []byte, err error) {
			i := bytes.IndexByte(data, delim)
			if i == -1 {
				if !atEOF {
					return 0, nil, nil
				}
				// If we have reached the end, return the last token.
				return 0, data, bufio.ErrFinalToken
			}

			// Otherwise, return the token before the comma.
			return i + 1, data[:i], nil
		}

		scanner.Split(split)

		for scanner.Scan() {
			w := scanner.Text()

			if !yield(w) {
				return
			}
		}

		if err := scanner.Err(); err != nil {
			fmt.Fprintln(os.Stderr, "An error occured:", err)
		}
	}
}
