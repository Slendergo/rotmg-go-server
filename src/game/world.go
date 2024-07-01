package game

type World struct {
	Id          int32
	Width       int32
	Height      int32
	IdName      string
	DisplayName string

	nextId  int32
	Players map[int32]*Player
	Enemies map[int32]*Enemy

	elapsed float64

	tiles [][]*Tile
}

type Tile struct {
	Type int32
}

type XMLWorld struct {
	IdName      string
	Width       int32
	Height      int32
	DisplayName string
}

func NewWorld(id int32, xmlWorld *XMLWorld) *World {
	world := &World{
		Id:          id,
		IdName:      xmlWorld.DisplayName,
		Width:       xmlWorld.Width,
		Height:      xmlWorld.Height,
		DisplayName: xmlWorld.DisplayName,

		nextId:  0,
		Players: make(map[int32]*Player),
		Enemies: make(map[int32]*Enemy),

		elapsed: 0.0,
	}

	world.tiles = make([][]*Tile, xmlWorld.Width)
	for x := range world.tiles {
		world.tiles[x] = make([]*Tile, xmlWorld.Height)
		for y := range world.tiles[x] {
			world.tiles[x][y] = &Tile{
				Type: 0xFF,
			}
		}
	}

	return world
}

func (world *World) ParseMap() bool {
	for x := range world.tiles {
		for y := range world.tiles[x] {
			tile := world.tiles[x][y]
			tile.Type = 0x36
		}
	}
	return true
}

func (world *World) CreatePlayer(connection *Connection, x float32, y float32) *Player {
	nextId := world.nextId
	world.nextId++

	player := NewPlayer(connection, nextId, x, y)
	world.Players[nextId] = player
	return player
}

func (world *World) Tick(dt float64) bool {
	world.elapsed += dt
	world.updateObjects(dt)
	world.sendNewTick(dt)
	world.clearObjectFlags()
	return true
}

func (world *World) updateObjects(dt float64) {
	world.updateEnemies(dt)
	world.updatePlayers(dt)
}

func (world *World) updateEnemies(dt float64) {
	for id, enemy := range world.Enemies {
		if !enemy.Update(dt) {
			delete(world.Enemies, id)
		}
	}
}

func (world *World) updatePlayers(dt float64) {
	for id, player := range world.Players {
		if !player.Update(dt) {
			delete(world.Players, id)
		}
	}
}

func (world *World) sendNewTick(dt float64) {
	for _, player := range world.Players {
		player.Connection.NewTick(dt)
	}
}

func (world *World) clearObjectFlags() {
	for _, player := range world.Players {
		player.ClearFlags()
	}

	for _, enemy := range world.Enemies {
		enemy.ClearFlags()
	}
}
