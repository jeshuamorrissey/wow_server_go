package components

type PlayerFeatures struct {
	SkinColor int
	Face      int
	HairStyle int
	HairColor int
	Feature   int
}

func (pf *PlayerFeatures) Bytes() uint32 {
	return uint32(pf.SkinColor) | uint32(pf.Face)<<8 | uint32(pf.HairStyle)<<16 | uint32(pf.HairColor)<<24
}

func (pf *PlayerFeatures) Bytes2() uint32 {
	return uint32(pf.Feature)
}
