package main

import (
	"testing"
)

func TestPerimeter(t *testing.T) {
	got := Perimeter(&Rectangle{
		Width:  10.0,
		Height: 10.0,
	})
	want := 40.0

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

func TestArea(t *testing.T) {
	r := &Rectangle{
		Width:  12.0,
		Height: 6.0,
	}
	got := r.Area()
	want := 72.0

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

func TestArea2(t *testing.T) {

	checkArea := func(t *testing.T, shape Shape, want float64) {
		t.Helper()
		got := shape.Area()
		if got != want {
			t.Errorf("got %g want %g", got, want)
		}
	}

	t.Run("rectangles", func(t *testing.T) {
		rectangle := &Rectangle{12, 6}
		checkArea(t, rectangle, 72.0)
	})

	t.Run("circles", func(t *testing.T) {
		circle := &Circle{10}
		checkArea(t, circle, 314.1592653589793)
	})
}

func TestArea3(t *testing.T) {

	// テーブル駆動テストは、同じ方法でテストできるテストケースのリストを作成する
	areaTests := []struct {
		name    string
		shape   Shape
		hasArea float64
	}{
		{name: "Rectangle", shape: Rectangle{Width: 12, Height: 6}, hasArea: 72.0},
		{name: "Circle", shape: Circle{Radius: 10}, hasArea: 314.1592653589793},
	}

	for _, tt := range areaTests {
		// using tt.name from the case to use it as the `t.Run` test name
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Area()
			if got != tt.hasArea {
				t.Errorf("%#v got %g want %g", tt.shape, got, tt.hasArea)
			}
		})

	}

}
