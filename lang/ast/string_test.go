// Use of this source code is governed by a BSD-style license that can be found
// in the LICENSE file.

package ast_test

import (
	"testing"

	"github.com/google/puffs/lang/parse"
	"github.com/google/puffs/lang/token"
)

func TestString(t *testing.T) {
	const filename = "test.puffs"
	testCases := []string{
		"1",
		"x",

		"f()",
		"f(i)",
		"f(i, j)",
		"f(i, j) + 1",
		"f(i, j)(k)",
		"f(i, j)(k, l, m + 2) + 3",

		"x[i]",
		"x[i][j]",
		"x[i:j]",
		"x[i](j)",
		"x(i)[j]",

		"x.y",
		"x.y.z.a",

		"+x",
		"-(x + y)",
		"not x",
		"not not x",
		"+++-x",

		"x + 42",
		"x and (y < z)",
		"x & (y as u8)",
		"x * ((a / b) - (i / j))",

		"x + y + z",
		"x + (i * j.k[l] * (-m << 4) * (n & o(p, q[:r.s + 5]))) + z",

		"x as bool",
		"x as u32",
		"x as T",
		"x as T",
		"x as T[i:]",
		"x as T[:j]",
		"x as T[i:j]",
		"x as pkg.T",
		"x as ptr T",
		"x as [4] T",
		"x as [8 + (2 * N)] ptr [4] ptr pkg.T[i:j]",
	}

	idMap := &token.IDMap{}
	for _, tc := range testCases {
		tokens, _, err := token.Tokenize(idMap, filename, []byte(tc))
		if err != nil {
			t.Errorf("Tokenize(%q): %v", tc, err)
			continue
		}
		expr, err := parse.ParseExpr(idMap, filename, tokens)
		if err != nil {
			t.Errorf("ParseExpr(%q): %v", tc, err)
			continue
		}
		got := expr.String(idMap)
		if got != tc {
			t.Errorf("got %q, want %q", got, tc)
			continue
		}
	}
}