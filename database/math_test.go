package database

import (
	"testing"
)

// Test generated using Keploy
func TestIsHappyNumber_HappyNumber_027(t *testing.T) {
	result := isHappyNumber(19)
	if !result {
		t.Errorf("Expected true, got false")
	}
}

// Test generated using Keploy

func TestFibonacci_PositiveInteger_013(t *testing.T) {
	result, err := fibonacci(7)
	expected := []int{0, 1, 1, 2, 3, 5, 8}
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if len(result) != len(expected) {
		t.Errorf("Expected length %d, got %d", len(expected), len(result))
	}
	for i := range result {
		if result[i] != expected[i] {
			t.Errorf("At index %d, expected %d, got %d", i, expected[i], result[i])
		}
	}
}

// Test generated using Keploy

func TestPrimeFactors_CompositeNumber_019(t *testing.T) {
	result := primeFactors(28)
	expected := []int{2, 2, 7}
	if len(result) != len(expected) {
		t.Errorf("Expected length %d, got %d", len(expected), len(result))
	}
	for i := range result {
		if result[i] != expected[i] {
			t.Errorf("At index %d, expected %d, got %d", i, expected[i], result[i])
		}
	}
}

// Test generated using Keploy

func TestIsAnagram_AnagramStrings_023(t *testing.T) {
	result := isAnagram("listen", "silent")
	if !result {
		t.Errorf("Expected true, got false")
	}
}

// Test generated using Keploy

func TestFactorial_PositiveInteger_011(t *testing.T) {
	result, err := factorial(5)
	expected := 120
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if result != expected {
		t.Errorf("Expected %d, got %d", expected, result)
	}
}

// Test generated using Keploy

func TestIsPalindrome_PalindromeString_020(t *testing.T) {
	result := isPalindrome("racecar")
	if !result {
		t.Errorf("Expected true, got false")
	}
}

// Test generated using Keploy

func TestIsPerfectNumber_PerfectNumber_026(t *testing.T) {
	result := isPerfectNumber(28)
	if !result {
		t.Errorf("Expected true, got false")
	}
}

// Test generated using Keploy

func TestIsFibonacci_FibonacciNumber_028(t *testing.T) {
	result := isFibonacci(8)
	if !result {
		t.Errorf("Expected true, got false")
	}
}

// Test generated using Keploy

func TestIsPrime_PrimeNumber_017(t *testing.T) {
	result := isPrime(7)
	if !result {
		t.Errorf("Expected true, got false")
	}
}

// Test generated using Keploy

func TestReverseString_ValidString_022(t *testing.T) {
	result := reverseString("hello")
	expected := "olleh"
	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}
}

// Test generated using Keploy

// Test generated using Keploy

func TestIsPerfectSquare_PerfectSquareNumber_029(t *testing.T) {
	result := isPerfectSquare(16)
	if !result {
		t.Errorf("Expected true, got false")
	}
}

func TestDivide_PositiveIntegers_004(t *testing.T) {
	result, err := divide(10, 2)
	expected := 5
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if result != expected {
		t.Errorf("Expected %d, got %d", expected, result)
	}
}

// Test generated using Keploy

func TestSquareRoot_PositiveInteger_009(t *testing.T) {
	result, err := squareRoot(16)
	expected := 4
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if result != expected {
		t.Errorf("Expected %d, got %d", expected, result)
	}
}

// Test generated using Keploy

func TestModulus_PositiveIntegers_006(t *testing.T) {
	result, err := modulus(10, 3)
	expected := 1
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if result != expected {
		t.Errorf("Expected %d, got %d", expected, result)
	}
}

// Test generated using Keploy

func TestSum_PositiveIntegers_001(t *testing.T) {
	result := sum(3, 5)
	expected := 8
	if result != expected {
		t.Errorf("Expected %d, got %d", expected, result)
	}
}

// Test generated using Keploy

func TestModulus_ZeroDivisor_007(t *testing.T) {
	_, err := modulus(10, 0)
	if err == nil {
		t.Errorf("Expected error, got nil")
	}
}

// Test generated using Keploy

func TestFactorial_NegativeInteger_012(t *testing.T) {
	_, err := factorial(-5)
	if err == nil {
		t.Errorf("Expected error, got nil")
	}
}

// Test generated using Keploy

func TestIsPrime_NonPrimeNumber_018(t *testing.T) {
	result := isPrime(10)
	if result {
		t.Errorf("Expected false, got true")
	}
}

// Test generated using Keploy

func TestIsPalindrome_NonPalindromeString_021(t *testing.T) {
	result := isPalindrome("hello")
	if result {
		t.Errorf("Expected false, got true")
	}
}

// Test generated using Keploy

func TestSubtract_PositiveIntegers_002(t *testing.T) {
	result := subtract(10, 4)
	expected := 6
	if result != expected {
		t.Errorf("Expected %d, got %d", expected, result)
	}
}

// Test generated using Keploy

func TestMultiply_PositiveIntegers_003(t *testing.T) {
	result := multiply(3, 5)
	expected := 15
	if result != expected {
		t.Errorf("Expected %d, got %d", expected, result)
	}
}

// Test generated using Keploy

func TestDivide_ZeroDivisor_005(t *testing.T) {
	_, err := divide(10, 0)
	if err == nil {
		t.Errorf("Expected error, got nil")
	}
}

// Test generated using Keploy

func TestSquareRoot_NegativeInteger_010(t *testing.T) {
	_, err := squareRoot(-16)
	if err == nil {
		t.Errorf("Expected error, got nil")
	}
}

// Test generated using Keploy

func TestFibonacci_NegativeInteger_014(t *testing.T) {
	_, err := fibonacci(-5)
	if err == nil {
		t.Errorf("Expected error, got nil")
	}
}

// Test generated using Keploy

func TestLCM_PositiveIntegers_016(t *testing.T) {
	result := lcm(4, 6)
	expected := 12
	if result != expected {
		t.Errorf("Expected %d, got %d", expected, result)
	}
}

// Test generated using Keploy

func TestIsAnagram_NonAnagramStrings_024(t *testing.T) {
	result := isAnagram("hello", "world")
	if result {
		t.Errorf("Expected false, got true")
	}
}

// Test generated using Keploy

func TestPower_PositiveIntegers_008(t *testing.T) {
	result := power(2, 3)
	expected := 8
	if result != expected {
		t.Errorf("Expected %d, got %d", expected, result)
	}
}

// Test generated using Keploy
