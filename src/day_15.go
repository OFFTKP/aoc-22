package main

import (
	"fmt"
	"math"
	"sort"
	"time"
)

type Sensor struct {
	x, y         int
	bx, by       int
	topY, botY   int
	topBotCached bool
}

type Range struct {
	minx, maxx int
}

var sensors []Sensor
var border int = 4000000
var quit chan bool
var startTime time.Time

func main() {
	parseSensors()
	part1()
	startTime = time.Now()
	for y := 0; y <= border; y++ {
		go part2(y)
	}
}

func parseSensors() {
	lines := GetLines(15)
	gridminx := math.MaxInt32
	gridmaxx := math.MinInt32
	for _, line := range lines {
		var sensor Sensor
		fmt.Sscanf(line, "Sensor at x=%d, y=%d: closest beacon is at x=%d, y=%d", &sensor.x, &sensor.y, &sensor.bx, &sensor.by)
		if sensor.bx < gridminx {
			gridminx = sensor.bx
		}
		if sensor.bx > gridmaxx {
			gridmaxx = sensor.bx
		}
		sensors = append(sensors, sensor)
	}
}

func part2(y int) {
	select {
	case <-quit:
		return
	default:
	}
	cur := getRanges(y)
	sort.Slice(cur, func(i, j int) bool {
		if cur[i].minx == cur[j].minx {
			return cur[i].maxx < cur[j].maxx
		}
		return cur[i].minx < cur[j].minx
	})
	var merged []Range
	var current *Range = &cur[0]
	for i := 1; i < len(cur); i++ {
		if current.maxx >= cur[i].minx {
			current.maxx = MaxOf(current.maxx, cur[i].maxx)
		} else {
			merged = append(merged, *current)
			current = &cur[i]
		}
	}
	if len(merged) == 0 {
		merged = append(merged, *current)
	}
	for _, r := range merged {
		if r.minx > 0 {
			x := r.minx - 1
			fmt.Printf("x: %d, y: %d, part 2: %d\n", x, y, (x*4000000)+y)
			fmt.Printf("time: %v\n", time.Since(startTime))
			quit <- true
		} else if r.maxx < border {
			x := r.maxx - 1
			fmt.Printf("x: %d, y: %d, part 2: %d\n", x, y, (x*4000000)+y)
			fmt.Printf("time: %v\n", time.Since(startTime))
			quit <- true
		}
	}
}

func part1() {
	part1Ranges := getRanges(2000000)
	rminx := math.MaxInt32
	rmaxx := math.MinInt32
	for _, r := range part1Ranges {
		if r.minx < rminx {
			rminx = r.minx
		}
		if r.maxx > rmaxx {
			rmaxx = r.maxx
		}
	}
	sum := 0
	for i := rminx; i <= rmaxx; i++ {
		for _, r := range part1Ranges {
			if i >= r.minx && i <= r.maxx {
				dontadd := false
				for _, s := range sensors {
					if s.bx == i && s.by == 2000000 {
						dontadd = true
						break
					}
				}
				if !dontadd {
					sum++
				}
				break
			}
		}
	}
	fmt.Printf("part 1: %d\n", sum)
}

func getRanges(yline int) []Range {
	var ranges []Range
	for _, sensor := range sensors {
		topY, botY := findTopBotY(&sensor)
		if yline >= topY && yline <= botY {
			var n int
			if yline >= sensor.y {
				n = sensor.y - topY - (yline - sensor.y)
			} else {
				n = yline - topY
			}
			minx := sensor.x - n
			maxx := sensor.x + n
			ranges = append(ranges, Range{minx, maxx})
		}
	}
	return ranges
}

func findTopBotY(sensor *Sensor) (int, int) {
	if sensor.topBotCached {
		return sensor.topY, sensor.botY
	}
	by := sensor.by + sign(sensor.by-sensor.y)*abs(sensor.bx-sensor.x)
	temp := abs(by - sensor.y)
	sensor.topY = sensor.y - temp
	sensor.botY = sensor.y + temp
	sensor.topBotCached = true
	return sensor.topY, sensor.botY
}
