// Copyright (c) 2017 Hervé Gouchet. All rights reserved.
// Use of this source code is governed by the MIT License
// that can be found in the LICENSE file.

package charset_test

import (
	"fmt"
	"testing"

	"github.com/rvflash/charset"
)

func ExampleFormatUtf8() {
	s := charset.FormatUtf8("Warning: ⚠ 1 minute left ⌚ before boom 💥 ➡ click on the button !")
	fmt.Println(s)
	// Output: Warning: ⚠ 1 minute left ⌚ before boom  ➡ click on the button !
}

func TestFormatUtf8(t *testing.T) {
	var dt = []struct {
		in, out string
	}{
		{
			in:  "Hello, 世界",
			out: "Hello, 世界",
		},
		{
			in:  "Oups: �",
			out: "Oups: ",
		},
		{
			in:  "Warning: ⚠ 1 minute left ⌚ before boom 💥 ➡ click on the button !",
			out: "Warning: ⚠ 1 minute left ⌚ before boom  ➡ click on the button !",
		},
	}
	for i, tt := range dt {
		if out := charset.FormatUtf8(tt.in); out != tt.out {
			t.Errorf("%d. error mismatch: got=%q, exp=%q", i, out, tt.out)
		}
	}
}

func TestFormatUtf8mb4(t *testing.T) {
	var dt = []struct {
		in, out string
	}{
		{
			in:  "Hello, 世界",
			out: "Hello, 世界",
		},
		{
			in:  "Oups: �",
			out: "Oups: ",
		},
		{
			in:  "Warning: ⚠ 1 minute left ⌚ before boom 💥 ➡ click on the button !",
			out: "Warning: ⚠ 1 minute left ⌚ before boom 💥 ➡ click on the button !",
		},
	}
	for i, tt := range dt {
		if out := charset.FormatUtf8mb4(tt.in); out != tt.out {
			t.Errorf("%d. error mismatch: got=%q, exp=%q", i, out, tt.out)
		}
	}
}
