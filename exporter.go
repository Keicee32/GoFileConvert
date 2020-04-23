package exporter

import(
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	"os"
)

type SocialMedia interface{
	Feed() []string
}

func TextFile(sm SocialMedia, filename string) error{
	e, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY,0755)
	defer e.Close()

	if err != nil{
		return errors.New("There was an error opening this file: " + err.Error())
	}
	for _, kc := range sm.Feed() {
		f, err := e.Write([]byte(kc +"\n"))
		if err != nil {
			return errors.New("There was an error opening this file: " + err.Error())
		}
		fmt.Printf("wrote %d bytes\n", f)
	}
	return nil
}

//JSON FILE
func JSONFile(sm SocialMedia, filename string) error {
	e, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY, 0755)
	defer e.Close()

	if err != nil {
		return errors.New("There was an error opening this file: " + err.Error())
	}
	for _, kc := range sm.Feed() {
		g, err := json.MarshalIndent(kc, "\n", "\n")
		if err != nil {
			return errors.New("There was an error opening this file: " + err.Error())
		}
		
		byteWritten, err := e.Write(g)

		if err != nil {
			return errors.New("There was an error opening this file: " + err.Error())
		}
		
		fmt.Printf("wrote %d bytes\n", byteWritten)

	}
	return nil
}


//XML FILE
func XMLFile(sm SocialMedia, filename string) error {
	e, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY, 0755)
	defer e.Close()

	if err != nil {
		return errors.New("There was an error opening this file: " + err.Error())
	}
	for _, ke := range sm.Feed() {
		h, err := xml.MarshalIndent(ke, "\n", "\n")
		if err != nil {
			return errors.New("There was an error opening this file: " + err.Error())
		}
		
		byteWritten, err := e.Write(h)

		if err != nil {
			return errors.New("There was an error opening this file: " + err.Error())
		}
		
		fmt.Printf("wrote %d bytes\n", byteWritten)

	}
	return nil
}