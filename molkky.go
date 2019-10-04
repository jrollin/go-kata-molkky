package molkky

// Molkky game
type Molkky struct {
	score int
	rolls []int
}

func startMolkky() Molkky {
	return Molkky{score: 0}
}

func (m *Molkky) getScore() int {
	return m.score
}

func (m *Molkky) setFallenkeels(kells []int) {
	m.rolls = kells
	if len(kells) == 0 {
		m.score += 0
	} else {

		if m.isValidKeels(kells) {
			if len(kells) > 1 {
				m.score += len(kells)
			} else {
				m.score += kells[0]
			}
		}

	}
}

func (m *Molkky) isValidKeels(keels []int) bool {
	for _, keel := range keels {
		if !m.isValidKeel(keel) {
			return false
		}
	}
	return true
}

func (m *Molkky) isValidKeel(val int) bool {
	if val >= 1 && val <= 12 {
		return true
	}
	return false
}
