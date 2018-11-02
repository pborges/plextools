package plextools

import "fmt"

type Movie struct {
	Title string
	File  string
}

func (m Movie) FormattedFileName() string {
	return sanitize(m.Title)
}

func (m Movie) FilePath() string {
	return m.File
}

func (p Client) Movies() ([]Movie, error) {
	sections, err := p.Sections()
	if err != nil {
		return nil, err
	}
	movies := make([]Movie, 0)
	for _, section := range sections {
		if section.Type == "movie" {
			sectionMediaContainer, err := p.request(fmt.Sprintf("http://%s/library/sections/%d/all", p.addr, section.Id))
			if err != nil {
				return nil, err
			}
			for _, v := range sectionMediaContainer.Videos {
				for _, m := range v.Media {
					movies = append(movies, Movie{
						Title: v.Title,
						File:  m.Part.File,
					})
				}
			}
		}
	}
	return movies, nil
}
