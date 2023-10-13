package units

import (
	"testing"
)

// quantityForUnitTests is a special quantity used for unit tests.
const quantityForUnitTests = "<|_quantity_for_unit_tests_|>"

func Test_NewUnit(t *testing.T) {
	tests := []struct {
		name       string
		unitName   string
		unitSymbol string
		wantErr    bool
	}{
		{
			name:       "Add TestUnit_1",
			unitName:   "TestUnit_1",
			unitSymbol: "T1",
			wantErr:    false,
		},
		{
			name:       "Test empty name",
			unitName:   "",
			unitSymbol: "T1",
			wantErr:    true,
		},
		{
			name:       "Test duplicate name (TestUnit_1)",
			unitName:   "TestUnit_1",
			unitSymbol: "",
			wantErr:    true,
		},
		{
			name:       "Testing for duplicate symbol (T1)",
			unitName:   "TestUnit_1.0",
			unitSymbol: "T1",
			wantErr:    true,
		},
		{
			name:       "Testing for case sensitive symbol (t1 != T1)",
			unitName:   "TestUnit_1.1",
			unitSymbol: "t1",
			wantErr:    false,
		},
	}

	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				u, err := NewUnit(tt.unitName, tt.unitSymbol, UnitOptionQuantity(quantityForUnitTests))
				if (err != nil) != tt.wantErr {
					t.Errorf("NewUnit() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
				if u != nil {
					if u.Name != tt.unitName {
						t.Errorf("NewUnit() = %v, want %v", u.Name, tt.unitName)
					}
					if u.Symbol != tt.unitSymbol {
						t.Errorf("NewUnit() = %v, want %v", u.Symbol, tt.unitSymbol)
					}
				}
			},
		)
	}
}
