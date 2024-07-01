package game

import (
	"fmt"
	"time"
)

var GlobalWorldManager *WorldManager

type WorldManager struct {
	worlds      map[int32]*World
	nextWorldId int32
}

func NewWorldManager() *WorldManager {
	return &WorldManager{
		worlds: make(map[int32]*World),
	}
}

func (manager *WorldManager) GetWorld(id int32) *World {
	world, exists := manager.worlds[id]
	if !exists {
		return nil
	}
	return world
}

func (manager *WorldManager) CreateWorld(idName string) *World {
	startTime := time.Now()

	nextId := manager.nextWorldId
	manager.nextWorldId++

	xmlWorld := &XMLWorld{
		IdName:      idName,
		Width:       64,
		Height:      64,
		DisplayName: idName,
	}

	world := NewWorld(nextId, xmlWorld)

	success := world.ParseMap()
	if !success {
		return nil
	}
	manager.worlds[world.Id] = world

	elapsed := time.Since(startTime)
	fmt.Printf("CreateWorld %s %f\n", idName, elapsed.Seconds())

	return world
}

func (manager *WorldManager) RemoveWorld(id int32) {
	delete(manager.worlds, id)
}

func (manager *WorldManager) TickAllWorlds(dt float64) {
	for id, world := range manager.worlds {
		if !world.Tick(dt) {
			delete(manager.worlds, id)
		}
	}
}
