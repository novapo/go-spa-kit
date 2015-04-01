package user

// User repressents a user in db
type User struct {
	ID           int
	Name         string
	PasswordHash string
}

func (u User) Save() {

}

func (u User) Delete() {

}
