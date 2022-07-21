package entity

import "github.com/labstack/gommon/log"

const (
	Bow = iota
	OneHandedSword
	TwoHandedSword
	Polearm
	MagicWeapon
)

func ToWeapon(v byte) (weapon string) {
	defer func() {
		if v := recover(); v != nil {
			log.Error("error weapon value ", v)
			weapon = ""
		}
	}()
	switch v {
	case 0:
		return "弓箭"
	case 1:
		return "单手剑"
	case 2:
		return "双手剑"
	case 3:
		return "长柄武器"
	case 4:
		return "法器"
	}
	panic(v)
}

func FromWeapon(desc string) (v byte) {
	defer func() {
		if desc := recover(); desc != nil {
			log.Error("error weapon desc ", desc)
			v = 5
		}
	}()
	switch desc {
	case "弓箭", "弓":
		return Bow
	case "单手剑":
		return OneHandedSword
	case "双手剑":
		return TwoHandedSword
	case "长柄武器":
		return Polearm
	case "法器":
		return MagicWeapon
	}
	panic(desc)
}

const (
	Water = iota
	Fire
	Ice
	Lightning
	Grass
	Earth
	Wind
)

func ToElementType(v byte) (elementType string) {
	defer func() {
		if v := recover(); v != nil {
			log.Error("error weapon value ", v)
			elementType = ""
		}
	}()
	switch v {
	case 0:
		return "水"
	case 1:
		return "火"
	case 2:
		return "冰"
	case 3:
		return "雷"
	case 4:
		return "草"
	case 5:
		return "岩"
	case 6:
		return "风"
	}
	panic(v)
}

func FromElementType(desc string) (v byte) {
	defer func() {
		if desc := recover(); desc != nil {
			log.Error("error element type value ", v)
			v = 7
		}
	}()
	switch desc {
	case "水":
		return Water
	case "火":
		return Fire
	case "冰":
		return Ice
	case "雷":
		return Lightning
	case "草":
		return Grass
	case "岩":
		return Earth
	case "风":
		return Wind
	}
	panic(desc)
}
