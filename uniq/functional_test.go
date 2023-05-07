package uniq

import (
	"bytes"
	"strings"
	"testing"
)

func TestFunctional(t *testing.T) {
	testCases := []struct {
		name        string
		input       string
		expected    string
		expectedErr error
		options     *Options
	}{
		{
			name:  "no options",
			input: "foo\nbar\nbaz\n",
			expected: `foo
bar
baz
`,
			options: &Options{},
		},
		{
			name:  "num_fields",
			input: "foo 1\nbar 2\nbaz 3\n",
			expected: `1
2
3
`,
			options: &Options{
				Num_fields: 1,
			},
		},
		{
			name:  "num_chars",
			input: "foo 123\nbar 456\nbaz 789\n",
			expected: `23
56
89
`,
			options: &Options{
				Num_chars: 2,
			},
		},
		{
			name:  "ignore_case",
			input: "foo\nBar\nbaz\n",
			expected: `foo
baz
`,
			options: &Options{
				RegiPtr: true,
			},
		},
		{
			name:  "copy",
			input: "foo\nbar\nfoo\n",
			expected: `2 foo
1 bar
1 foo
`,
			options: &Options{
				CopyPtr: true,
			},
		},
		{
			name:  "double",
			input: "foo\nbar\nfoo\n",
			expected: `foo
foo
`,
			options: &Options{
				DoublePtr: true,
			},
		},
		{
			name:  "unique",
			input: "foo\nbar\nfoo\n",
			expected: `bar
`,
			options: &Options{
				UniqPtr: true,
			},
		},
		{
			name:  "input_file",
			input: "",
			expected: `foo
bar
baz
`,
			options: &Options{
				InputFile: "testdata/input.txt",
			},
		},
		{
			name:        "output_file",
			input:       "foo\nbar\nbaz\n",
			expected:    "",
			expectedErr: nil,
			options: &Options{
				OutputFile: "testdata/output.txt",
			},
		},
		{
			name:        "too_many_args",
			input:       "",
			expected:    "",
			expectedErr: ErrTooMuchFlagArgs,
			options:     &Options{},
		},
		{
			name:        "uncombined_params",
			input:       "",
			expected:    "",
			expectedErr: ErrUncombinedParams,
			options: &Options{
				CopyPtr:   true,
				DoublePtr: true,
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			var output bytes.Buffer
			input := strings.NewReader(tc.input)
			Functional(tc.options, &output, input)
		})
	}
}

func TestStrsIsEqual(t *testing.T) {
	previous := "foo"
	current := "foo"
	opt := Options{RegiPtr: false}

	if !StrsIsEqual(&previous, &current, &opt) {
		t.Errorf("expected true, but got false")
	}

	previous = "FOO"
	current = "foo"
	opt.RegiPtr = false

	if StrsIsEqual(&previous, &current, &opt) {
		t.Errorf("expected false, but got true")
	}
}

func TestDeleteFieldsHandler(t *testing.T) {
	str := "foo bar baz"
	opt := Options{Num_fields: 1}

	result, err := DeleteFieldsHandler(str, &opt)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	expected := "bar baz"

	if result != expected {
		t.Errorf("expected %q, but got %q", expected, result)
	}
}

func TestDeleteCharsHandler(t *testing.T) {
	str := "foobar"
	opt := Options{Num_chars: 3}

	result, err := DeleteCharsHandler(str, &opt)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	expected := "bar"

	if result != expected {
		t.Errorf("expected %q, but got %q", expected, result)
	}
}
