# go-randomuser

A GO library to get a random user information from http://randomuser.me/

## Usage

```
import "github.com/petehouston/go-randomuser"
```

```go
c := &randomuser.QueryConfig{
    5, "", "", "",
}
users, err := randomuser.Generate(c)

if err != nil {
    fmt.Println(err)
    return
}

for _, user := range users {
    fmt.Println(user.Name.Title + " : " + user.Name.First + " " + user.Name.Last)
}
```

## Data structure

### QueryConfig

```go
type QueryConfig struct {
	MaxResults int
	Gender     string
	Password   string
	Seed       string
}
```

* **MaxResults** : specify the number of users to query.
* **Gender** : specify gender, can be `male`, `female`. By default, service will return both.
* **Password** : specify the password rule for user account. Read the [randomuser.me docs - Passwords](https://randomuser.me/documentation#passwords)
* **Seed** : to generate a same set of users.

Please refer to [randomuser.me - docs](https://randomuser.me/documentation) to get more information.


### User

```go
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
```

## License

MIT.