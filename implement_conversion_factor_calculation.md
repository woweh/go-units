# Plan: Implement Calculated Conversion Factors for All Derived Units

This plan systematically refactors all unit files to derive conversion factors from base units where possible, reviews hardcoded factors for accuracy, and verifies tests.

## Steps
Implement conversion factor calculations for derived units.

- Make sure the base SI unit is mentioned/documented per quantity.

- Add helper functions for the factor calculation.
  - Follow the "*Factor naming pattern", like in 'area_units.go: `func areaFactor(length Unit) float64`
  - Document dimensional analysis:
    - Add comments showing the base unit breakdown for each derived unit (e.g., // N = kg·m/s²) to make derivations transparent.
  - Where possible, derive the values. Only use empirical values if you can't derive a value.

- Review hardcoded conversion factors for units that cannot be derived.
  - Verify against authoritative sources (NIST, ISO, https://qudt.org)
  - Document sources and values:
    - Add comments showing the URLs and values of the authoritative sources (NIST, ISO).

- Review and update unit tests
  - Update test expected values if deriving from base units produces more accurate factors.
  - Add test cases for newly calculated conversions
  - Verify precision/rounding matches physical reality.
    - Note that there is a default value (conversions_test.go, _defaultTolerance = 1e-8), which can be overwritten in 'type conversionTest > tol'.
  - Make sure all non-SI units are tested.
  - For metric SI units, it's sufficient to test just some units, e.g. the largest (e.g. 'QuettaMeter') and smallest (e.g. QuectoMeter) and one or two in between.
  - Make sure that all Revit conversion factors are covered (> RevitUnits.csv) - for those Revit units that are supported.
    - NOTE that the Revit internal reference units, are often NOT SI units!!!
    - Revit often uses imperial units are base units, sometimes SI, sometimes a mixture of imperial and SI!!!
    - Document the Revit base unit!!!
    - If there is a significant difference between Revit units and 'our' units, there are at least two possibilities:
      a) our calculations, formulae or conversion factors are wrong.
      b) Revit is wrong.
      Document those cases, mark them as TODO, so they can be investigated.

- While your checking all units per quantity, check the naming too:
  - A UnitQuantity should just have the name of the quantity, like `var Area = NewUnitQuantity("area")`, NOT `var AreaQuantity = NewUnitQuantity("area")`
  - Again, the helper factor calculation functions should follow the naming convention "{quantity}Factor", e.g. `func areaFactor`.

A good example/template is the 'area' unit (area_units.go).

Apply to the Uber Go Style Guide
https://github.com/uber-go/guide/blob/master/style.md

Since this is a library, only a portion of the style guide is applicable.
