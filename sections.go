package plextools

import (
	"fmt"
	"strconv"
)

type Section struct {
	Id    int
	Title string
	Path  string
	Type  string
}

type Entry struct {
	Id    int
	Title string
}

func (p Client) Sections() ([]Section, error) {
	m, err := p.request(fmt.Sprintf("http://%s/library/sections", p.addr))
	if err != nil {
		return nil, err
	}
	s := make([]Section, len(m.Directories))
	for i, d := range m.Directories {
		id, err := strconv.Atoi(d.Key)
		if err != nil {
			return nil, err
		}
		section := Section{
			Id:    id,
			Title: d.Title,
			Type:  d.Type,
			Path:  d.Location.Path,
		}
		s[i] = section
	}
	return s, nil
}