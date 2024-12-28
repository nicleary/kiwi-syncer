package kiwix_library

import (
	"encoding/xml"
	"github.com/google/uuid"
)

type Feed struct {
	XMLName xml.Name `xml:"feed" json:"feed_name"`
	Entries []Entry  `xml:"entry" json:"entries"`
}

type Entry struct {
	XMLName      xml.Name  `xml:"entry" json:"entry"`
	Id           uuid.UUID `xml:"id" json:"id"`
	Title        string    `xml:"title" json:"title"`
	Updated      string    `xml:"updated" json:"updated"`
	Summary      string    `xml:"summary" json:"summary"`
	Language     string    `xml:"language" json:"language"`
	Name         string    `xml:"name" json:"name"`
	Flavour      string    `xml:"flavour" json:"flavour"`
	Category     string    `xml:"category" json:"category"`
	Tags         string    `xml:"tags" json:"tags"`
	ArticleCount int       `xml:"articleCount" json:"article_count"`
	MediaCount   int       `xml:"mediaCount" json:"media_count"`
	Author       Author    `xml:"author" json:"author"`
	Publisher    Publisher `xml:"publisher" json:"publisher"`
}

type Author struct {
	XMLName xml.Name `xml:"author" json:"author"`
	Name    string   `xml:"name" json:"name"`
}

type Publisher struct {
	XMLName xml.Name `xml:"publisher" json:"publisher"`
	Name    string   `xml:"name" json:"name"`
}