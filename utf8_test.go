// Copyright (c) 2017 Hervé Gouchet. All rights reserved.
// Use of this source code is governed by the MIT License
// that can be found in the LICENSE file.

package charset_test

import (
	"testing"

	"github.com/rvflash/charset"
)

func TestUtf8(t *testing.T) {
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
		if out := charset.Utf8(tt.in); out != tt.out {
			t.Errorf("%d. error mismatch: got=%q, exp=%q", i, out, tt.out)
		}
	}
}

func TestUtf8mb4(t *testing.T) {
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
		if out := charset.Utf8mb4(tt.in); out != tt.out {
			t.Errorf("%d. error mismatch: got=%q, exp=%q", i, out, tt.out)
		}
	}
}
