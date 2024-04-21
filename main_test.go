package main

import (
  "fmt"
  "testing"
)

func TestCalculator(t *testing.T) {
  t.Run("TestAddition", func(t *testing.T) {
    result := addition(3, 5)
    if result != 8 {
      t.Errorf("Expected result: 8, but got %f", result)
    }
  })

  t.Run("TestSubtraction", func(t *testing.T) {
    result := subtraction(10, 4)
    if result != 6 {
      t.Errorf("Expected result: 6, but got %f", result)
    }
  })

  t.Run("TestMultiplication", func(t *testing.T) {
    result := multiplication(6, 4)
    if result != 24 {
      t.Errorf("Expected result: 24, but got %f", result)
    }
  })

  t.Run("TestDivision", func(t *testing.T) {
    result, err := division(10, 2)
    if err != nil {
      t.Errorf("Error encountered: %v", err)
    }
    if result != 5 {
      t.Errorf("Expected result: 5, but got %f", result)
    }
  })

  t.Run("TestDivisionByZero", func(t *testing.T) {
    _, err := division(10, 0)
    if err == nil {
      t.Errorf("Expected error for division by zero, but got nil")
    }
  })
}

func addition(a, b float64) float64 {
  return a + b
}

func subtraction(a, b float64) float64 {
  return a - b
}

func multiplication(a, b float64) float64 {
  return a * b
}

func division(a, b float64) (float64, error) {
  if b == 0 {
    return 0, fmt.Errorf("Division by zero")
  }
  return a / b, nil
}
