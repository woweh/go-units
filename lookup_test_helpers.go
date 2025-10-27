package units

// This file contains helpers for testing the find/lookup of names, aliases and symbols.

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// lookUpTestUnit is a test unit for looking up names, aliases or symbols.
type lookUpTestUnit struct {
	expect Unit   // the expected unit, or nil if an error is expected
	key    string // the name, alias or symbol to look up
}

type lookUpTests []lookUpTestUnit

func (u *lookUpTestUnit) string() string {
	if u.expect == nil {
		return u.key + "->nil"
	}
	return u.key + "->" + u.expect.Name
}

func (u *lookUpTestUnit) verifyFind() error {
	got, err := Find(u.key)

	if err != nil {
		if u.expect != nil {
			return fmt.Errorf("expected %q, got error %v", u.expect.Name, err)
		}
		return nil
	}

	if got == nil {
		if u.expect != nil {
			return fmt.Errorf("expected %q, got nil", u.expect.Name)
		}
		return nil
	}

	if u.expect == nil {
		return fmt.Errorf("expected nil, got %q", got.Name)
	}

	if got.Name != u.expect.Name {
		return fmt.Errorf("expected %q, got %q", u.expect.Name, got.Name)
	}

	return nil
}

func (u *lookUpTestUnit) verifyHasNameOrSymbol() error {
	if u.expect == nil {
		return nil
	}
	if !u.expect.HasName(u.key) && !u.expect.HasSymbol(u.key) {
		return fmt.Errorf("expected %q to have name or symbol %q", u.expect.Name, u.key)
	}
	return nil
}

// testLookupNamesAndSymbols tests that names, aliases or symbols are correctly mapped to their units.
func testLookupNamesAndSymbols(t *testing.T, tests lookUpTests) {
	for _, tt := range tests {
		t.Run(
			tt.string(), func(t *testing.T) {
				err := tt.verifyFind()
				assert.Nil(t, err)
				err = tt.verifyHasNameOrSymbol()
				assert.Nil(t, err)
			},
		)
	}
}
