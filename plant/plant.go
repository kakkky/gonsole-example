package plant

import "fmt"

// Plant interface defines common behavior for all plants
type Plant interface {
	Grow() string
	Bloom() string
	Describe() string
	Water() Plant     // For method chaining
	Fertilize() Plant // For method chaining
	GetHealth() string
	ChangeName(newName string) Plant // Change plant name
}

// BasePlant is the base structure for all plants
type BasePlant struct {
	Name    string
	Age     int // Age in years
	Watered bool
	Healthy bool
}

// Flower represents a flower with color information
type Flower struct {
	BasePlant
	Color string // Flower color
}

// Tree represents a tree with height information
type Tree struct {
	BasePlant
	Height float64 // Height in meters
}

// Constants for default values
const (
	DefaultColor  = "Red"
	DefaultHeight = 1.0
)

// NewFlower creates a new flower instance.
func NewFlower(name string, age int) *Flower {
	return &Flower{
		BasePlant: BasePlant{
			Name:    name,
			Age:     age,
			Watered: false,
			Healthy: true,
		},
		Color: DefaultColor,
	}
}

// NewTree creates a new tree instance.
func NewTree(name string, age int) *Tree {
	return &Tree{
		BasePlant: BasePlant{
			Name:    name,
			Age:     age,
			Watered: false,
			Healthy: true,
		},
		Height: DefaultHeight,
	}
}

// Grow makes the flower grow and increases its age.
func (f *Flower) Grow() string {
	f.Age++
	return "The flower is growing beautifully"
}

// Bloom makes the flower bloom.
func (f *Flower) Bloom() string {
	return fmt.Sprintf("Beautiful %s flowers have bloomed", f.Color)
}

// Describe returns detailed information about the flower.
func (f *Flower) Describe() string {
	return fmt.Sprintf("%s (%d years old, %s flowers)", f.Name, f.Age, f.Color)
}

// Water waters the flower. Supports method chaining.
func (f *Flower) Water() Plant {
	f.Watered = true
	return f // For method chaining
}

// Fertilize fertilizes the flower. Supports method chaining.
func (f *Flower) Fertilize() Plant {
	f.Healthy = true
	return f // For method chaining
}

// GetHealth returns the flower's health status.
func (f *Flower) GetHealth() string {
	// Short variable declaration
	health := "Normal"
	if f.Watered && f.Healthy {
		health = "Very Healthy"
	} else if f.Watered {
		health = "Well Watered"
	} else if f.Healthy {
		health = "Healthy"
	} else {
		health = "Needs Care"
	}
	return health
}

// SetColor sets the flower's color.
func (f *Flower) SetColor(color string) {
	f.Color = color
}

// ChangeName changes the flower's name. Supports method chaining.
func (f *Flower) ChangeName(newName string) Plant {
	f.Name = newName
	return f
}

// Grow makes the tree grow, increasing its age and height.
func (t *Tree) Grow() string {
	t.Age++
	t.Height += 0.5
	return "The tree is growing bigger and taller"
}

// Bloom makes the tree bloom.
func (t *Tree) Bloom() string {
	return "Small flowers have bloomed on the tree"
}

// Describe returns detailed information about the tree.
func (t *Tree) Describe() string {
	return fmt.Sprintf("%s (%d years old, %.1fm tall)", t.Name, t.Age, t.Height)
}

// Water waters the tree. Supports method chaining.
func (t *Tree) Water() Plant {
	t.Watered = true
	return t // For method chaining
}

// Fertilize fertilizes the tree. Supports method chaining.
func (t *Tree) Fertilize() Plant {
	t.Healthy = true
	return t // For method chaining
}

// GetHealth returns the tree's health status.
func (t *Tree) GetHealth() string {
	// Short variable declaration
	health := "Normal"
	if t.Watered && t.Healthy {
		health = "Lush and Green"
	} else if t.Watered {
		health = "Well Watered"
	} else if t.Healthy {
		health = "Strong"
	} else {
		health = "Withering"
	}
	return health
}

// SetHeight sets the tree's height.
func (t *Tree) SetHeight(height float64) {
	t.Height = height
}

// ChangeName changes the tree's name. Supports method chaining.
func (t *Tree) ChangeName(newName string) Plant {
	t.Name = newName
	return t
}
