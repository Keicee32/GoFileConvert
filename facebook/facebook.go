package facebook

type Facebook struct{
	URL		string
	Name	string
	Friends	int
}

func (f *Facebook) Fame() int{
	return f.Friends
}

func (f *Facebook) Feed() []string{
	return []string{
		"Facebook Feeds",
		"Hey, here's my cool new selfie",
		"What is code?",
	}
}