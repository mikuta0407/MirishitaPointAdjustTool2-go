package calc

import (
	"fmt"
	"math"
	"os"

	"github.com/mikuta0407/mltd-pat/config"
)

var pointLimit = 10000000

func Main(pointlist config.PointList, mode string, currentPoint int, targetPoint int) error {
	fmt.Println(pointlist)

	if currentPoint < 0 || targetPoint < 0 || targetPoint > pointLimit {
		return fmt.Errorf("input value between 0 and %d", pointLimit)
	}

	var err error
	if mode == "aniv" {
		err = anivCalc(pointlist, currentPoint, targetPoint)
	}
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return nil
}

func anivCalc(pointlist config.PointList, currentPoint int, targetPoint int) error {
	remain := targetPoint - currentPoint

	minLiveCount := make([]int, remain+1, remain+1)
	for i := 1; i <= remain; i++ {
		minLiveCount[i] = math.MaxInt64
	}

	for i := 1; i <= remain; i++ {
		for j := 1; j <= len(pointlist.ANNIVERSARY.LIVE); j++ {
			live := pointlist.ANNIVERSARY.LIVE[j]
			if i+live.POINT > remain {
				continue
			}

			if j >= pointlist.ANNIVERSARY.LIVE[i].POINT {
				//				minLiveCount[j] = min
			}
		}
	}

	lastLiveIndex := make([]int, remain+1, remain+1)
	for i := 1; i <= remain; i++ {
		minLiveCount[i] = -1
	}

	for i := 0; i <= remain; i++ {
		for j := 0; j < len(pointlist.ANNIVERSARY.LIVE); j++ {
			live := pointlist.ANNIVERSARY.LIVE[j]
			if i+live.POINT > remain {
				continue
			}

			if minLiveCount[i]+1 < minLiveCount[i+live.POINT] {
				minLiveCount[i+live.POINT] = minLiveCount[i] + 1
				lastLiveIndex[i+live.POINT] = j
			}
		}
	}

	liveCount := make([]int, len(pointlist.ANNIVERSARY.LIVE))

	point := remain

	for {
		if point == 0 {
			break
		}

		lastLive := lastLiveIndex[point]
		liveCount[lastLive]++
		point -= pointlist.ANNIVERSARY.LIVE[lastLive].POINT
	}

	fmt.Println("live count is below.")
	for i := 0; i < len(pointlist.ANNIVERSARY.LIVE); i++ {
		fmt.Println(pointlist.ANNIVERSARY.LIVE[i].NAME, liveCount[i])
	}

	return nil
}
