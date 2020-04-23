package linkedin

type Linkedin struct{
	URL		string
	Name	string
	Connections	int
}

func (l *Linkedin) Feed() []string{
	return []string{
		"Linkedin Feeds",
		"Hey, I just started a new position at Hotels.ng",
	}
}

func (l *Linkedin) Fame() int{
	return l.Connections
}