package deck

import (
	"fmt"
	"math/rand"
	"testing"
)

func ExampleCard() {
	fmt.Println(Card{Rank: Ace, Suit: Heart})
	fmt.Println(Card{Rank: Queen, Suit: Club})
	fmt.Println(Card{Rank: Seven, Suit: Diamond})
	fmt.Println(Card{Rank: Ace, Suit: Spade})
	fmt.Println(Card{Suit: Joker})

	// Output:
	// Ace of Hearts
	// Queen of Clubs
	// Seven of Diamonds
	// Ace of Spades 
	// Joker
}

func TestNew(t *testing.T) {
	cards := New()
	// 13 rank * 4 suits
	if len(cards) != 13*4 {
		t.Error("Wrong number of cards in new deck.")
	}
}

func TestDefaultSort(t *testing.T) {
	cards := New(DefaultSort)
	exp := Card{Rank: Ace, Suit: Spade} 
	if cards[0] != exp {
		t.Error("Expected Ace of spades as first card. Recieved:", cards[0])
	}
}

func TestSort(t *testing.T) {
	cards := New(Sort(Less))
	exp := Card{Rank: Ace, Suit: Spade} 
	if cards[0] != exp {
		t.Error("Expected Ace of spades as first card. Recieved:", cards[0])
	}
}

func TestShuffle(t *testing.T) {
	
	shuffleRand = rand.New(rand.NewSource(0))
	og := New()
	first := og[40]
	second:= og[35]
	cards := New(Shuffle)
	if cards[0] != first {
		t.Errorf("Expected the first card to be %s, recieved %s.", first, cards[0])
	}
	if cards[1] != second {
		t.Errorf("Expected the second card to be %s, recieved %s.", second, cards[1])
	}

}

func TestJokers(t *testing.T) {
	cards := New(Jokers(3))
	count := 0
	for _, c := range cards {
		if c.Suit == Joker{
			count++
		}
	}

	if count != 3 {
		t.Error("Expected 3 Jokers, recieved:", count)
	}
}

func TestFilter(t *testing.T) {
	filter := func(card Card) bool {
		return card.Rank == 2 || card.Rank == 3
	}

	cards := New(Filter(filter))

	for _, c := range cards {
		if c.Rank == Two || c.Rank == Three {
			t.Error("Expected all twos and threes to be filtered out.")
		}
	}
}

func TestDeck(t *testing.T) {
	cards := New(Deck(3))

	if len(cards) != 13 * 4 * 3 {
		t.Errorf("expected %d cards, recieved %d cards", 13*4*3, len(cards))
	}

}