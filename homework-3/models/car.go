package models

import (
	"errors"
	"fmt"
)

// Engine -
type Engine struct {
	Acceleration int
	FuelCapacity int
	FuelRemains  int
	running      bool
}

// Body -
type Body struct {
	Color         string
	Capacity      int
	TrunkCapacity int
	windowsState  [4]bool
}

// Car -
type Car struct {
	Engine
	Body
	Name       string
	IssuedYear int
}

// LeftFrontWindow
const (
	LeftFrontWindow = iota
	RightFrontWindow
	LeftBackWindow
	RightBackWindow
)

// TurnOn -
func (engine *Engine) TurnOn() {
	engine.running = true
}

// OpenWindow -
func (body *Body) OpenWindow(windowID byte) {
	body.windowsState[windowID] = true
}

// Fill -
func (engine *Engine) Fill(count int) error {
	countAfterFilling := engine.FuelRemains + count

	if countAfterFilling > engine.FuelCapacity {
		engine.FuelRemains = engine.FuelCapacity
		return errors.New("over the capacity")
	}

	engine.FuelRemains = countAfterFilling

	return nil
}

func (car Car) String() string {
	carModel := fmt.Sprintf("Model: %v", car.Name)
	issuedYear := fmt.Sprintf("Issued: %v year", car.IssuedYear)
	fuelPercent := fmt.Sprintf("Fuel: %v%%", float32(car.FuelRemains)/float32(car.FuelCapacity)*100)
	runningState := fmt.Sprintf("Is running: %v", car.running)

	openedWindowsCount := 0
	for _, windowState := range car.windowsState {
		if windowState {
			openedWindowsCount++
		}
	}

	windowsOpened := ""
	if openedWindowsCount > 0 {
		windowsOpened = fmt.Sprintf("Windows aren't closed\n")
	}

	return fmt.Sprintf("%v\n%v\n%v\n%v\n%v", carModel, issuedYear, fuelPercent, runningState, windowsOpened)
}
