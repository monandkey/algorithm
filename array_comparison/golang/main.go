package main

func search(source, target []string) (store []string) {
	for _, v := range source {
		if str := searchDiff(v, target); str != "" {
			store = append(store, str)
		}
	}
	return store
}

func searchDiff(v string, target []string) string {
	for _, w := range target {
		if v == w {
			return w
		}
	}
	return ""
}
