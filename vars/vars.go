package vars

var (
	DOOM_RES          = []int{320, 200}
	SCALE             = float32(4.0)
	WIN_RES           = []int{DOOM_RES[0] * int(SCALE), DOOM_RES[1] * int(SCALE)}
	H_wIDTH, H_HEIGHT = WIN_RES[0] / 2, WIN_RES[1] / 2
)
