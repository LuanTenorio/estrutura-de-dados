package greedy

type Activity struct{ Start, End int }

func ActivitySelection(activities []Activity) []Activity {
	n := len(activities)

	setActitives := make([]Activity, n/2)
	k := 1

	for m := 2; m < n; m++ {
		if activities[m].Start >= activities[k].End {
			setActitives = append(setActitives, activities[m])
			k = m
		}
	}

	return setActitives
}
