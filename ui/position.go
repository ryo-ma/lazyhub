package ui

type Position struct {
	prc    float32
	margin int
}

func (p Position) getCoordinate(max int) int {
	return int(p.prc*float32(max)) - p.margin
}

type ViewPosition struct {
	x0, y0, x1, y1 Position
}

func (vp ViewPosition) GetCoordinates(maxX int, maxY int) (int, int, int, int) {
	x0 := vp.x0.getCoordinate(maxX)
	y0 := vp.y0.getCoordinate(maxY)
	x1 := vp.x1.getCoordinate(maxX)
	y1 := vp.y1.getCoordinate(maxY)
	return x0, y0, x1, y1
}

//var viewPositions = map[string]viewPosition{
//	RepositoryView: {
//		position{0.0, 0},
//		position{0.0, 0},
//		position{0.3, 2},
//		position{0.9, 2},
//	},
//	TextView: {
//		position{0.3, 0},
//		position{0.0, 0},
//		position{1.0, 2},
//		position{0.9, 2},
//	},
//	PathView: {
//		position{0.0, 0},
//		position{0.89, 0},
//		position{1.0, 2},
//		position{1.0, 2},
//	},
//	SearchView: {
//		position{0.1, 0},
//		position{0.35, 0},
//		position{0.9, 2},
//		position{0.5, 2},
//	},
//}
