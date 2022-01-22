package data

type Types struct {
	// // numerics
	// Int8    int8
	// Int16   int16
	// Int32   int32
	// Int64   int64
	// UInt8   uint8
	// UInt16  uint16
	// UInt32  uint32
	// UInt64  uint64
	// Int     int
	// UInt    uint
	// Rune    rune
	// Byte    byte
	// Float32 float32
	// Float64 float64
	// IntPtr  *int
	// UIntPtr *uint

	// // string
	// String    string
	// StringPtr *string

	// structs
	Struct Struct
	// StructPtr *Struct

	// // arrays
	// IntArray       []int
	// StringArray    []string
	// StructArray    []Struct
	// IntPtrArray    []*int
	// StringPtrArray []*string
	// StructPtrArray []*Struct

	// // others
	// Bool    bool
	// BoolPtr *bool
}

type Struct struct {
	One string
	Two string
}
