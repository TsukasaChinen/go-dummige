package structs

type Images struct {
	S string `query:"size"`
	C string `query:"color"`
	N string
	W int
	H int
}

type RGB struct {
	R uint8
	G uint8
	B uint8
}
