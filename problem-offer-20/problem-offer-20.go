package problem_offer_20

import "strings"

type StateType int

const (
	StateTypeLookup0 StateType = iota
	StateTypeLookup1
	StateTypeLookup2
	StateTypeDecimal0
	StateTypeDecimal1
	StateTypeDecimal2
	StateTypeE
	StateTypeLookupE1
	StateTypeLookupE2
)

func isNumber(s string) bool {
	s = strings.TrimSpace(s)
	pos := 0
	state := StateTypeLookup0
	for pos < len(s) {
		c := s[pos]
		switch state {
		case StateTypeLookup0:
			if c == '+' || c == '-' || c == '.' || c >= '0' && c <= '9' {
				state = StateTypeLookup1
				if c == '+' || c == '-' {
					pos++
				}
			} else {
				return false
			}
		case StateTypeLookup1:
			if c >= '0' && c <= '9' {
				state = StateTypeLookup2
			} else if c == '.' {
				state = StateTypeDecimal0
			} else {
				return false
			}
			pos++
		case StateTypeLookup2:
			if c >= '0' && c <= '9' {
				state = StateTypeLookup2
			} else if c == '.' {
				state = StateTypeDecimal1
			} else if c == 'e' || c == 'E' {
				state = StateTypeE
			} else {
				return false
			}
			pos++
		case StateTypeDecimal0:
			if c >= '0' && c <= '9' {
				state = StateTypeDecimal2
			} else {
				return false
			}
			pos++
		case StateTypeDecimal1:
			if c >= '0' && c <= '9' {
				state = StateTypeDecimal2
			} else if c == 'e' || c == 'E' {
				state = StateTypeE
			} else {
				return false
			}
			pos++
		case StateTypeDecimal2:
			if c >= '0' && c <= '9' {
				state = StateTypeDecimal2
			} else if c == 'e' || c == 'E' {
				state = StateTypeE
			} else {
				return false
			}
			pos++
		case StateTypeE:
			if c == '+' || c == '-' || c >= '0' && c <= '9' {
				if c == '+' || c == '-' {
					pos++
				}
				state = StateTypeLookupE1
			} else {
				return false
			}
		case StateTypeLookupE1, StateTypeLookupE2:
			if c >= '0' && c <= '9' {
				state = StateTypeLookupE2
			} else {
				return false
			}
			pos++
		}
	}

	return state == StateTypeLookup2 || state == StateTypeLookupE2 || state == StateTypeDecimal1 || state == StateTypeDecimal2
}
