package main

func search(source, target []string) (store []string) {
	for _, v := range source {
		store = searchDiff(v, target, store)
	}
	return store
}

func searchDiff(v string, target []string, store []string) []string {
	for _, w := range target {
		if v == w {
			store = append(store, v)
			return store
		}
	}
	return store
}
