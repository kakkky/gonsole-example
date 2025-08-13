# Gonsole Example

A simple Go language sample project that includes various Go elements such as interfaces, structs, variables, short variable declarations, methods, constants, and functions.

## Project Structure

```
gonsole-example/
├── go.mod              # Go module file
├── README.md           # This file
├── animal/
│   ├── animal.go       # Animal interface and implementations (dog, cat)
│   └── utils/
│       └── utils.go    # Animal-related utility functions
├── plant/
│   ├── plant.go        # Plant interface and implementations (flower, tree)
│   └── utils/
│       └── utils.go    # Plant-related utility functions
└── vehicle/
    ├── vehicle.go      # Vehicle interface and implementations (car, bike)
    └── utils/
        └── utils.go    # Vehicle-related utility functions
```

## Features

### Go Language Elements Included

1. **Interface**: `Animal`, `Plant`, `Vehicle` - Define common behaviors for each category
2. **Struct**: 
   - Base structures: `BaseAnimal`, `BasePlant`, `BaseVehicle`
   - Concrete structures: `Dog`, `Cat`, `Flower`, `Tree`, `Car`, `Bike`
3. **Var**: Package-level variable declarations
4. **Short variable declaration**: Variable declarations using `:=` (used within methods)
5. **Method**: Struct methods with method chaining support
6. **Const**: Constant declarations (`DefaultBreed`, `DefaultColor`, `DefaultEngine`, etc.)
7. **Function**: Package-level functions and utility functions

### Package Structure

#### animal package
- Defines basic animal behaviors (sound, describe, feed, exercise)
- Implementations for Dog and Cat
- Utilities for age conversion and name formatting

#### plant package
- Defines basic plant behaviors (grow, bloom, describe, water, fertilize)
- Implementations for Flower and Tree
- Utilities for growth stage determination and name formatting
#### vehicle package
- Defines basic vehicle behaviors (start, stop, describe, refuel, accelerate)
- Implementations for Car and Bike
- Utilities for age calculation and fuel efficiency calculation

### Method Chaining

All objects support method chaining:

```go
// Animals
dog := animal.NewDog("Buddy", 3)
dog.Feed().Exercise()

// Plants
flower := plant.NewFlower("Rose", 2)
flower.Water().Fertilize()

// Vehicles
car := vehicle.NewCar("Toyota", "Prius", 2020)
car.Refuel().Accelerate()
```

### Usage Example

```go
package main

import (
    "fmt"
    "gonsole-example/animal"
    "gonsole-example/plant"
    "gonsole-example/vehicle"
)

func main() {
    // Animal example
    dog := animal.NewDog("Buddy", 3)
    fmt.Println(dog.Describe())
    fmt.Println(dog.Sound())
    dog.Feed().Exercise()
    fmt.Println("Status:", dog.GetStatus())

    // Plant example
    flower := plant.NewFlower("Rose", 2)
    fmt.Println(flower.Describe())
    fmt.Println(flower.Bloom())
    flower.Water().Fertilize()
    fmt.Println("Health:", flower.GetHealth())

    // Vehicle example
    car := vehicle.NewCar("Toyota", "Prius", 2020)
    fmt.Println(car.Describe())
    fmt.Println(car.Start())
    car.Refuel().Accelerate()
    fmt.Println("Status:", car.GetStatus())
}
```

## Learning Points

This project provides practical learning opportunities for the following Go concepts:

1. **Interface Design**: Abstracting common behaviors
2. **Struct Embedding**: Inheritance-like design using base structs
3. **Method Receivers**: Method implementation using pointer receivers
4. **Method Chaining**: Implementation of fluent interfaces
5. **Package Organization**: Functional package structure
6. **Utility Functions**: Separation of common processing
7. **Constants and Constructors**: Object initialization patterns


