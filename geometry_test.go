package facialgeometry

import (
	"math"
	"testing"
)

func TestCanthalTiltDeg(t *testing.T) {
	deg, err := CanthalTiltDeg(Point2D{100, 200}, Point2D{140, 196})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if math.Abs(deg-5.710593) > 1e-4 {
		t.Fatalf("expected ~5.7106, got %f", deg)
	}
}

func TestComputeFacialThirds(t *testing.T) {
	ft, err := ComputeFacialThirds(0, 100, 200, 300)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if math.Abs(ft.BalanceScore-1.0) > 1e-9 {
		t.Fatalf("expected balance 1.0, got %f", ft.BalanceScore)
	}
}

func TestBilateralSymmetry(t *testing.T) {
	pairs := []LandmarkPair{
		{Left: Point2D{80, 100}, Right: Point2D{120, 100}},
		{Left: Point2D{75, 140}, Right: Point2D{125, 140}},
	}
	s, err := BilateralSymmetry(pairs, 100, 120)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if math.Abs(s-1.0) > 1e-9 {
		t.Fatalf("expected symmetry 1.0, got %f", s)
	}
}
