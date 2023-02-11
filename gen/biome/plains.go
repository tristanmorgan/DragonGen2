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
	"github.com/TacoError/DragonGen2/gen/perlin"
	"github.com/df-mc/dragonfly/server/world"
	"github.com/df-mc/dragonfly/server/world/chunk"
)

type plains struct {
}

func (plains) IsPosInBiome(world.ChunkPos) bool {
	return true
}

func (plains) GenChunk(chunk *chunk.Chunk, pos world.ChunkPos) {
	worldX := pos.X() << 4
	worldZ := pos.Z() << 4
	var x uint8
	var z uint8
	var y int16
	for x = 0; x <= 15; x++ {
		for z = 0; z <= 15; z++ {
			y = perlin.OctavePerlin(
				Noise,
				3,
				float64(x)+float64(worldX),
				float64(z)+float64(worldZ),
				40,
			)
			chunk.SetBlock(x, y, z, 0, grass)
			for yy := y - 1; yy >= 1; yy-- {
				chunk.SetBlock(x, yy, z, 0, stone)
			}
		}
	}
}
