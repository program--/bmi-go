package bmi

const BMI_SUCCESS bool = true
const BMI_FAILURE bool = false
const MAX_COMPONENT_NAME int = 2048
const MAX_VAR_NAME int = 2048
const MAX_UNITS_NAME int = 2048
const MAX_TYPE_NAME int = 2048

type Bmi interface {
	// Model control functions
	Initialize(config_file string)
	Update()
	UpdateUntil(time float64)
	Finalize()

	// Model information functions
	GetComponentName() string
	GetInputItemCount() int
	GetOutputItemCount() int
	GetInputVarNames() []string
	GetOutputVarNames() []string

	// Variable information functions
	GetVarGrid(name string) int
	GetVarType(name string) string
	GetVarUnits(name string) string
	GetVarItemsize(name string) int
	GetVarNbytes(name string) int
	GetVarLocation(name string) string

	GetCurrentTime() float64
	GetStartTime() float64
	GetEndTime() float64
	GetTimeUnits() string
	GetTimeStep() float64

	// Variable getters
	GetValue(name string, dest any)
	GetValuePtr(name string) any
	GetValueAtIndices(name string, dest any, inds []int, count int)

	// Variable setters
	SetValue(name string, src any)
	SetValueAtIndices(name string, inds []int, count int, src any)

	// Grid information functions
	GetGridRank(grid int) int
	GetGridSize(grid int) int
	GetGridType(grid int) string

	GetGridShape(grid int, shape []int)
	GetGridSpacing(grid int, spacing []float64)
	GetGridOrigin(grid int, origin []float64)

	GetGridX(grid int, x []float64)
	GetGridY(grid int, y []float64)
	GetGridZ(grid int, z []float64)

	GetGridNodeCount(grid int) int
	GetGridEdgeCount(grid int) int
	GetGridFaceCount(grid int) int

	GetGridEdgeNodes(grid int, edge_nodes []int)
	GetGridFaceEdges(grid int, face_edges []int)
	GetGridFaceNodes(grid int, face_nodes []int)
	GetGridNodesPerFace(grid int, nodes_per_face []int)
}
