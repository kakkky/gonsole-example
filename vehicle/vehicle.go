package vehicle

import "fmt"

// Vehicle interface defines common behavior for all vehicles
type Vehicle interface {
	Start() string
	Stop() string
	Describe() string
	Refuel() Vehicle     // For method chaining
	Accelerate() Vehicle // For method chaining
	GetStatus() string
	ChangeName(newBrand, newModel string) Vehicle // Change vehicle name (brand and model)
}

// BaseVehicle is the base structure for all vehicles
type BaseVehicle struct {
	Brand    string
	Model    string
	Year     int
	Running  bool
	FuelFull bool
}

// Car represents a car with engine information
type Car struct {
	BaseVehicle
	Engine string // Engine type
}

// Bike represents a motorcycle with type information
type Bike struct {
	BaseVehicle
	Type string // Bike type
}

// Constants for default values
const (
	DefaultEngine = "Gasoline"
	DefaultType   = "Sport"
)

// NewCar creates a new car instance.
func NewCar(brand, model string, year int) *Car {
	return &Car{
		BaseVehicle: BaseVehicle{
			Brand:    brand,
			Model:    model,
			Year:     year,
			Running:  false,
			FuelFull: false,
		},
		Engine: DefaultEngine,
	}
}

// NewBike creates a new bike instance.
func NewBike(brand, model string, year int) *Bike {
	return &Bike{
		BaseVehicle: BaseVehicle{
			Brand:    brand,
			Model:    model,
			Year:     year,
			Running:  false,
			FuelFull: false,
		},
		Type: DefaultType,
	}
}

// Start starts the car engine.
func (c *Car) Start() string {
	c.Running = true
	return "Vroom vroom... Engine started"
}

// Stop stops the car engine.
func (c *Car) Stop() string {
	c.Running = false
	return "Engine stopped"
}

// Describe returns detailed information about the car.
func (c *Car) Describe() string {
	return fmt.Sprintf("%s %s (%d model year, %s engine)", c.Brand, c.Model, c.Year, c.Engine)
}

// Refuel refuels the car. Supports method chaining.
func (c *Car) Refuel() Vehicle {
	c.FuelFull = true
	return c // For method chaining
}

// Accelerate accelerates the car. Supports method chaining.
func (c *Car) Accelerate() Vehicle {
	c.Running = true
	return c // For method chaining
}

// GetStatus returns the car's current status.
func (c *Car) GetStatus() string {
	// Short variable declaration
	status := "Parked"
	if c.Running && c.FuelFull {
		status = "Driving Smoothly"
	} else if c.Running {
		status = "Driving (Low Fuel)"
	} else if c.FuelFull {
		status = "Parked (Full Tank)"
	}
	return status
}

// SetEngine sets the car's engine type.
func (c *Car) SetEngine(engine string) {
	c.Engine = engine
}

// ChangeName changes the car's brand and model. Supports method chaining.
func (c *Car) ChangeName(newBrand, newModel string) Vehicle {
	c.Brand = newBrand
	c.Model = newModel
	return c
}

// Start starts the bike engine.
func (b *Bike) Start() string {
	b.Running = true
	return "Vrooom vrooom... Bike started"
}

// Stop stops the bike engine.
func (b *Bike) Stop() string {
	b.Running = false
	return "Bike stopped"
}

// Describe returns detailed information about the bike.
func (b *Bike) Describe() string {
	return fmt.Sprintf("%s %s (%d model year, %s type)", b.Brand, b.Model, b.Year, b.Type)
}

// Refuel refuels the bike. Supports method chaining.
func (b *Bike) Refuel() Vehicle {
	b.FuelFull = true
	return b // For method chaining
}

// Accelerate accelerates the bike. Supports method chaining.
func (b *Bike) Accelerate() Vehicle {
	b.Running = true
	return b // For method chaining
}

// GetStatus returns the bike's current status.
func (b *Bike) GetStatus() string {
	// Short variable declaration
	status := "Parked"
	if b.Running && b.FuelFull {
		status = "Cruising"
	} else if b.Running {
		status = "Riding (Low Fuel)"
	} else if b.FuelFull {
		status = "Parked (Full Tank)"
	}
	return status
}

// SetType sets the bike's type.
func (b *Bike) SetType(bikeType string) {
	b.Type = bikeType
}

// ChangeName changes the bike's brand and model. Supports method chaining.
func (b *Bike) ChangeName(newBrand, newModel string) Vehicle {
	b.Brand = newBrand
	b.Model = newModel
	return b
}
