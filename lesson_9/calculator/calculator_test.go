package main

import (
	"log"
	"testing"
)

func TestAdd(t *testing.T) {
	//1. 1-ый test-case
	if res := add(20, 10); res != 30 {
		log.Fatalf("Add(20, 10) fail. expected %d, got %d\n", 30, res) // log.Fatal провоцирует завершение работы кода
	}
}
func TestSub(t *testing.T) {
	//1. 1-ый test-case
	if res := Sub(20, 10); res != 10 {
		log.Fatalf("Sub(20, 10) fail. expected %d, got %d\n", 10, res) // log.Fatal провоцирует завершение работы кода
	}
}

func TestMult(t *testing.T) {
	//1. 1-ый test-case
	if res := Mult(20, 10); res != 200 {
		log.Fatalf("Mult(20, 10) fail. expected %d, got %d\n", 200, res) // log.Fatal провоцирует завершение работы кода
	}
}

func TestDiv(t *testing.T) {
	//1. 1-ый test-case
	if res := Div(20, 10); res != 2 {
		log.Fatalf("Div(20, 10) fail. expected %d, got %d\n", 2, res) // log.Fatal провоцирует завершение работы кода
	}
}
