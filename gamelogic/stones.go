package gamelogic

import "fmt"

// A group of stones, defined by player color and two lists of Points, stones and liberties.
type StoneGroup struct {
	Color Player
	Stones []Point
	Liberties []Point
}

// Method implements Stringer interface for StoneGroup struct.
func (sg StoneGroup) String() string {
	s1, s2, s3 := "StoneGroup{ ", " -Stones: ", "\n -Liberties: "
	c := "Black\n"
	if sg.Color != Black {
		c = "White\n"
	}
	return fmt.Sprint(s1, c, s2, sg.Stones, s3, sg.Liberties, " }")
}

// Method returns the number of liberties in a StoneGroup.
func (sg StoneGroup) NumLiberties() int {
	return len(sg.Liberties)
}

// Method returns the number of stones in a StoneGroup.
func (sg StoneGroup) NumStones() int {
	return len(sg.Stones)
}

// Method for adding a liberty to a StoneGroup.
func (sg *StoneGroup) AddLiberty(p Point) {
	if _, b := contains(sg.Liberties, p); !b {
		sg.Liberties = append(sg.Liberties, p)
	}
}

// Method for removing a liberty to a StoneGroup.
func (sg *StoneGroup) RemoveLiberty(p Point) {
	if i, b := contains(sg.Liberties, p); b {
		s := sg.Liberties
		s[len(s)-1], s[i] = s[i], s[len(s)-1]
    	sg.Liberties = s[:len(s)-1]
	}
}

// Method for merging a StoneGroup with another passed StoneGroup.
func (sg *StoneGroup) MergeWith(mg StoneGroup) {
	if sg.Color != mg.Color {
		return
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