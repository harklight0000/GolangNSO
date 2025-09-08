package sqlplugins

import . "nso/ainterfaces"

type compoundCondition struct {
	Conditions []ICondition
}
