package api

import(
	"math/rand"
)

type CardFace int

const (
	One CardFace = iota
	Two
	Three
	Four
	Five
	Six
	Seven
)

var cardFaceNames = map[CardFace]string{
	One:   "One",
	Two:   "Two",
	Three: "Three",
	Four:  "Four",
	Five:  "Five",
	Six:   "Six",
	Seven: "Seven",
}

func (cf CardFace) String() string {
	return cardFaceNames[cf]
}

var NumberFaces = []CardFace{One, Two, Three, Four, Five, Six, Seven}
//var SpecialFaces = []CardFace{Skip, Reverse, DrawTwo, Wild, WildDrawFour}

type Color int 

const (
	Red Color = iota
	Yellow
	Blue
	Black
	White
)

var colorNames = map[Color]string{
	Red:    "Red",
	Yellow: "Yellow",
	Blue:   "Blue",
	Black:  "Black",
	White:  "White",
}

func (c Color) String() string {
	return colorNames[c]
}

var CardColors = []Color{Red, Yellow, Blue, Black, White}

type Card struct {
	Face  CardFace
	Color Color
}

func (c Card) String() string {
	return c.Color.String() + " " + c.Face.String()
}

func Compare(c1, c2 Card) bool {
	return c1.Face == c2.Face || c1.Color == c2.Color
}

func NewDeck() []Card {
	deck := make([]Card, 0)
	for color := 0; color < len(CardColors); color++ {
		for face := 0; face < len(NumberFaces); face++ {
			deck = append(deck, 
				Card{Face: NumberFaces[face], Color: CardColors[color]}, 
				Card{Face: NumberFaces[face], Color: CardColors[color]}, 
				Card{Face: NumberFaces[face], Color: CardColors[color]})
		}
	}

	return Shuffle(deck)
}

func Shuffle(deck []Card) []Card {
	for i := range deck {
		j := rand.Intn(i + 1)
		deck[i], deck[j] = deck[j], deck[i]
	}
	return deck
}

func Draw(deck []Card, n int) ([]Card, []Card) {
	return deck[n:], deck[:n]
}
