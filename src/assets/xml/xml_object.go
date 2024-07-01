package assets

import (
	"main/utils"

	"github.com/beevik/etree"
)

type XMLObjectProperties struct {
	IdName                          string
	Type                            int32
	DisplayId                       string
	Size                            int32
	minSize                         int32
	maxSize                         int32
	sizeStep                        int32
	IsStatic                        bool
	IsItem                          bool
	IsPlayer                        bool
	IsEnemy                         bool
	IsFriendly                      bool
	BlocksSight                     bool
	FullOccupy                      bool
	OccupySquare                    bool
	EnemyOccupySquare               bool
	InterGamePortal                 bool
	IsContainer                     bool
	CanPutNormalObjects             bool
	CanPutSoulboundObjects          bool
	ProtectFromGroundDamage         bool
	ProtectFromGroundEffects        bool
	Connects                        bool
	IsLoot                          bool
	SpawnPoint                      bool
	CanUseTexes                     bool
	LeachHealth                     bool
	BlockNearbyPlayerPortalCreation bool
	PortalBlockRadius               float32
	DungeonName                     string
	GuildItem                       string
	GuildItemParams                 string
	YardType                        int32
	Price                           int32
	Fame                            int32
	SlotTypes                       []int32
	StartingEquipment               []int32
	// todo at some point
	// Projectiles map[int] *XMLProjectile
}

func NewXMLObjectProperties() *XMLObjectProperties {
	return &XMLObjectProperties{}
}

// func (x *XMLObjectProperties) GetProjectile(index int32) *XMLProjectile {
// 	xmlProjectile, ok := x.projectiles[index]
// 	if ok {
// 		return xmlProjectile
// 	}
// 	return nil
// }

func (x *XMLObjectProperties) GetSize() int32 {
	if x.minSize == x.maxSize {
		return x.minSize
	}
	return utils.NextInt32(x.minSize/x.sizeStep, x.minSize/x.sizeStep) * x.sizeStep
}

func (x *XMLObjectProperties) Parse(elem *etree.Element) {
	x.DisplayId = StringElementValue(elem, "DisplayId", x.IdName)

	x.Size = IntElementValue(elem, "size", 100)
	x.minSize = IntElementValue(elem, "MinSize", 100)
	x.maxSize = IntElementValue(elem, "MaxSize", 100)
	x.sizeStep = IntElementValue(elem, "SizeStep", 5)

	x.IsItem = HasElement(elem, "Item")
	x.IsPlayer = HasElement(elem, "Player")
	x.IsStatic = HasElement(elem, "Static")
	x.IsEnemy = HasElement(elem, "Enemy")
	x.IsFriendly = HasElement(elem, "Friendly")
	x.BlocksSight = HasElement(elem, "BlocksSight")
	x.FullOccupy = HasElement(elem, "FullOccupy")
	x.OccupySquare = HasElement(elem, "OccupySquare")
	x.EnemyOccupySquare = HasElement(elem, "EnemyOccupySquare")
	x.InterGamePortal = HasElement(elem, "InterGamePortal")
	x.IsContainer = HasElement(elem, "Container")
	x.CanPutNormalObjects = HasElement(elem, "CanPutNormalObjects")
	x.CanPutSoulboundObjects = HasElement(elem, "CanPutSoulboundObjects")
	x.ProtectFromGroundDamage = HasElement(elem, "ProtectFromGroundDamage")
	x.ProtectFromGroundEffects = HasElement(elem, "ProtectFromGroundEffects")
	x.Connects = HasElement(elem, "Connects")
	x.IsLoot = HasElement(elem, "Loot")
	x.SpawnPoint = HasElement(elem, "SpawnPoint")
	x.CanUseTexes = HasElement(elem, "CanUseTexes")
	x.LeachHealth = HasElement(elem, "LeachHealth")

	x.BlockNearbyPlayerPortalCreation = HasElement(elem, "BlockPlayerPortals")
	x.PortalBlockRadius = FloatElementValue(elem, "BlockPlayerPortals", 6.0)

	x.DungeonName = StringElementValue(elem, "DungeonName", "")
	x.GuildItem = StringElementValue(elem, "GuildItem", "")
	x.GuildItemParams = StringElementValue(elem, "GuildItemParams", "")
	x.YardType = IntElementValue(elem, "YardType", 0)
	x.Price = IntElementValue(elem, "Price", 0)
	x.Fame = IntElementValue(elem, "Fame", 0)
	x.SlotTypes = IntArrayElementValue(elem, "SlotTypes")
	x.StartingEquipment = IntArrayElementValue(elem, "Equipment")
}
