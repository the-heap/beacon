package main

// ByDate implements sort.Interface for []Log based
// on the Date field
type ByDate []Log

func (d ByDate) Len() int           { return len(d) }
func (d ByDate) Swap(i, j int)      { d[i], d[j] = d[j], d[i] }
func (d ByDate) Less(i, j int) bool { return d[i].Date > d[j].Date }
