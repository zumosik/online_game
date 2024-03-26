package models

import (
	"bytes"
	"math"
	"reflect"
	"testing"
)

func TestVecor(t *testing.T) {
	nums := [][]float64{
		{2.3, 4.3}, {2, 3}, {666.77, 9.2}, {0.1, -99.2}, {math.Inf(2), math.Inf(3)},
	}

	for i := 0; i < len(nums); i++ {
		var vec Vector
		vec.X = nums[i][0]
		vec.Y = nums[i][1]

		var buf bytes.Buffer
		vec.Serialize(&buf)

		t.Log(buf.Bytes())

		var vec2 Vector
		err := vec2.Deserialize(&buf)
		if err != nil {
			t.Error(err)
			t.Fail()
		}

		if reflect.DeepEqual(vec, vec2) {
			t.Log("ok")
		} else {
			t.Errorf("got: %v, want %v", vec2, vec)
			t.Fail()
		}
	}
}
