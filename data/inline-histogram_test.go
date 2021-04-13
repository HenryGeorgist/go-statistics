package data

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"testing"

	"github.com/HydrologicEngineeringCenter/go-statistics/statistics"
)

func TestAddObservationBinData(t *testing.T) {
	ih := Init(1.0, 0.0, 5.0)
	vals := []float64{.5, 1.5, 2.5, 3.5, 4.5}
	expected := []int64{1, 1, 1, 1, 1}
	for idx := range vals {
		ih.AddObservation(vals[idx])
	}
	for idx, val := range expected {
		if ih.GetBins()[idx] != val {
			t.Errorf("Bin(%d) = %d; expected %d", idx, ih.GetBins()[idx], val)
		}
	}
	fmt.Println(ih)
	b, err := json.Marshal(ih)
	if err != nil {
		panic(err)
	}
	s := string(b)
	fmt.Println(s)
}
func TestAddObservationBinEdges(t *testing.T) {
	ih := Init(1.0, 0.0, 5.0)
	vals := []float64{0, 1, 2, 3, 4, 5}
	expected := []int64{1, 1, 1, 1, 1, 1}
	for idx := range vals {
		ih.AddObservation(vals[idx])
	}
	for idx, val := range expected {
		if ih.GetBins()[idx] != val {
			t.Errorf("Bin(%d) = %d; expected %d", idx, ih.GetBins()[idx], val)
		}
	}
}
func TestAddObservationBinEdges_cruel(t *testing.T) {
	ih := Init(1.0, 0.0, 5.0)
	vals := []float64{0, 1, 2, 3, 4, 6}
	expected := []int64{1, 1, 1, 1, 1, 0, 1}
	for idx := range vals {
		ih.AddObservation(vals[idx])
	}
	for idx, val := range expected {
		if ih.GetBins()[idx] != val {
			t.Errorf("Bin(%d) = %d; expected %d", idx, ih.GetBins()[idx], val)
		}
	}
}
func TestAddObservationExceedUpper(t *testing.T) {
	ih := Init(1.0, 0.0, 5.0)
	vals := []float64{.5, 1.5, 2.5, 3.5, 4.5, 5.5}
	expected := []int64{1, 1, 1, 1, 1, 1}
	for idx := range vals {
		ih.AddObservation(vals[idx])
	}
	for idx, val := range expected {
		if ih.GetBins()[idx] != val {
			t.Errorf("Bin(%d) = %d; expected %d", idx, ih.GetBins()[idx], val)
		}
	}
}
func TestAddObservationBelowLower(t *testing.T) {
	ih := Init(1.0, 0.0, 5.0)
	vals := []float64{-.5, .5, 1.5, 2.5, 3.5, 4.5}
	expected := []int64{1, 1, 1, 1, 1, 1}
	for idx := range vals {
		ih.AddObservation(vals[idx])
	}
	for idx, val := range expected {
		if ih.GetBins()[idx] != val {
			t.Errorf("Bin(%d) = %d; expected %d", idx, ih.GetBins()[idx], val)
		}
	}
}
func TestAddObservationNormalDist(t *testing.T) {
	rand.Seed(1234)
	ih := Init(.01, -1.0, 1.0)
	n := statistics.NormalDistribution{Mean: 0, StandardDeviation: 1}
	iterations := 1000
	for i := 0; i < iterations; i++ {
		ih.AddObservation(n.InvCDF(rand.Float64()))
	}
	for _, val := range ih.GetBins() {
		fmt.Println(val)
	}
	fmt.Println(fmt.Sprintf("numbins %d", len(ih.GetBins())))
	fmt.Println(fmt.Sprintf("min %f", ih.pm.GetMin()))
	fmt.Println(fmt.Sprintf("max %f", ih.pm.GetMax()))
	fmt.Println(fmt.Sprintf("mean %f", ih.pm.GetMean()))
	fmt.Println(fmt.Sprintf("variance %f", ih.pm.GetSampleVariance()))
	fmt.Println(fmt.Sprintf("sample size %d", ih.pm.GetSampleSize()))
}
func TestAddObservationNormalDistConvergenceLeastStrict(t *testing.T) {
	rand.Seed(1234)
	ih := Init(.01, -1.0, 1.0)
	n := statistics.NormalDistribution{Mean: 0, StandardDeviation: 1}
	var convergence bool = false
	var iterations int64 = 1000
	for convergence != true {
		var i int64 = 0
		for i < iterations {
			ih.AddObservation(n.InvCDF(rand.Float64()))
			i++
		}
		convergence, iterations = ih.TestForConvergence(.05, .95, .95, .001) //upper confidence limit test, lower confidence limit test, confidenece, error tolerance
		fmt.Println(fmt.Sprintf("Computed %d estimated to need %d more iterations", ih.pm.GetSampleSize(), iterations))
	}
	fmt.Println("****Converged 95, 5******")
	fmt.Println(fmt.Sprintf("numbins %d", len(ih.GetBins())))
	fmt.Println(fmt.Sprintf("min %f", ih.pm.GetMin()))
	fmt.Println(fmt.Sprintf("max %f", ih.pm.GetMax()))
	fmt.Println(fmt.Sprintf("mean %f", ih.pm.GetMean()))
	fmt.Println(fmt.Sprintf("variance %f", ih.pm.GetSampleVariance()))
	fmt.Println(fmt.Sprintf("sample size %d", ih.pm.GetSampleSize()))
}
func TestAddObservationNormalDistConvergenceMiddleStrict(t *testing.T) {
	rand.Seed(1234)
	ih := Init(.01, -1.0, 1.0)
	n := statistics.NormalDistribution{Mean: 0, StandardDeviation: 1}
	var convergence bool = false
	var iterations int64 = 1000
	for convergence != true {
		var i int64 = 0
		for i < iterations {
			ih.AddObservation(n.InvCDF(rand.Float64()))
			i++
		}
		convergence, iterations = ih.TestForConvergence(.025, .975, .95, .001) //upper confidence limit test, lower confidence limit test, confidenece, error tolerance
		fmt.Println(fmt.Sprintf("Computed %d estimated to need %d more iterations", ih.pm.GetSampleSize(), iterations))
	}
	fmt.Println("****Converged 97.5 2.5******")
	fmt.Println(fmt.Sprintf("numbins %d", len(ih.GetBins())))
	fmt.Println(fmt.Sprintf("min %f", ih.pm.GetMin()))
	fmt.Println(fmt.Sprintf("max %f", ih.pm.GetMax()))
	fmt.Println(fmt.Sprintf("mean %f", ih.pm.GetMean()))
	fmt.Println(fmt.Sprintf("variance %f", ih.pm.GetSampleVariance()))
	fmt.Println(fmt.Sprintf("sample size %d", ih.pm.GetSampleSize()))
}
func TestAddObservationNormalDistConvergenceMostStrict(t *testing.T) {
	rand.Seed(1234)
	ih := Init(.01, -1.0, 1.0)
	n := statistics.NormalDistribution{Mean: 0, StandardDeviation: 1}
	var convergence bool = false
	var iterations int64 = 1000
	for convergence != true {
		var i int64 = 0
		for i < iterations {
			ih.AddObservation(n.InvCDF(rand.Float64()))
			i++
		}
		convergence, iterations = ih.TestForConvergence(.01, .99, .95, .001) //upper confidence limit test, lower confidence limit test, confidenece, error tolerance
		fmt.Println(fmt.Sprintf("Computed %d estimated to need %d more iterations", ih.pm.GetSampleSize(), iterations))
	}
	fmt.Println("****Converged 99 1******")
	fmt.Println(fmt.Sprintf("numbins %d", len(ih.GetBins())))
	fmt.Println(fmt.Sprintf("min %f", ih.pm.GetMin()))
	fmt.Println(fmt.Sprintf("max %f", ih.pm.GetMax()))
	fmt.Println(fmt.Sprintf("mean %f", ih.pm.GetMean()))
	fmt.Println(fmt.Sprintf("variance %f", ih.pm.GetSampleVariance()))
	fmt.Println(fmt.Sprintf("sample size %d", ih.pm.GetSampleSize()))
}

func TestBins(t *testing.T) {
	ih := Init(1.0, 0.0, 5.0)
	vals := []float64{-.5, .5, 1.5, 2.5, 3.5, 4.5}
	expected := []int64{-1, 0, 1, 2, 3, 4}
	for idx := range vals {
		ih.AddObservation(vals[idx])
	}
	for idx, val := range expected {
		if ih.Bins()[idx] != float64(val) {
			t.Errorf("BinStart(%d) = %f; expected %f", idx, ih.Bins()[idx], float64(val))
		}
	}
}

func TestBinsAgain(t *testing.T) {
	ih := Init(2.0, 1.0, 5.0)
	vals := []float64{-.5, .5, 1.5, 2.5, 3.5, 4.5}
	expected := []int64{-1, 1, 3}
	for idx := range vals {
		ih.AddObservation(vals[idx])
	}
	for idx, val := range expected {
		if ih.Bins()[idx] != float64(val) {
			t.Errorf("BinStart(%d) = %f; expected %f", idx, ih.Bins()[idx], float64(val))
		}
	}
}

func TestBinCounts(t *testing.T) {
	ih := Init(1.0, 0.0, 5.0)
	vals := []float64{-.5, .5, 1.5, 2.5, 3.5, 4.5}
	expected := []int64{1, 1, 1, 1, 1, 1}
	for idx := range vals {
		ih.AddObservation(vals[idx])
	}
	for idx, val := range expected {
		if ih.BinCounts()[idx] != val {
			t.Errorf("Bin(%d) = %v; expected %v", idx, ih.BinCounts()[idx], val)
		}
	}
}
