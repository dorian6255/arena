package char

import ()

type Char struct {
	STR  int
	DEX  int
	CON  int
	INTE int
	WIS  int
	CHA  int
}
type Player struct {
	Char
}

const MAX_ALLO_STATS = 27 //Based on baldur's gate 3 character creator
