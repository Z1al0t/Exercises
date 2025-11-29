package utility

import (
	"math/rand"
	"slices"
)

// RandomExerciseIndex gets random exercise which wasn't solved yet
func RandomExerciseIndex(m int, done []int) ([]int, int, bool) {
	moreExercises := true
	index := 0
	for i := 0; i < m; i++ {
		randIndex := rand.Intn(m)
		if slices.Contains(done, randIndex) {
			if len(done) == m {
				moreExercises = false
				break
			}
			i--
			continue
		}
		done = append(done, randIndex)
		index = randIndex
		break
	}
	return done, index, moreExercises
}
