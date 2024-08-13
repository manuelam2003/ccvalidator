package luhn

import "testing"

func TestCheckLuhn(t *testing.T) {
	testCases := []struct {
		desc     string
		input    string
		expected bool
	}{
		{
			desc:     "Test1",
			input:    "17893729974",
			expected: true,
		},
		{
			desc:     "Test2",
			input:    "17893729973",
			expected: false,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			res := CheckLuhn(tC.input)

			if tC.expected != res {
				t.Errorf("got %t want %t given, %s", res, tC.expected, tC.input)
			}
		})
	}

}
