package decision

type Quality struct {
	Id          string
	Name        string
	Importance  float64
	Probability float64
}

type Option struct {
	Id   string
	Name string
	Pros []*Quality
	Cons []*Quality
}

type Problem struct {
	Id      string
	Name    string
	Options []*Option
}

type Result struct {
	OptionsRating map[string]float64
}

type Decision struct {
	Id        string
	ProblemId string
	UserId    string
	Result    Result
}
