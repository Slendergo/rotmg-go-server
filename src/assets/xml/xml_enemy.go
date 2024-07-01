package assets

import (
	"main/utils"
	"math"

	"github.com/beevik/etree"
)

type XMLSpawn struct {
	count  int32
	mean   float32
	stdDev float32
	min    int32
	max    int32
}

func newXMLSpawn(elem *etree.Element) *XMLSpawn {
	xmlSpawn := &XMLSpawn{}
	xmlSpawn.count = IntElementValue(elem, "Count", -1)
	xmlSpawn.mean = FloatElementValue(elem, "Mean", 1)
	xmlSpawn.stdDev = FloatElementValue(elem, "StdDev", 0.0)
	xmlSpawn.min = IntElementValue(elem, "Min", 1)
	xmlSpawn.max = IntElementValue(elem, "Max", 1)
	return xmlSpawn
}

func (x *XMLSpawn) NumToSpawn() int32 {
	if x.count != -1 {
		return x.count
	}

	if x.min == x.max {
		return x.min
	}

	return int32(math.Min(float64(x.min), math.Max(float64(x.max), float64(utils.Normal32(x.mean, x.stdDev)))))
}

type XMLEnemy struct {
	IdName                  string
	Type                    int32
	Group                   string
	Size                    int32
	Exp                     int32
	MaxHitPoints            int32
	Level                   int32
	Defense                 int32
	PerRealmMax             int32
	Spawn                   *XMLSpawn
	SpawnProb               float32
	IsStatic                bool
	IsHero                  bool
	IsQuest                 bool
	IsGod                   bool
	IsEncounter             bool
	IsStasisImmune          bool
	IsStunImmune            bool
	IsParalyzeImmune        bool
	IsSlowImmune            bool
	IsDazedImmune           bool
	IsInvincible            bool
	KeepDamageRecord        bool
	NoCollisionWhenMovement bool
	AlwaysPositiveHealth    bool
}

func NewXMLEnemy() *XMLEnemy {
	return &XMLEnemy{}
}

func (x *XMLEnemy) Parse(elem *etree.Element) {
	x.Group = StringElementValue(elem, "Group", "")
	x.Size = IntElementValue(elem, "Size", 0)
	x.Exp = IntElementValue(elem, "Exp", 0)
	x.MaxHitPoints = IntElementValue(elem, "MaxHitPoints", 100)
	x.Level = IntElementValue(elem, "Level", 1)
	x.Defense = IntElementValue(elem, "Defense", 0)
	x.PerRealmMax = IntElementValue(elem, "PerRealmMax", -1)
	x.Spawn = newXMLSpawn(elem)
	x.SpawnProb = FloatElementValue(elem, "SpawnProb", 0.0)
	x.IsStatic = HasElement(elem, "Static")
	x.IsHero = HasElement(elem, "Hero")
	x.IsQuest = HasElement(elem, "Quest")
	x.IsGod = HasElement(elem, "God")
	x.IsEncounter = HasElement(elem, "Encounter")
	x.IsStasisImmune = HasElement(elem, "StasisImmune")
	x.IsStunImmune = HasElement(elem, "StunImmune")
	x.IsParalyzeImmune = HasElement(elem, "ParalyzeImmune")
	x.IsSlowImmune = HasElement(elem, "SlowImmune")
	x.IsDazedImmune = HasElement(elem, "DazedImmune")
	x.IsInvincible = HasElement(elem, "Invincible")
	x.KeepDamageRecord = HasElement(elem, "KeepDamageRecord")
	x.NoCollisionWhenMovement = HasElement(elem, "NoCollisionWhenMovement")
	x.AlwaysPositiveHealth = HasElement(elem, "AlwaysPositiveHealth")
}
