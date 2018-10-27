package mash

import "testing"

/** Test the loss in liters from boil off using the default BOILOFF rate*/
func TestMash_ExpectedBoilLoss(t *testing.T) {

	m := Mash{}
	boilOff:= m.ExpectedBoilLoss();

	if boilOff != DEFAULT_BOILOFF_PER_HOUR {
		t.Errorf("Expected %f but got %f", DEFAULT_BOILOFF_PER_HOUR, boilOff)
	}
}

/*Test default boiloff rate for 90 minute mash*/
func TestMash_ExpectedBoilLoss90Min(t *testing.T) {
	expected := DEFAULT_BOILOFF_PER_HOUR*1.5

	m := Mash{boilTime: 90}
	boilOff := m.ExpectedBoilLoss()
	if boilOff != expected {
		t.Errorf("Expected %f but got %f", expected, boilOff)
	}
}
