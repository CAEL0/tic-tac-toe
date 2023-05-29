package board

type Board struct {
	state [9]int
	turn  int
}

func New() *Board {
	return &Board{
		state: [9]int{0, 0, 0, 0, 0, 0, 0, 0, 0},
		turn:  1,
	}
}

func (b *Board) State() [9]int {
	return b.state
}

func (b *Board) clear() {
	for i := range b.state {
		b.state[i] = 0
	}
	b.turn = 1
}
