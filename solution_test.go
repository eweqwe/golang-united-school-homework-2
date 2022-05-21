package square

import "testing"

type testSuiteData struct {
	Expected float64
	SidesNum intCustomType
	SideLen  float64
}

func TestCalcSquare(t *testing.T) {

	testSuites := []testSuiteData{
		{
			Expected: 314.1592653589793,
			SidesNum: 0,
			SideLen:  10,
		},
		{
			Expected: 43.30127018922193,
			SidesNum: 3,
			SideLen:  10,
		},
		{
			Expected: 100,
			SidesNum: 4,
			SideLen:  10,
		},
		{
			Expected: 0,
			SidesNum: 5,
			SideLen:  10,
		},
	}

	for _, value := range testSuites {
		actual := CalcSquare(value.SideLen, value.SidesNum)

		if actual != value.Expected {
			t.Errorf("CalcSquare() returns: \"%v\". Expected result: \"%v\"", actual, value.Expected)
		}
	}

}
