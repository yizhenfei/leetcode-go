package leetcode

// Coorditaion of grid
type coord struct {
	X int
	Y int
}

// Context for the calculation
type context struct {
	Width           int
	Height          int
	Grids           [][]byte
	IsMerged        [][]byte
	NumberOfIslands int
}

func newContext(grid [][]byte) *context {
	c := context{
		Width:           len(grid[0]),
		Height:          len(grid),
		Grids:           grid,
		NumberOfIslands: 0,
	}

	isMerged := make([][]byte, c.Height)
	for i := 0; i < c.Height; i++ {
		isMerged[i] = make([]byte, c.Width)
	}

	c.IsMerged = isMerged

	return &c
}

type coordQueue struct {
	coords []coord
}

func newCoordQueue() *coordQueue {
	return &coordQueue{
		coords: []coord{},
	}
}

func (q *coordQueue) push(coord coord) {
	q.coords = append(q.coords, coord)
}

func (q *coordQueue) pop() coord {
	c := q.coords[0]
	q.coords = q.coords[1:]
	return c
}

func (q *coordQueue) isEmpty() bool {
	return len(q.coords) == 0
}

func isUnmergedLand(coord coord, c *context) bool {
	// If is out of range, regard as ocean
	if coord.X < 0 || coord.X >= c.Width || coord.Y < 0 || coord.Y >= c.Height {
		return false
	}

	// If is ocean, return false
	if c.Grids[coord.Y][coord.X] == '0' {
		return false
	}

	// If is merged, return false
	if c.IsMerged[coord.Y][coord.X] == 1 {
		return false
	}

	return true
}

func mergeLand(coord coord, c *context, q *coordQueue) {
	// Mark as merged
	c.IsMerged[coord.Y][coord.X] = 1

	// Push into queue
	q.push(coord)
}

func isNewIsland(coord coord, c *context) bool {
	return isUnmergedLand(coord, c)
}

func maybeMerge(coord coord, c *context, q *coordQueue) {
	if isUnmergedLand(coord, c) {
		mergeLand(coord, c, q)
	}
}

func mergeAround(co coord, c *context, q *coordQueue) {
	maybeMerge(coord{co.X + 1, co.Y}, c, q)
	maybeMerge(coord{co.X - 1, co.Y}, c, q)
	maybeMerge(coord{co.X, co.Y + 1}, c, q)
	maybeMerge(coord{co.X, co.Y - 1}, c, q)
}

func mergeIsland(coord coord, c *context) {
	q := newCoordQueue()

	// Merge initial lands
	mergeLand(coord, c, q)

	// Expand island
	for !q.isEmpty() {
		cur := q.pop()
		mergeAround(cur, c, q)
	}
}

func countIslandByGrid(coord coord, c *context) {
	if isNewIsland(coord, c) {
		c.NumberOfIslands++
		mergeIsland(coord, c)
	}
}

func numIslands(grid [][]byte) int {
	// Validate input
	if grid == nil || len(grid) == 0 || grid[0] == nil || len(grid[0]) == 0 {
		return 0
	}

	// Setup context
	c := newContext(grid)

	// Iterate over all grids and count island
	for y := 0; y < c.Height; y++ {
		for x := 0; x < c.Width; x++ {
			countIslandByGrid(coord{x, y}, c)
		}
	}

	// Return result
	return c.NumberOfIslands
}
