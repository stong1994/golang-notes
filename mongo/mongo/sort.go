package mongo

type SortCond interface {
	And(...SortCond) SortCond
	IsValid() bool
}

type SingleSortCondition struct {
	Key string
	Desc bool
}

func (s SingleSortCondition) And(conds ...SortCond) SortCond {
	return and(s, and(conds...))
}

func (s SingleSortCondition) IsValid() bool {
	return s.Key != ""
}

func CondExpr(key string, desc bool) SortCond {
	return SingleSortCondition{
		Key:  key,
		Desc: desc,
	}
}

type emptySortConditions struct {}

func NewEmptySortCond() SortCond {
	return emptySortConditions{}
}

func (e emptySortConditions) And(conds ...SortCond) SortCond {
	return and(conds...)
}

func (e emptySortConditions) IsValid() bool {
	return false
}

type sortConditions []SortCond

var _ SortCond = sortConditions{}

func (c sortConditions) And(conds ...SortCond) SortCond {
	return and(c, and(conds...))
}

func (c sortConditions) IsValid() bool {
	return len(c) > 0
}

func and(conds ...SortCond) SortCond{
	ss := make(sortConditions, 0, len(conds))
	for _, cond := range conds {
		if cond == nil || !cond.IsValid() {
			continue
		}
		ss = append(ss, cond)
	}
	return ss
}