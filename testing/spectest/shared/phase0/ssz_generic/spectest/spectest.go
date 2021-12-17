package ssz_generic

/*type SingleFieldTestStruct struct {
	A byte
}*/

type SmallTestStruct struct {
	A uint16
	B uint16
}

type FixedTestStruct struct {
	A uint8
	B uint64
	C uint32
}

type VarTestStruct struct {
	A uint16
	B []uint16 `json:"b" ssz-max:"1024"`
	C uint8
}

type ComplexTestStruct struct {
	A uint16
	B []uint16 `json:"b" ssz-max:"128"`
	C uint8
	D []byte           `json:"d" ssz-max:"256"`
	E *VarTestStruct   `json:"e"`
	F *FixedTestStruct `json:"f ssz-size:"4"`
	G *VarTestStruct   `json:"g ssz-size:"2"`
}

//TODO BitsStruct
