package main

type EXP struct {
	trees map[string]*tree
}

type tree struct {
	childs []*tree
	method string
}

func (rt *EXP) Group(grp string) {

}
