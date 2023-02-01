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
	"github.com/df-mc/dragonfly/server/world"
	"github.com/df-mc/dragonfly/server/world/chunk"
)

// GenerateChunk Finds the correct biome for the chunk, and generates
// the chunk accordingly.
func (g Generator) GenerateChunk(pos world.ChunkPos, chunk *chunk.Chunk) {
	biome.Generate(chunk, pos)
}
