package main

import (
	exporter "exporter"
	"exporter/facebook"
	"exporter/twitter"
	"exporter/linkedin"
)

type SocialMedia interface{
	Feed() []string
	Fame() int
}

func main(){

	txt := (exporter.TextFile)
	son := (exporter.JSONFile)
	xml := (exporter.XMLFile)

	fb := new(facebook.Facebook)
	twtr := new(twitter.Twitter)
	lnkd := new(linkedin.Linkedin)

 	err := txt(twtr, "twtrdata.txt")
	err = txt(fb, "fbdata.txt")
	err = txt(lnkd, "lnkddata.txt")

	if err != nil {
		panic(err)
	}

	b := son(twtr, "twtrdata.json")
	b = son(fb, "fbdata.json")
	b = son(lnkd, "lnkddata.json")

	if b != nil {
		panic(err)
	}

	c := xml(twtr, "twtrdata.xml")
	c = xml(fb, "fbdata.xml")
	c = xml(lnkd, "lnkddata.xml")

	if c != nil {
		panic(err)
	}
}