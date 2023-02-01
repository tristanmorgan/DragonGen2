package perlin

import (
	goperlin "github.com/aquilax/go-perlin"
)

/**
    ____                               ______          ___
   / __ \_________ _____ _____  ____  / ____/__  ____ |__ \
  / / / / ___/ __ `/ __ `/ __ \/ __ \/ / __/ _ \/ __ \__/ /
 / /_/ / /  / /_/ / /_/ / /_/ / / / / /_/ /  __/ / / / __/
/_____/_/   \__,_/\__, /\____/_/ /_/\____/\___/_/ /_/____/
                 /____/
	(c) Taylor R (https://github.com/TacoError) 2023
*/

// OctavePerlin Combines multiple 'octaves' of perlin noise
// in order to create a more natural looking terrain
func OctavePerlin(noise *goperlin.Perlin, octaves int, x float64, z float64, persistence float64) int16 {
	total := 0.0
	frequency := 1
	amplitude := 100.0
	maxValue := 0.0

	for i := 0; i <= octaves; i++ {
		total += noise.Noise2D(x/100, z/100) * amplitude
		maxValue += amplitude
		amplitude *= persistence
		frequency *= 2
	}

	return int16(float64(total/maxValue) * 20)
}
