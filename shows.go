package plextools

import (
	"fmt"
	"strconv"
)

type Show struct {
	Title    string
	Episodes []Episode
}

type Episode struct {
	Show    string
	Season  int
	Episode int
	Title   string
	File    string
}

func (v Episode) FormattedFileName() string {
	show := sanitize(v.Show)
	title := sanitize(v.Title)
	return fmt.Sprintf("%s.S%02dE%02d.%s", show, v.Season, v.Episode, title)
}

func (v Episode) FilePath() string {
	return v.File
}

func (p Client) Shows() ([]Show, error) {
	sections, err := p.Sections()
	if err != nil {
		return nil, err
	}
	shows := make([]Show, 0)
	for _, section := range sections {
		if section.Type == "show" {
			sectionMediaContainer, err := p.request(fmt.Sprintf("http://%s/library/sections/%d/all", p.addr, section.Id))
			if err != nil {
				return nil, err
			}

			for _, dir := range sectionMediaContainer.Directories {
				id, err := strconv.Atoi(dir.RatingKey)
				if err != nil {
					return nil, err
				}
				epMediaContainer, err := p.request(fmt.Sprintf("http://%s/library/metadata/%v/allLeaves", p.addr, id))
				if err != nil {
					return nil, err
				}
				episodes := make([]Episode, len(epMediaContainer.Videos))
				for i, v := range epMediaContainer.Videos {
					for _, m := range v.Media {
						episodes[i] = Episode{
							Show:    dir.Title,
							Title:   v.Title,
							File:    m.Part.File,
							Season:  v.ParentIndex,
							Episode: v.Index,
						}
					}
				}
				shows = append(shows, Show{
					Title:    dir.Title,
					Episodes: episodes,
				})
			}
		}
	}
	return shows, nil
}
