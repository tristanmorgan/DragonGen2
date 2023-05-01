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
	"github.com/tristanmorgan/dragongen/gen/biome"
)

// Generator Is the generator struct
// that will be used to generate.
type Generator struct {
}

// NewGenerator Returns a pointer to the generator struct.
// This can be changed by setting conf.Generator.
func NewGenerator(seed int64) *Generator {
	biome.RegisterBiomes(seed)
	return &Generator{}
}
