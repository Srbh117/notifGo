package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"testing"
)

func TestStorage(t *testing.T) {
	s := NewMemoryStorage()

	for i := range 100 {
		num := fmt.Sprintf("%s", strconv.Itoa(rand.Intn(i*100+1)))
		latestOffset, err := s.Push([]byte(num))
		if err != nil {
			t.Log(err.Error())
		}
		_, err = s.Fetch(latestOffset)
		if err != nil {
			t.Log(err)
		}
		// log.Default(data)
	}
	fmt.Println(s)

}
