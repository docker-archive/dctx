package httpsrvctx

import (
	"reflect"
	"testing"

	"github.com/docker/dctx"
)

func TestWithVars(t *testing.T) {
	vars := map[string]string{
		"foo": "asdf",
		"bar": "qwer",
	}

	ctx := WithVars(dctx.Background(), vars)
	for _, testcase := range []struct {
		key      string
		expected interface{}
	}{
		{
			key:      "vars",
			expected: vars,
		},
		{
			key:      "vars.foo",
			expected: "asdf",
		},
		{
			key:      "vars.bar",
			expected: "qwer",
		},
	} {
		v := ctx.Value(testcase.key)

		if !reflect.DeepEqual(v, testcase.expected) {
			t.Fatalf("%q: %v != %v", testcase.key, v, testcase.expected)
		}
	}
}
