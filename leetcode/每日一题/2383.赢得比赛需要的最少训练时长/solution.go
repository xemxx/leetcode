// @algorithm @lc id=2459 lang=golang
// @title minimum-hours-of-training-to-win-a-competition

package main

// @test(5,3,[1,4,3,2],[2,6,3,1])=8
// @test(2,4,[1],[3])=0
// @test(1, 1, [1,1,1,1], [1,1,1,50])=51
func minNumberOfHours(initialEnergy int, initialExperience int, energy []int, experience []int) int {
	cEg, cEp := initialEnergy, initialExperience
	addEg, addEp := 0, 0
	for i := 0; i < len(experience); i++ {
		cEg -= energy[i]
		if cEp <= experience[i] {
			addEp += (experience[i] - cEp + 1)
			cEp = (experience[i] + 1)
		}
		cEp += experience[i]
	}
	if cEg < 1 {
		addEg = 0 - cEg + 1
	}

	return addEg + addEp
}
