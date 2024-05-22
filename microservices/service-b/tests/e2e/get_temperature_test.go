package e2e_test

import (
	"net/http"
	"testing"
)

const baseURL = "http://localhost:8080/cep/"

func TestValidCEPWithMask(t *testing.T) {
	resp, err := http.Get(baseURL + "99150-000")
	if err != nil {
		t.Fatalf("Failed to make request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status code 200, got %d", resp.StatusCode)
	}
}

func TestValidCEPWithouthMask(t *testing.T) {
	resp, err := http.Get(baseURL + "99150000")
	if err != nil {
		t.Fatalf("Failed to make request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status code 200, got %d", resp.StatusCode)
	}
}

func TestInvalidCEPFormat(t *testing.T) {
	resp, err := http.Get(baseURL + "abcde")
	if err != nil {
		t.Fatalf("Failed to make request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusUnprocessableEntity {
		t.Errorf("Expected status code 422, got %d", resp.StatusCode)
	}
}

func TestNonExistingCEP(t *testing.T) {
	resp, err := http.Get(baseURL + "99999-999")
	if err != nil {
		t.Fatalf("Failed to make request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusNotFound {
		t.Errorf("Expected status code 404, got %d", resp.StatusCode)
	}
}
