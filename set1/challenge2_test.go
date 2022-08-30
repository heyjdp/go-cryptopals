package set1

import "testing"

type TestChallenge2Data struct {
	in       string
	xor      string
	expected string
}

func TestChallenge2(t *testing.T) {
	testData := []TestChallenge2Data{
		{
			"1c0111001f010100061a024b53535009181c",
			"686974207468652062756c6c277320657965",
			"746865206b696420646f6e277420706c6179",
		},
		{
			"1ce0f9663ba06c67760709354a6bc55327a8",
			"ed81352672aab6649d915d8bf1e12dd0e2e3",
			"f161cc40490ada03eb9654bebb8ae883c54b",
		},
		{
			"44c653d9a75cdaa67afa39e75de4b853bd82",
			"19b1a326de3001e1da460d59f9c6a8042ddb",
			"5d77f0ff796cdb47a0bc34bea42210579059",
		},
		{
			"dad949538355e9c526a800382f9d054ae1b2",
			"afaf545967ddfd80178a4b81fffd63bbc7c1",
			"75761d0ae488144531224bb9d06066f12673",
		},
	}

	for _, testCase := range testData {
		in := testCase.in
		xor := testCase.xor
		out := Challenge2(in, xor)
		if out != testCase.expected {
			t.Errorf("FAILED Challenge2\n- Expected: %s\n- Got     : %s", testCase.expected, out)
		} else {
			t.Logf("PASSED Challenge2\n- Expected: %s\n- Got     : %s", testCase.expected, out)
		}
	}
}
