package main

import (
	"testing"
)


func TestReadDataset(t *testing.T) {
	name := "News_Category_Dataset_v2.json"
	_,err := ReadDataset(name)
	if err != nil {
		t.Errorf("Could not read from %s correctly",name)
	}
}
