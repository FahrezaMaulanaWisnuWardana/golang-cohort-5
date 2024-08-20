package helpers

type Peoples struct {
	ID       int
	Email    string
	Password string
	Address  string
	Job      string
	Reason   string
}

var People = []Peoples{
	{ID: 1, Email: "john.doe@example.com", Password: "123", Address: "Love For Imperfect Things", Job: "Software Engineer", Reason: "Haemin Sunim"},
	{ID: 2, Email: "jane.smith@example.com", Password: "123", Address: "The Power of Nunchi", Job: "Software Engineer", Reason: "Euny Hong"},
	{ID: 3, Email: "michale.jackson@example.com", Password: "123", Address: "Winter in Tokyo", Job: "Software Engineer", Reason: "Ilana Tan"},
	{ID: 4, Email: "sarah.wilson@example.com", Password: "123", Address: "Winter in Tokyo", Job: "Software Engineer", Reason: "Ilana Tan"},
	{ID: 5, Email: "robert.jackson@example.com", Password: "123", Address: "Winter in Tokyo", Job: "Software Engineer", Reason: "Ilana Tan"},
}
