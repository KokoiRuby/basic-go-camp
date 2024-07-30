package domain

// User domain obj, DDD entity, BO Business Object
type User struct {
	Email    string
	Password string
}

func (u *User) EncryptPassword(password string) {

}
