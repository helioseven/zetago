package gamelogic

import (
	"errors"
	"fmt"
)

// A group of stones, defined by player color and two lists of Points, stones and liberties.
type StoneGroup struct {
	Color Player
	Stones []Point
	Liberties []Point
}

// Method returns a deep copy of a StoneGroup struct.
func (sg *StoneGroup) Copy() *StoneGroup {
	ns := make([]Point, sg.NumStones())
	copy(ns, sg.Stones)
	nl := make([]Point, sg.NumLiberties())
	copy(nl, sg.Liberties)
	return &StoneGroup{sg.Color, ns, nl}
}

// Method returns a boolean comparison of two StoneGroup structs.
func (sg *StoneGroup) Equal(c *StoneGroup) bool {
	if c == nil {
		return false
	}
	b1, b2, b3 := c.Color != sg.Color, c.NumStones() != sg.NumStones(), c.NumLiberties() != sg.NumLiberties()
	if b1 || b2 || b3 {
		return false
	}
	for _, e := range sg.Stones {
		if _, t := contains(c.Stones, e); !t {
			return false
		}
	}
	for _, e := range sg.Liberties {
		if _, t := contains(c.Liberties, e); !t {
			return false
		}
	}
	return true
}

// Method implements Stringer interface for StoneGroup struct.
func (sg *StoneGroup) String() string {
	s1, s2, s3 := "StoneGroup{\n -Color: ", " -Stones: ", "\n -Liberties: "
	c := "Black\n"
	if sg.Color != Black {
		c = "White\n"
	}
	return fmt.Sprint(s1, c, s2, sg.Stones, s3, sg.Liberties, " }")
}

// Method returns the number of liberties in a StoneGroup.
func (sg *StoneGroup) NumLiberties() int {
	return len(sg.Liberties)
}

// Method returns the number of stones in a StoneGroup.
func (sg *StoneGroup) NumStones() int {
	return len(sg.Stones)
}

// Method for adding a liberty to a StoneGroup.
func (sg *StoneGroup) AddLiberty(p Point) error {
	if _, b := contains(sg.Liberties, p); b {
		return errors.New("StoneGroup already contains the given liberty.")
	} else {
		sg.Liberties = append(sg.Liberties, p)
		return nil
	}
}

// Method for removing a liberty to a StoneGroup.
func (sg *StoneGroup) RemoveLiberty(p Point) error {
	if i, b := contains(sg.Liberties, p); !b {
		return errors.New("StoneGroup doesn't have the given liberty.")
	} else {
		s, l := sg.Liberties, sg.NumLiberties()
		s[l-1], s[i] = s[i], s[l-1]
    	sg.Liberties = s[:l-1]
    	return nil
	}
}

// Method for merging in a StoneGroup with another passed StoneGroup.
func (sg *StoneGroup) MergeIn(mg *StoneGroup) error {
	if sg.Color != mg.Color {
		return errors.New("Cannot merge StoneGroup of different player color.")
	}
	cs := sg.Stones
	for _, e := range mg.Stones {
		if _, b := contains(cs, e); !b {
			cs = append(cs, e)
		}
	}
	cl := sg.Liberties
	for _, e := range mg.Liberties {
		if _, b1 := contains(cl, e); !b1 {
			if _, b2 := contains(cs, e); !b2 {
				cl = append(cl, e)
			}
		}
	}
	sg.Stones = cs
	sg.Liberties = cl
	return nil
}

// Utility function checks set membership in a list of points.
func contains(s []Point, e Point) (int, bool) {
	for i, a := range s {
		if a == e {
			return i, true
		}
	}
	return -1, false
}