package parse

type stringArraySet []string

func (set stringArraySet) append(val string) stringArraySet {
	if set.contains(val) {
		return set
	}
	return append(set, val)
}

func (set stringArraySet) contains(val string) bool {
	for _, s := range set {
		if val == s {
			return true
		}
	}
	return false
}
