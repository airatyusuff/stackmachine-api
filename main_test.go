package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRootHandler(t *testing.T) {
	//Arrange
	req, err := http.NewRequest("GET", "/", nil)

	if err != nil {
		t.Fatal(err)
	}

	testResponse := httptest.NewRecorder()
	handler := http.HandlerFunc(rootHandler)

	// Act
	handler.ServeHTTP(testResponse, req)

	//Assert
	if status := testResponse.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	expected := "Stack machine API"
	if testResponse.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", testResponse.Body.String(), expected)
	}
}

func TestExecuteHandler(t *testing.T) {
	//Arrange
	requestBody, err := json.Marshal(Command{Text: "35 40 +"})
	req, err := http.NewRequest("POST", "/execute", bytes.NewBuffer(requestBody))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Add("Content-Type", "application/json")

	testResponse := httptest.NewRecorder()
	handler := http.HandlerFunc(executeCommand)

	// Act
	handler.ServeHTTP(testResponse, req)

	//Assert
	if status := testResponse.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	var result Result
	json.NewDecoder(testResponse.Body).Decode(&result)

	expected := 75
	if result.Data != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", result.Data, expected)
	}
}
