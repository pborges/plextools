package plextools

import (
	"net/http"
	"encoding/xml"
	"strings"
)

func sanitize(s string) string {
	s = strings.Replace(s, "/", "-", -1)
	s = strings.Replace(s, " & ", " and ", -1)
	s = strings.Replace(s, "&", "and", -1)
	s = strings.Replace(s, ".", "", -1)
	s = strings.Replace(s, "'", "", -1)
	s = strings.Replace(s, "!", "", -1)
	s = strings.Replace(s, ";", "", -1)
	s = strings.Replace(s, ":", "", -1)
	s = strings.Replace(s, ",", "", -1)
	s = strings.Replace(s, " ", ".", -1)
	return s
}

type NamedEntry interface {
	FormattedFileName() string
	FilePath() string
}

func Dial(addr string) Client {
	return Client{
		addr: addr,
	}
}

type Client struct {
	addr string
}

func (p Client) request(url string) (mediaContainer, error) {
	//fmt.Printf("request: %s\n", url)
	m := mediaContainer{}

	res, err := http.Get(url)
	if err != nil {
		return m, err
	}

	err = xml.NewDecoder(res.Body).Decode(&m)
	if err != nil {
		return m, err
	}
	return m, nil
}

type mediaContainer struct {
	Identifier  string      `xml:"identifier,attr"`
	Directories []directory `xml:"Directory"`
	Videos      []video     `xml:"Video"`
}

type video struct {
	Key              string `xml:"key,attr"`
	RatingKey        string `xml:"ratingKey,attr"`
	Title            string `xml:"title,attr"`
	ParentTitle      string `xml:"parentTitle,attr"`
	GrandParentTitle string `xml:"grandparentTitle,attr"`
	Index            int    `xml:"index,attr"`
	ParentIndex      int    `xml:"parentIndex,attr"`
	Media            []media  `xml:"Media"`
}

type media struct {
	Part struct {
		File string `xml:"file,attr"`
	} `xml:"Part"`
}

type directory struct {
	Key       string   `xml:"key,attr"`
	RatingKey string   `xml:"ratingKey,attr"`
	Title     string   `xml:"title,attr"`
	Type      string   `xml:"type,attr"`
	Location  location `xml:"Location"`
}

type location struct {
	Path string `xml:"path,attr"`
}
