package set1

import (
	"testing"
)

type TestData struct {
	in       string
	expected string
}

func TestChallenge1(t *testing.T) {
	testData := []TestData{
		{
			"49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d",
			"SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t",
		},
		{
			"a8d399e83503d88a75e9bf2faf3b3b80bd48ec6aabf6a213c7e192f6007d64f52709a6fc3170e75ff133b6230d4468d7",
			"qNOZ6DUD2Ip16b8vrzs7gL1I7Gqr9qITx+GS9gB9ZPUnCab8MXDnX/EztiMNRGjX",
		},
		{
			"a244d4210e3d28d844ec90b1a9a2d3e20ef1204a427dded9c27f9033dbc7534dbf8fafa559e5bb0906f1e822d7b52615",
			"okTUIQ49KNhE7JCxqaLT4g7xIEpCfd7Zwn+QM9vHU02/j6+lWeW7CQbx6CLXtSYV",
		},
		{
			"57a37ba556434991937c9d36e9b9bf4d41253d3d27a73fb393eded6ce6e3390ac4a1bfcd76e76d915a195a18d7e7d4b5",
			"V6N7pVZDSZGTfJ026bm/TUElPT0npz+zk+3tbObjOQrEob/NdudtkVoZWhjX59S1",
		},
	}

	for _, testCase := range testData {
		in := testCase.in
		out := Challenge1(in)
		if out != testCase.expected {
			t.Errorf("FAILED Challenge1\n- Expected: %s\n- Got     : %s", testCase.expected, out)
		} else {
			t.Logf("PASSED Challenge1\n- Expected: %s\n- Got     : %s", testCase.expected, out)

		}
	}
}
