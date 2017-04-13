package randomuser

type JsonRandomUser struct {
	Results []User `json:"results"`
	Info    Info   `json:"info"`
}

type User struct {
	Gender     string   `json:"gender"`
	Name       Name     `json:"name"`
	Location   Location `json:"location"`
	Email      string   `json:"email"`
	Login      Login    `json:"login"`
	Dob        string   `json:"dob"`
	Registered string   `json:"registered"`
	Phone      string   `json:"phone"`
	Cell       string   `json:"cell"`
	Id         Id       `json:"id"`
	Picture    Picture  `json:"picture"`
	Nat        string   `json:"nat"`
}

type Name struct {
	Title string `json:"title"`
	First string `json:"first"`
	Last  string `json:"last"`
}

type Location struct {
	Street   string      `json:"street"`
	City     string      `json:"city"`
	State    string      `json:"state"`
	Postcode interface{} `json:"postcode"`
}

type Login struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Salt     string `json:"salt"`
	Md5      string `json:"md5"`
	Sha1     string `json:"sha1"`
	Sha256   string `json:"sha256"`
}

type Id struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type Picture struct {
	Large     string `json:"large"`
	Medium    string `json:"medium"`
	Thumbnail string `json:"thumbnail"`
}

type Info struct {
	Seed    string `json:"seed"`
	Results int    `json:"results"`
	Page    int    `json:"page"`
	Version string `json:"version"`
}
