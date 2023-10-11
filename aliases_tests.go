package units

// This file contains helpers for testing aliases.

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type aliasTest struct {
	expect *Unit
	alias  string
}

func (a *aliasTest) string() string {
	if a.expect == nil {
		return a.alias + "->nil"
	}
	return a.alias + "->" + a.expect.Name
}

func (a *aliasTest) validate() error {
	if a.expect != nil && !a.expect.HasAlias(a.alias) {
		return fmt.Errorf("expected %q to have alias %q", a.expect.Name, a.alias)
	}
	got, err := Find(a.alias)
	if err != nil {
		if a.expect != nil {
			return fmt.Errorf("expected %q, got error %v", a.expect.Name, err)
		}
	} else if got == nil {
		if a.expect != nil {
			return fmt.Errorf("expected %q, got nil", a.expect.Name)
		}
	} else {
		if a.expect == nil {
			return fmt.Errorf("expected nil, got %q", got.Name)
		} else if got.Name != a.expect.Name {
			return fmt.Errorf("expected %q, got %q", a.expect.Name, got.Name)
		}
	}

	return nil
}

// testAliases tests that aliases are correctly mapped to their units.
func testAliases(t *testing.T, tests []aliasTest) {
	for _, tt := range tests {
		t.Run(
			tt.string(), func(t *testing.T) {
				err := tt.validate()
				assert.Nil(t, err)
			},
		)
	}
}
