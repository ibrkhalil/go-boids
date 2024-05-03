package main

import (
	"math"
	"math/rand"
	"time"
)

type Boid struct {
	position Vector2D
	velocity Vector2D
	id int
}

func (b *Boid) calcAcceleration() Vector2D {
	upper, lower := b.position.AddV(viewRadius), b.position.AddV(-viewRadius)
	averageVelcoity := Vector2D{x: 0, y: 0}
	count := 0.0

	for i := math.Max(lower.x, 0); i <= math.Min(upper.x, screenWidth); i++ {
		for j := math.Max(lower.y, 0); j <= math.Min(upper.y, screenHeight); j++ {
			if otherBoidID := boidMap[int(i)][int(j)]; otherBoidID != -1 && otherBoidID != b.id {
				if distance := boids[otherBoidID].position.Distance(b.position); distance < viewRadius {
					count++
					averageVelcoity = averageVelcoity.Add(boids[otherBoidID].velocity)
				}
			}
		}
	}



	accel := Vector2D{x: 0, y: 0}
	if count > 0 {
		averageVelcoity = averageVelcoity.DivisionV(count)
		accel = averageVelcoity.Subtract(b.velocity).MultiplyV(adjRate)
	}
	return accel 
}

func (b *Boid) moveOne()  {
	b.velocity = b.velocity.Add(b.calcAcceleration()).limit(-1, 1)
	boidMap[int(b.position.x)][int(b.position.y)] = -1
	// Bounce back if going out of view
	next := b.position.Add(b.velocity)
	if next.x >= screenWidth || next.x < 0 {
		b.velocity = Vector2D{-b.velocity.x, b.velocity.y}
	}

	if next.y >= screenHeight || next.y < 0 {
		b.velocity = Vector2D{b.velocity.y, -b.velocity.y}
	}

	b.position = b.position.Add(b.velocity)
	boidMap[int(b.position.x)][int(b.position.y)] = b.id
}

func (b *Boid) start()  {
	for {
		b.moveOne()
		time.Sleep(5 * time.Millisecond)
	}
}

func createBoid(bid int)  {
	b := Boid{
		position: Vector2D{x: rand.Float64() * screenWidth, y: rand.Float64() * screenHeight},
		velocity: Vector2D{x: (rand.Float64() * 2) - 1.0, y: (rand.Float64() * 2) - 1.0},
		id: bid,
	}
	boids[bid] = &b

	boidMap[int(b.position.x)][int(b.position.y)] = bid
	go b.start()
}
