package main

func earliestFinishTime(landStartTime []int, landDuration []int, waterStartTime []int, waterDuration []int) int {
	minTime := 200000
	
	for i := 0; i < len(landStartTime); i++  {

		for j := 0; j < len(waterStartTime); j++ {
			landEnd := landStartTime[i] + landDuration[i]
			waterStart := max(landEnd, waterStartTime[j])
			count1 := waterStart + waterDuration[j]

			waterEnd := waterStartTime[j] + waterDuration[j]
			landStart := max(waterEnd, landStartTime[i])
            count2 := landStart + landDuration[i]
            
			if count1 < minTime {
				minTime = count1
			}
			if count2 < minTime {
				minTime = count2
			}
            minTime = min(minTime, min(count1, count2))
		}
	}

	return minTime
}