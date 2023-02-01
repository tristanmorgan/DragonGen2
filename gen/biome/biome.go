package biome

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
	"github.com/df-mc/dragonfly/server/world"
	"github.com/df-mc/dragonfly/server/world/chunk"
)

type Biome interface {
	// IsPosInBiome Returns whether the given chunk
	// pos is in this biome
	IsPosInBiome(pos world.ChunkPos) bool
	// GenChunk Is the function that generates the chunk
	// for this specific biome.
	GenChunk(chunk *chunk.Chunk, pos world.ChunkPos)
}

// biomes Is a list of biomes that can be used
var biomes []Biome

// RegisterBiomes Is used to register all the
// biomes.
func RegisterBiomes() {

}

// Generate Will find the correct biome, and use that to generate
// the chunk.
func Generate(chunk *chunk.Chunk, pos world.ChunkPos) {
	for _, biome := range biomes {
		if biome.IsPosInBiome(pos) {
			biome.GenChunk(chunk, pos)
			return
		}
	}
	biomes[0].GenChunk(chunk, pos)
}
