// RNG project RNG.go
package slLib

type RNG struct {
	// LCG using GCC's constants
	m     uint32
	a     uint32
	c     uint32
	state uint32
}

func (rng *RNG) Init(seed uint32) {
	rng.m = 0x100000000 - 1 // 2**32;
	rng.a = 1103515245
	rng.c = 12345
	if seed == 0 {
		rng.state = 234236743 * (rng.m - 1)
	}
	if seed != 0 {
		rng.state = seed
	}
}

func (rng *RNG)  NextInt() uint32 {
	rng.state = (rng.a*rng.state + rng.c) % rng.m
	return rng.state
}

func (rng *RNG) NextFloat() float64 {
	// returns in range [0,1]
	return float64(rng.nextInt() / (rng.m - 1))
}

func (rng *RNG) NextRange(start uint32, end uint32) uint32 {
	// returns in range [start, end): including start, excluding end
	// can't modulu nextInt because of weak randomness in lower bits
	var rangeSize uint32
	rangeSize = end - start
	var randomUnder1 float32
	randomUnder1 = float32(rng.nextInt()) / float32(rng.m)
	return start + uint32(randomUnder1*float32(rangeSize))
}

func (rng *RNG) Dice(number int, dice int, plus int) int {
	var s int
	s = 0
	for i := 0; i < number; i++ {
		s = s + int(rng.nextRange(1, uint32(dice+1)))
	}
	return s + plus
}
