// Copyright (c) 2017 Hervé Gouchet. All rights reserved.
// Use of this source code is governed by the MIT License
// that can be found in the LICENSE file.

package charset

import (
	"bytes"
	"unicode/utf8"
)

// Utf8 returns a string with a maximum of three bytes per character.
func Utf8(in string) string {
	return stripBytes(in, 3)
}

// Utf8mb4 returns a string with a maximum of four bytes per character.
func Utf8mb4(in string) string {
	return stripBytes(in, 4)
}

func stripBytes(in string, max int) string {
	var b bytes.Buffer
	for len(in) > 0 {
		r, size := utf8.DecodeRuneInString(in)
		in = in[size:]
		if r == utf8.RuneError {
			continue
		}
		if size > max {
			continue
		}
		b.WriteRune(r)
	}
	return b.String()
}
