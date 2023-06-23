package model

/*
	"asset13", "yellow", "5", "Tom", "1300"
*/
type Asset struct {
	ID             string
	Color          string
	Size           string // Int
	Owner          string
	AppraisedValue string // Int
}

type Data struct {
	Id          string
	TenderID    string
	Accountcode string
	Account     string
	Name        string
	Currency    string
	Branch      string
	Amount      string
	Status      string
}
