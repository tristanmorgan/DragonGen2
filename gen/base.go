package gen

/**
    ____                               ______          ___
   / __ \_________ _____ _____  ____  / ____/__  ____ |__ \
  / / / / ___/ __ `/ __ `/ __ \/ __ \/ / __/ _ \/ __ \__/ /
 / /_/ / /  / /_/ / /_/ / /_/ / / / / /_/ /  __/ / / / __/
/_____/_/   \__,_/\__, /\____/_/ /_/\____/\___/_/ /_/____/
                 /____/
	(c) Taylor R (https://github.com/TacoError) 2023
*/

import (
	"DragonGen2/gen/biome"
	goperlin "github.com/aquilax/go-perlin"
)

// Generator Is the generator struct
// that will be used to generate.
type Generator struct {
	// noise Is a goperlin.Perlin object
	// used to generate perlin according
	// to a seed.
	noise *goperlin.Perlin
}

// NewGenerator Returns a pointer to the generator struct.
// This can be changed by setting conf.Generator.
func NewGenerator(seed int64) *Generator {
	biome.RegisterBiomes()
	return &Generator{
		noise: goperlin.NewPerlin(2., 2., 3, seed),
	}
}
