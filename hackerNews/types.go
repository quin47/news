package hackerNews

type article struct {
	Id int32
	Delete bool
	Type string `json:"type"`
	By string
	Time  int64
	Text string
	Dead bool
	Parent int32
	Poll string
	Kids []int32
	Url string
	Score int32
	Title string
	Parts []int32
	descendants int32
}
