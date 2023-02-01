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
	"github.com/df-mc/dragonfly/server/block"
	"github.com/df-mc/dragonfly/server/world"
)

var (
	grass = world.BlockRuntimeID(block.Grass{})
	stone = world.BlockRuntimeID(block.Stone{})
)
