package algorithm

import "math"

func Foo1() string {
    return "Hello from package1!"
}

type City struct {
	Name string
	X    float64
	Y    float64
}

type TSP struct {
	Cities []City
	N      int // Number of cities
}

func NewTSP(cities []City) *TSP {
	return &TSP{
		Cities: cities,
		N:      len(cities),
	}
}

func (t *TSP) distance(city1, city2 City) float64 {
	dx := city1.X - city2.X
	dy := city1.Y - city2.Y
	return math.Sqrt(dx*dx + dy*dy)
}


func (t *TSP) HopfieldEnergy(states [][]int, A, B, C, D float64) float64 {
	energyObj := 0.0
	energyConst := 0.0

	// Calculate the objective function part of the energy (E_obj)
	for X := 0; X < t.N; X++ {
		for i := 0; i < t.N; i++ {
			for Y := 0; Y < t.N; Y++ {
				if Y != X {
					energyObj += 0.5 * t.distance(t.Cities[X], t.Cities[Y]) * float64(states[X][i]) * (float64(states[Y][(i+1)%t.N]) + float64(states[Y][(i-1+t.N)%t.N]))
				}
			}
		}
	}
	// Calculate the constraint part of the energy (E_const)
	// Constraint 1: Each city is visited only once (No more than two neurons in each row outputting 1)
	for X := 0; X < t.N; X++ {
		for i := 0; i < t.N; i++ {
			for j := 0; j < t.N; j++ {
				if j != i {
					energyConst += A * 0.5 * float64(states[X][i]*states[X][j])
				}
			}
		}
	}

	// Constraint 2: Cannot visit two cities at the same time (No more than two neurons in each column outputting 1)
	for i := 0; i < t.N; i++ {
		for X := 0; X < t.N; X++ {
			for Y := 0; Y < t.N; Y++ {
				if Y != X {
					energyConst += B * 0.5 * float64(states[X][i]*states[Y][i])
				}
			}
		}
	}

	// Constraint 3: Visit all cities (Exactly N neurons outputting 1)
	countCities := 0
	for X := 0; X < t.N; X++ {
		for i := 0; i < t.N; i++ {
			countCities += states[X][i]
		}
	}
	energyConst += C * 0.5 * math.Pow(float64(countCities-t.N), 2)

	// Calculate the total energy (E_net)
	energyNet := 0.5 * energyObj + energyConst

	return energyNet
}
