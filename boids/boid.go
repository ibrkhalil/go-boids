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
	averagePosition, averageVelcoity, separation := Vector2D{x: 0, y: 0}, Vector2D{x: 0, y: 0}, Vector2D{x: 0, y: 0}
	count := 0.0
	rWlock.RLock()
	for i := math.Max(lower.x, 0); i <= math.Min(upper.x, screenWidth); i++ {
		for j := math.Max(lower.y, 0); j <= math.Min(upper.y, screenHeight); j++ {
			if otherBoidID := boidMap[int(i)][int(j)]; otherBoidID != -1 && otherBoidID != b.id {
				if distance := boids[otherBoidID].position.Distance(b.position); distance < viewRadius {
					count++
					averageVelcoity = averageVelcoity.Add(boids[otherBoidID].velocity)
					averagePosition = averagePosition.Add(boids[otherBoidID].position)
					separation = separation.Add(b.position.Subtract(boids[otherBoidID].position).DivisionV(distance))
				}
			}
		}
	}
	rWlock.RUnlock()

	accel := Vector2D{x: b.borderBounce(b.position.x, screenWidth), y: b.borderBounce(b.position.y, screenHeight)}
	if count > 0 {
		averagePosition, averageVelcoity = averagePosition.DivisionV(count), averageVelcoity.DivisionV(count)
		accelAlignment := averageVelcoity.Subtract(b.velocity).MultiplyV(adjRate)
		accelCohesion := averagePosition.Subtract(b.position).MultiplyV(adjRate)
		accelSeparation := separation.MultiplyV(adjRate)
		accel = accel.Add(accelAlignment).Add(accelCohesion).Add(accelSeparation)
	}
	return accel 
}

func (b *Boid) borderBounce(pos, maxBorderPos float64) float64 {
	if pos < viewRadius {
		return 1 / pos
	} else if pos > maxBorderPos - viewRadius {
		return 1 / (pos - maxBorderPos)
	}
	return 0
}

func (b *Boid) moveOne()  {
	acceleration := b.calcAcceleration()
	rWlock.Lock()
	b.velocity = b.velocity.Add(acceleration).limit(-1, 1)
	boidMap[int(b.position.x)][int(b.position.y)] = -1
	

	b.position = b.position.Add(b.velocity)
	boidMap[int(b.position.x)][int(b.position.y)] = b.id
	rWlock.Unlock()
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
