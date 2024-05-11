package main

import (
	"fmt"
	"math"
	"regexp"
	"strconv"
)

type Point2D struct {
	x int
	y int
}

var (r = regexp.MustCompile(`\((\d*),(\d*)\)`))

func findArea(pointStr string)  {
	var points []Point2D
	for _, p := range r.FindAllStringSubmatch(pointStr, -1) {
		x, _ := strconv.Atoi(p[1])
		y, _ := strconv.Atoi(p[2])
		points = append(points, Point2D{x: x, y: y})
	}

	area := 0.0
	for i := 0; i < len(points); i++ {
		a, b := points[i], points[(i+1) % len(points)]
		area += float64(a.x * b.y) - float64(a.y * b.x)
	}
	fmt.Println(math.Abs(area) / 2)
}

func main() {
	line := "(4,10),(12,8),(10,3),(2,2),(7,5)"
	findArea(line)
}
