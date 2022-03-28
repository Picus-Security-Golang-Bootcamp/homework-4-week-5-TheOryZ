package model

//Books model
type Books struct {
	Books []Book `json:"Books"`
}

//Book detail model
type Book struct {
	ID             int     `json:"Id"`
	Title          string  `json:"Title"`
	NumberOfPages  int     `json:"Number_Of_Pages"`
	NumberOfStocks int     `json:"Number_Of_Stocks"`
	Price          float64 `json:"Price"`
	ISBN           string  `json:"ISBN"`
	ReleaseDate    string  `json:"Release_Date"`
	AuthorID       int     `json:"Author_Id"`
}
