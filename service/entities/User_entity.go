package entities

// user entity
// one user map to one user entity

// ID -- user id -- the unique identification of user
// Key -- api key
// Username -- user's name
// Password -- user's password
// Email -- user's Email
// Phone -- user's phone number

type User struct {
     ID        int       `json:"id"`
     Key       string    `json:"key"`
     Username  stirng    `json:"username"`
     Password  string    `json:"password"`
     Email     string    `json:"Email"`
     Phone     string    `json:"phone"`
}

//Create a new User instance and renturn

func NewUser(username string, password string, email string, phone string) *User{
    return &User{-1,"",username, password, email, phone}
    if len(userName) == 0 {
        panic("UserName shold not null!")
    }
    if len(password) < 5 {
        panic("Password must be more than 4 words!")
    }
}
