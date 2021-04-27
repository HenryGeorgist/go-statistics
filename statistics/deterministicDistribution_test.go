package statistics

import (
	"testing"
)

func TestDeterministicDistribution_InvCDF(t *testing.T) {
	val := 0.5
	d, _ := InitDeterministic(val)
	probability := 0.8
	var expected float64 = 0.5
	got := d.InvCDF(probability)
	if expected != got {
		t.Errorf("d.InvCDF(probability): %v; expected: %v", got, expected)
	}
}

func TestDeterministicDistribution_PDFIsValue(t *testing.T) {
	val := 0.5
	d, _ := InitDeterministic(val)
	var expected float64 = 1.0
	got := d.PDF(val)
	if expected != got {
		t.Errorf("d.InvCDF(probability): %v; expected: %v", got, expected)
	}

}

func TestDeterministicDistribution_PDFIsNotValue(t *testing.T) {
	val := 0.5
	d, _ := InitDeterministic(val)
	notValue := 4.0
	var expected float64 = 0.0
	got := d.PDF(notValue)
	if expected != got {
		t.Errorf("d.PDF(notValue): %v; expected: %v", got, expected)
	}
}

func TestDeterministicDistribution_CDFIsValue(t *testing.T) {
	val := 0.5
	d, _ := InitDeterministic(val)
	var expected float64 = 1.0
	got := d.CDF(val)
	if expected != got {
		t.Errorf("d.CDF(val): %v; expected: %v", got, expected)
	}

}

func TestDeterministicDistribution_CDFGreaterThanValue(t *testing.T) {
	val := 0.5
	d, _ := InitDeterministic(val)
	greaterValue := 4.0
	var expected float64 = 1.0
	got := d.CDF(greaterValue)
	if expected != got {
		t.Errorf("d.PDF(notValue): %v; expected: %v", got, expected)
	}
}

func TestDeterministicDistribution_CDFLessThanValue(t *testing.T) {
	val := 0.5
	d, _ := InitDeterministic(val)
	lesserValue := 0.4
	var expected float64 = 0.0
	got := d.CDF(lesserValue)
	if expected != got {
		t.Errorf("d.PDF(notValue): %v; expected: %v", got, expected)
	}
}
