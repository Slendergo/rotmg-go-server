package game

import (
	"main/assets"
)

type World struct {
	Id          int32
	Width       int32
	Height      int32
	IdName      string
	DisplayName string

	nextEntityId int32
	Players      map[int32]*Player
	Entities     map[int32]*Entity

	elapsed float64

	tiles [][]*Tile
}

type Tile struct {
	Type         int32
	StaticEntity Entity
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

		nextEntityId: 0,
		Players:      make(map[int32]*Player),
		Entities:     make(map[int32]*Entity),

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

func (world *World) InBoundsFloat(x float32, y float32) bool {
	return x >= 0 && y >= 0 && x < float32(world.Width) && y < float32(world.Height)
}

func (world *World) InBoundsInt(x int, y int) bool {
	return x >= 0 && y >= 0 && x < int(world.Width) && y < int(world.Height)
}

func (world *World) InBoundsInt32(x int32, y int32) bool {
	return x >= 0 && y >= 0 && x < world.Width && y < world.Height
}

func (world *World) ParseMap() bool {

	mapData := assets.GlobalMapLibrary.GetMapData("nexus.jm")
	if mapData == nil {
		return false
	}

	worldCenterX := world.Width / 2
	worldCenterY := world.Height / 2

	centerX := float32(mapData.Width) / 2.0
	centerY := float32(mapData.Height) / 2.0

	for x := range mapData.Width {
		for y := range mapData.Height {

			cX := int(float32(x) - centerX + float32(worldCenterX))
			cY := int(float32(y) - centerY + float32(worldCenterY))

			if !world.InBoundsInt(cX, cY) {
				continue
			}

			index := x + y*mapData.Width

			tileData := mapData.Tiles[index]

			tile := world.tiles[cX][cY]

			if tileData.Ground != 0xFF {
				tile.Type = tileData.Ground
			}

			if tileData.Object != -1 {
				world.CreateEntity(tileData.Object, float32(cX)+0.5, float32(cY)+0.5)
			}

			if tileData.Region != -1 {
				// todo
			}
		}
	}
	return true
}

func (world *World) CreatePlayer(connection *Connection, typ int32, x float32, y float32) *Player {
	nextId := world.nextEntityId
	world.nextEntityId++

	props := assets.GlobalXMLLibrary.GetXMLObjectProperties(typ)
	if props == nil {
		return nil
	}

	player := NewPlayer(connection, props, nextId, x, y)
	world.Players[nextId] = player
	return player
}

func (world *World) CreateEntity(typ int32, x float32, y float32) *Entity {
	nextId := world.nextEntityId
	world.nextEntityId++

	props := assets.GlobalXMLLibrary.GetXMLObjectProperties(typ)
	if props == nil {
		return nil
	}

	entity := NewEntity(props, nextId, x, y)
	world.Entities[nextId] = entity
	return entity
}

func (world *World) Tick(dt float64) bool {
	world.elapsed += dt
	world.updateObjects(dt)
	world.sendNewTick(dt)
	world.clearObjectFlags()
	return true
}

func (world *World) updateObjects(dt float64) {
	world.updateEntities(dt)
	world.updatePlayers(dt)
}

func (world *World) updateEntities(dt float64) {
	for id, entity := range world.Entities {
		if !entity.Update(dt) {
			delete(world.Entities, id)
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
		if !player.Dead && player.Connection.Connected {
			player.Connection.NewTick(dt)
		}
	}
}

func (world *World) clearObjectFlags() {
	for _, player := range world.Players {
		player.ClearFlags()
	}

	for _, entity := range world.Entities {
		entity.ClearFlags()
	}
}
