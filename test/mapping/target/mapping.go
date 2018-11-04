package target

type AllSupported struct {
	Int         int
	Float       float64
	Bool        bool
	String      string
	Nested      *AllSupported
	StringSlice []string
	BoolSlice   []bool
	IntSlice    []int
	FloatSlice  []float64
	MixedSlice  []interface{}
}
