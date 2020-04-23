package twitter

type Twitter struct{
	URL		string
	Name	string
	Followers	int
}

func (t *Twitter) Feed() []string{
	return []string{
		"Twitter Feeds",
		"Coding is basically copying other people's work",
	}
}

func (t *Twitter) Fame() int{
	//t.Followers = 34
	return t.Followers
}