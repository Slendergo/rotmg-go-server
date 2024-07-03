package game

import (
	assets "main/assets"
	xml "main/assets/xml"
	"main/network"
	"main/structures"
	"math"
)

type Entity struct {
	Id        int32
	MaxHealth int32
	Health    int32
	Size      float32
	Type      int32
	X         float32
	Y         float32
	LastX     float32
	LastY     float32
	flags     int32

	// if its a entity it will have this
	XMLEnemy *xml.XMLEnemy

	ObjectProps *xml.XMLObjectProperties
	Dead        bool
}

func NewEntity(objectProps *xml.XMLObjectProperties, id int32, x float32, y float32) *Entity {

	if objectProps.IsPlayer {
		return nil
	}

	entity :=
		&Entity{
			Id:          id,
			Size:        float32(objectProps.Size),
			Type:        objectProps.Type,
			X:           x,
			Y:           y,
			LastX:       x,
			LastY:       y,
			flags:       structures.NoFlags,
			ObjectProps: objectProps,
			Dead:        false,
		}

	if objectProps.IsEnemy {
		entity.XMLEnemy = assets.GlobalXMLLibrary.GetXMLEnemy(objectProps.Type)

		entity.MaxHealth = entity.XMLEnemy.MaxHitPoints
		entity.Health = entity.XMLEnemy.MaxHitPoints
		entity.Size = float32(entity.XMLEnemy.Size)
	}

	return entity
}

func (entity *Entity) NewObjectData() network.NewObjectData {
	newObjectData := network.NewObjectData{
		ObjectType: entity.Type,
		StatusData: network.StatusData{
			ObjectId: entity.Id,
			X:        entity.X,
			Y:        entity.Y,
			Stats:    []network.StatData{},
		},
	}
	return newObjectData
}

func (entity *Entity) HasFlag(flag int32) bool {
	return (entity.flags & flag) != 0
}

func (entity *Entity) Facing() float32 {
	dx := entity.X - entity.LastX
	dy := entity.Y - entity.LastY
	return float32(math.Atan2(float64(dy), float64(dx)))
}

func (entity *Entity) SetPosition(x float32, y float32) {
	entity.LastX = entity.X
	entity.LastY = entity.Y
	entity.X = x
	entity.Y = y
	entity.flags |= structures.MovedFlag
}

func (entity *Entity) Update(dt float64) bool {
	if entity.Dead {
		return false
	}

	// logic

	return !entity.Dead
}

func (entity *Entity) TakeDamage(damage int32) {
}

func (entity *Entity) Kill() {
	entity.Dead = true
}

func (entity *Entity) ClearFlags() {
	entity.flags &= structures.NoFlags
}
