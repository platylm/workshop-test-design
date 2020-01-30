package testing

import "testing"

func Test_Input_NumberA_3_And_NumberB_2_And_NumberC_1_Should_Be_Maximum_3(t *testing.T) {
	expectedMaximum := 3

	numberA := 3
	numberB := 2
	numberC := 1

	actualMaximum := maxAndMean(numberA, numberB, numberC)
	if expectedMaximum != actualMaximum {
		t.Errorf("Expected is %d but actual is %d", expectedMaximum, actualMaximum)
	}
}

func Test_Input_NumberA_5_And_NumberB_10_And_NumberC_7_Should_Be_Maximum_10(t *testing.T) {
	expectedMaximum := 10

	numberA := 5
	numberB := 10
	numberC := 7

	actualMaximum := maxAndMean(numberA, numberB, numberC)
	if expectedMaximum != actualMaximum {
		t.Errorf("Expected is %d but actual is %d", expectedMaximum, actualMaximum)
	}
}

func Test_Input_NumberA_1_And_NumberB_5_And_NumberC_7_Should_Be_Maximum_7(t *testing.T) {
	expectedMaximum := 7

	numberA := 1
	numberB := 5
	numberC := 7

	actualMaximum := maxAndMean(numberA, numberB, numberC)
	if expectedMaximum != actualMaximum {
		t.Errorf("Expected is %d but actual is %d", expectedMaximum, actualMaximum)
	}
}
