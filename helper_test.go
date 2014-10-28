package filters

import "testing"

func TestNormalize(t *testing.T) {
	actual := []float64{0.0, 1.0, 2.0, 3.0, 4.0}
	expected := []float64{0.0, 0.1, 0.2, 0.3, 0.4}
	Normalize(&actual)

	sum := 0.0
	for i, a := range actual {
		e := expected[i]
		if a != e {
			t.Errorf("At the %d of normalized slice should be %f, but actual %f", i, e, a)
		}
		sum += a
	}
	if sum != 1 {
		t.Errorf("The sum of all elements of actual sould be 1, but actual %f", sum)
	}
}
