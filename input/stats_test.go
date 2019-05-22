package input

import (
  "testing"
  "os"
)

var nClients, nCalls = 1024, 128
var veryPopular = Parameters{3, 5, 20, "buzz", "fizz"}
var lessPopular = Parameters{5, 7, 40, "Hello", "World"}
var isItPopular = Parameters{6, 7, 50, "AnswerTo", "TheQuestion"}

func client(n int, done chan int) {
  for i := 0; i < nCalls; i++ {
    switch (i + n) % 4 {
    case 0, 2:
      RegisterHit(veryPopular)
    case 1:
      RegisterHit(lessPopular)
    case 3:
      RegisterHit(isItPopular)
    }
  }
  done <- 0
}

func TestMain(m *testing.M) {
  // Concurretly registering lots of stats...
  done := make(chan int, nClients)
  for i := 0; i < nClients; i++ {
    go client(1, done)
  }
  // Waiting for all to be registered...
  for n := nClients; n > 0; n-- {
    <- done
  }
  // Testing results
  os.Exit(m.Run())
}

func TestMostPopular(t *testing.T) {
  expected := HitStats{veryPopular, nClients * nCalls / 2}
  if got := MostPopular(); got != expected {
    t.Errorf("Does not work well under pressure: got %+v; want %+v.", got, expected)
  }
}

func TestRegisterHit(t *testing.T) {
  var newParams Parameters
  RegisterHit(newParams)
  if stats[newParams] != 1 {
    t.Error("Failed registering new parameters.")
  }
}
