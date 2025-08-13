package animal

import "fmt"

// Animal interface defines common behavior for all animals
type Animal interface {
	// Sound returns the animal's vocalization
	Sound() string
	// Describe returns a description of the animal
	Describe() string
	// Feed feeds the animal (method chaining support)
	Feed() Animal
	// Exercise makes the animal exercise (method chaining support)
	Exercise() Animal
	// GetStatus returns the current status of the animal
	GetStatus() string
	// ChangeName changes the animal's name
	ChangeName(newName string) Animal
}

// BaseAnimal is the base structure for all animals
type BaseAnimal struct {
	Name  string
	Age   int
	Fed   bool
	Tired bool
}

// Dog represents a dog with breed information
type Dog struct {
	BaseAnimal
	Breed string // Dog breed
}

// Cat represents a cat with color information
type Cat struct {
	BaseAnimal
	Color string // Fur color
}

// Constants for default values
const (
	DefaultBreed = "Mixed"
	DefaultColor = "Brown"
)

// NewDog creates a new dog instance.
// name: the dog's name
// age: the dog's age
// Returns: pointer to initialized dog
func NewDog(name string, age int) *Dog {
	return &Dog{
		BaseAnimal: BaseAnimal{
			Name:  name,
			Age:   age,
			Fed:   false,
			Tired: false,
		},
		Breed: DefaultBreed,
	}
}

// NewCat creates a new cat instance.
// name: the cat's name
// age: the cat's age
// Returns: pointer to initialized cat
func NewCat(name string, age int) *Cat {
	return &Cat{
		BaseAnimal: BaseAnimal{
			Name:  name,
			Age:   age,
			Fed:   false,
			Tired: false,
		},
		Color: DefaultColor,
	}
}

// Sound returns the dog's bark sound.
func (d *Dog) Sound() string {
	return "Woof Woof"
}

// Describe returns detailed information about the dog as a string.
func (d *Dog) Describe() string {
	return fmt.Sprintf("Dog named %s (%d years old, %s breed)", d.Name, d.Age, d.Breed)
}

// Feed feeds the dog and makes it full. Supports method chaining.
func (d *Dog) Feed() Animal {
	d.Fed = true
	return d // For method chaining
}

// Exercise makes the dog exercise and become tired. Supports method chaining.
func (d *Dog) Exercise() Animal {
	d.Tired = true
	return d // For method chaining
}

// GetStatus returns the dog's current status as a string.
// Status is determined based on feeding and exercise state.
func (d *Dog) GetStatus() string {
	// Short variable declaration
	status := "Energetic"
	if d.Fed && d.Tired {
		status = "Satisfied"
	} else if d.Fed {
		status = "Full"
	} else if d.Tired {
		status = "Tired"
	}
	return status
}

// SetBreed sets the dog's breed.
func (d *Dog) SetBreed(breed string) {
	d.Breed = breed
}

// ChangeName changes the dog's name. Supports method chaining.
func (d *Dog) ChangeName(newName string) Animal {
	d.Name = newName
	return d
}

// NewAnimal creates a new animal instance based on the animal type.
// name: the animal's name
// age: the animal's age
// animalType: type of animal ("dog" or "cat")
// Returns: created animal interface, nil for unknown types
func NewAnimal(name string, age int, animalType string) Animal {
	switch animalType {
	case "dog":
		return NewDog(name, age)
	case "cat":
		return NewCat(name, age)
	default:
		return nil // Unknown animal type
	}
}

// Sound returns the cat's meow sound.
func (c *Cat) Sound() string {
	return "Meow Meow"
}

// Describe returns detailed information about the cat as a string.
func (c *Cat) Describe() string {
	return fmt.Sprintf("Cat named %s (%d years old, %s color)", c.Name, c.Age, c.Color)
}

// Feed feeds the cat and makes it full. Supports method chaining.
func (c *Cat) Feed() Animal {
	c.Fed = true
	return c // For method chaining
}

// Exercise makes the cat exercise and become tired. Supports method chaining.
func (c *Cat) Exercise() Animal {
	c.Tired = true
	return c // For method chaining
}

// GetStatus returns the cat's current status as a string.
// Status is determined based on feeding and exercise state.
func (c *Cat) GetStatus() string {
	// Short variable declaration
	status := "Relaxed"
	if c.Fed && c.Tired {
		status = "Sleepy"
	} else if c.Fed {
		status = "Satisfied"
	} else if c.Tired {
		status = "Lazy"
	}
	return status
}

// SetColor sets the cat's fur color.
func (c *Cat) SetColor(color string) {
	c.Color = color
}

// ChangeName changes the cat's name. Supports method chaining.
func (c *Cat) ChangeName(newName string) Animal {
	c.Name = newName
	return c
}
