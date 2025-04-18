package vec

type VecI struct {
	X, Y int
}

type VecF32 struct {
	X, Y float32
}

type VecF64 struct {
	X, Y float64
}

func NewI(x, y int) *VecI {
	return &VecI{ X: x, Y: y }
}

func NewF32(x, y float32) *VecF32 {
	return &VecF32{ X: x, Y: y }
}

func NewF64(x, y float64) *VecF64 {
	return &VecF64{ X: x, Y: y }
}

func (v *VecI) ToVecF32() *VecF32 {
	return &VecF32{
		X: float32(v.X),
		Y: float32(v.Y),
	}
}

func (v *VecF64) ToVecF32() *VecF32 {
	return &VecF32{
		X: float32(v.X),
		Y: float32(v.Y),
	}
}

func (v *VecI) ToVecF64() *VecF64 {
	return &VecF64{
		X: float64(v.X),
		Y: float64(v.Y),
	}
}

func (v *VecF32) ToVecF64() *VecF64 {
	return &VecF64{
		X: float64(v.X),
		Y: float64(v.Y),
	}
}

func (v *VecF32) ToVecI() *VecI {
	return &VecI{
		X: int(v.X),
		Y: int(v.Y),
	}
}

func (v *VecI) ToVecI() *VecI {
	return &VecI{
		X: int(v.X),
		Y: int(v.Y),
	}
}
