package model

import (
	"sort"
	"time"
)

type Task struct {
	Id          int       `json:"id"`
	Description string    `json:"description"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

type TaskSlice []Task
type TaskMap map[int]Task

func (s TaskSlice) Len() int           { return len(s) }
func (s TaskSlice) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s TaskSlice) Less(i, j int) bool { return s[i].Id < s[j].Id }

func (m TaskMap) OrderById() TaskSlice {
	ids := make([]int, 0, len(m))

	for id := range m {
		ids = append(ids, id)
	}

	sort.Ints(ids)

	var result []Task

	for _, id := range ids {
		result = append(result, m[id])
	}

	return result
}
