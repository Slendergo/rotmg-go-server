package assets

import "github.com/beevik/etree"

type XMLGround struct {
	IdName    string
	Type      int32
	DisplayId string
	Speed     float32
	MinDamage int32
	MaxDamage int32
	NoWalk    bool
	Sinking   bool
}

func NewXMLGround() *XMLGround {
	return &XMLGround{}
}

func (x *XMLGround) Parse(elem *etree.Element) {
	x.DisplayId = StringElementValue(elem, "DisplayId", x.IdName)
	x.Speed = FloatElementValue(elem, "Speed", 1.0)
	x.MinDamage = IntElementValue(elem, "MinDamage", 0)
	x.MaxDamage = IntElementValue(elem, "MaxDamage", 0)
	x.NoWalk = HasElement(elem, "NoWalk")
	x.Sinking = HasElement(elem, "Sinking")
}
