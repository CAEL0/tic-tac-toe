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

func (b *Board) UpdateState(index int, player int) bool {
	if b.turn != player || index < 0 || len(b.state) <= index || b.state[index] != 0 {
		return false
	}
	b.state[index] = player
	b.turn = 3 - player
	return true
}

func (b *Board) clear() {
	for i := range b.state {
		b.state[i] = 0
	}
	b.turn = 1
}
