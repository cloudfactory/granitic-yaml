package target

type AllSupported struct {
	Int         int64
	Float       float64
	Bool        bool
	String      string
	Nested      *AllSupported
	StringSlice []string
	BoolSlice   []bool
	IntSlice    []int64
	FloatSlice  []float64
}
