// Package facialgeometry provides pure-math 2D facial landmark geometry helpers
// for canthal tilt, facial thirds proportions, golden ratio delta, and bilateral symmetry.
//
// Coordinates follow top-down image space (X right, Y down).
// Reference formulas and interactive calculator: https://looksmaxscale.com/
package facialgeometry

import (
	"errors"
	"math"
)

// Phi is the golden ratio constant (1 + sqrt(5)) / 2.
const Phi = 1.618033988749895

// Point2D represents a 2D facial landmark coordinate.
type Point2D struct {
	X float64
	Y float64
}

// LandmarkPair holds corresponding left and right facial landmarks.
type LandmarkPair struct {
	Left  Point2D
	Right Point2D
}

// FacialThirds holds normalized upper, middle, and lower facial ratios plus a balance score.
type FacialThirds struct {
	Upper        float64
	Middle       float64
	Lower        float64
	BalanceScore float64
}

// CanthalTiltDeg computes the medial-to-lateral canthal tilt angle in degrees.
// Positive degrees indicate an upward lateral tilt in top-down image coordinates.
func CanthalTiltDeg(medial, lateral Point2D) (float64, error) {
	dx := math.Abs(lateral.X - medial.X)
	if dx <= 1e-12 {
		return 0, errors.New("medial and lateral x coordinates must differ")
	}
	dyUp := medial.Y - lateral.Y
	return math.Atan2(dyUp, dx) * (180.0 / math.Pi), nil
}

// ComputeFacialThirds computes vertical facial thirds from strictly increasing top-down Y values.
func ComputeFacialThirds(trichionY, glabellaY, subnasaleY, mentonY float64) (FacialThirds, error) {
	if !(trichionY < glabellaY && glabellaY < subnasaleY && subnasaleY < mentonY) {
		return FacialThirds{}, errors.New("y coordinates must strictly increase from trichion to menton")
	}
	total := mentonY - trichionY
	u := (glabellaY - trichionY) / total
	m := (subnasaleY - glabellaY) / total
	l := (mentonY - subnasaleY) / total
	const target = 1.0 / 3.0
	l1 := math.Abs(u-target) + math.Abs(m-target) + math.Abs(l-target)
	score := math.Max(0.0, math.Min(1.0, 1.0-(l1/(4.0/3.0))))
	return FacialThirds{Upper: u, Middle: m, Lower: l, BalanceScore: score}, nil
}

// GoldenRatioDelta returns height/width and its signed deviation from Phi.
func GoldenRatioDelta(faceHeight, faceWidth float64) (ratio float64, delta float64, err error) {
	if faceHeight <= 0 || faceWidth <= 0 {
		return 0, 0, errors.New("dimensions must be positive")
	}
	ratio = faceHeight / faceWidth
	return ratio, ratio - Phi, nil
}

// BilateralSymmetry returns a normalized mirror symmetry score in [0, 1].
func BilateralSymmetry(pairs []LandmarkPair, midlineX, faceWidth float64) (float64, error) {
	if len(pairs) == 0 || faceWidth <= 0 {
		return 0, errors.New("pairs must be non-empty and faceWidth must be positive")
	}
	var sumErr float64
	for _, p := range pairs {
		mirroredLX := 2.0*midlineX - p.Left.X
		dx := mirroredLX - p.Right.X
		dy := p.Left.Y - p.Right.Y
		sumErr += math.Hypot(dx, dy)
	}
	meanErr := sumErr / float64(len(pairs))
	score := 1.0 - (meanErr / faceWidth)
	return math.Max(0.0, math.Min(1.0, score)), nil
}
