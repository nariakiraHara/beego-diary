package models

import (
	"errors"
	"strconv"
	"time"

	"github.com/astaxie/beego/orm"
)

var (
	UserList map[string]*User
)

func init() {
	UserList = make(map[string]*User)
	u := User{1, "astaxie", "11111", "example@gmail.com", 1, time.Now(), time.Now()}
	UserList["user_11111"] = &u
}

// User is struct user info
type User struct {
	ID       int64  `orm:"column(id);pk" json:"id"`
	Username string `orm:"default(1)" json:"userName"`
	Password string `json:"-"`
	Email    string
	IsActive int
	Created  time.Time `orm:"auto_now_add;type(datetime)" json:"created"`
	Updated  time.Time `orm:"auto_now_add;type(datetime)" json:"updated"`
}

// TableName return tableName string
func (this *User) TableName() string {
	return "users"
}

// AddUser returns result ? ok : ng
func AddUser(u *User) string {
	o := orm.NewOrm()
	o.Using("users")
	u.IsActive = 1
	res, err := o.Insert(u)
	if err != nil {
		panic(err.Error())
	}
	return strconv.FormatInt(res, 10)
}

// GetUser returns user information from userID
func GetUser(uid string) (u *User, err error) {
	if u, ok := UserList[uid]; ok {
		return u, nil
	}
	return u, errors.New("User not exists")
}

// GetAllUsers is called GET METHOD
// returns all users information
func GetAllUsers() map[string]*User {
	return UserList
}

func UpdateUser(uid string, uu *User) (a *User, err error) {
	if u, ok := UserList[uid]; ok {
		if uu.Username != "" {
			u.Username = uu.Username
		}
		if uu.Password != "" {
			u.Password = uu.Password
		}
		return u, nil
	}
	return nil, errors.New("User Not Exist")
}

func Login(username, password string) bool {
	for _, u := range UserList {
		if u.Username == username && u.Password == password {
			return true
		}
	}
	return false
}

func DeleteUser(uid string) {
	delete(UserList, uid)
}
