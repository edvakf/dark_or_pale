package main

import "testing"

func TestDarkness(t *testing.T) {
	img := openImageFile("./images/rgb.png")
	actual := Darkness(img)
	expected := 70
	if actual != expected {
		t.Fatalf("Expected darkness %d, actual %d", expected, actual)
	}
}

func TestDarkness2(t *testing.T) {
	img := openImageFile("./images/gray.png")
	actual := Darkness(img)
	expected := 70
	if actual != expected {
		t.Fatalf("Expected darkness %d, actual %d", expected, actual)
	}
}
