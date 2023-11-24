package main

import (
	"fmt"
	"math"
	"math/rand"

	"ising/graphics"
)

type Ising struct {
	T   float64
	J   int
	N   int
	Kb  float64
	Net [][]int
}

func (s *Ising) M() int {
	m := 0
	for i := 0; i < s.N; i++ {
		for j := 0; j < s.N; j++ {
			m += s.Net[i][j]
		}
	}
	return m
}

func (s *Ising) dE(spin, H int) int {
	return 2 * spin * H * s.J
}

func (s *Ising) P(dE float64) float64 {
	return math.Exp((-1.0) * dE / (s.Kb * s.T))
}

func (s *Ising) Hi(i, j int) int {
	return s.Net[i][(j-1+s.N)%s.N] + s.Net[i][(j+1+s.N)%s.N] + s.Net[(i-1+s.N)%s.N][j] + s.Net[(i+1+s.N)%s.N][j]
}

func (s *Ising) H() int {
	h := 0
	for i := 0; i < s.N; i++ {
		for j := 0; j < s.N; j++ {
			h += s.Net[i][j] * s.Hi(i, j)
		}
	}
	return (-1) * s.J * h / 2
}

func (s *Ising) Fill() {
	s.Net = make([][]int, 0, s.N)
	for i := 0; i < s.N; i++ {
		s.Net = append(s.Net, make([]int, s.N))
		for j := 0; j < s.N; j++ {
			s.Net[i][j] = rand.Intn(2)*2 - 1 // -1 or 1
		}
	}
}

func (s *Ising) Switch(i, j int) {
	dE := s.dE(s.Net[i][j], s.Hi(i, j))

	if dE <= 0 || s.P(float64(dE)) > rand.Float64() {
		s.Net[i][j] *= -1
	}
}

func (s *Ising) Calculate(n int) (int, int) {
	sumH := 0
	sumM := 0

	for i := 0; i < n; i++ {
		x := rand.Intn(s.N)
		y := rand.Intn(s.N)
		s.Switch(x, y)
		sumH += s.H()
		sumM += s.M()
	}

	return sumH / n, int(math.Abs(float64(sumM / n)))
}

func (s *Ising) Print() {
	for i := 0; i < s.N; i++ {
		fmt.Println(s.Net[i])
	}
}

func main() {
	s := Ising{
		T:  0,
		J:  1,
		Kb: 1.38,
		N:  10,
	}
	s.Fill()
	s.Print()

	//графики
	n := 50
	maxT := 10.0
	hs := make([]float64, 0, n)
	ms := make([]float64, 0, n)
	ts := make([]float64, 0, n)

	for t := 0.0; t < maxT; t += maxT / float64(n) {
		s.T = t
		h, m := s.Calculate(10000)
		hs = append(hs, float64(h))
		ms = append(ms, float64(m))
		ts = append(ts, t)
	}

	err := graphics.PlotGraph(ts, hs, "Динамика средней энергии от температуры")
	if err != nil {
		fmt.Errorf("plot error: %w", err)
	}
	err = graphics.PlotGraph(ts, ms, "Динамика среднего магнитного момента от температуры")
	if err != nil {
		fmt.Errorf("plot error: %w", err)
	}

	//решетки
	s.T = 0.1
	h, m := s.Calculate(10000)
	fmt.Printf("\nT: %.1f, H: %d,\t|M|: %d \n", s.T, h, m)
	s.Print()
	s.T = 2.3
	h, m = s.Calculate(10000)
	fmt.Printf("\nT: %.1f, H: %d,\t|M|: %d \n", s.T, h, m)
	s.Print()
	s.T = 5
	h, m = s.Calculate(10000)
	fmt.Printf("\nT: %.1f, H: %d,\t|M|: %d \n", s.T, h, m)
	s.Print()
	s.T = 25
	h, m = s.Calculate(10000)
	fmt.Printf("\nT: %.1f, H: %d,\t|M|: %d \n", s.T, h, m)
	s.Print()

}
