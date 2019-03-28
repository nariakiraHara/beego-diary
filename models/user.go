package models

import (
	"errors"
	"strconv"
	"time"

	"github.com/astaxie/beego/orm"
	"golang.org/x/crypto/bcrypt"
)

var (
	UserList map[string]*User
)

// User is struct user info
type User struct {
	ID       int64  `orm:"column(id);pk" json:"id"`
	Username string `orm:"default(1)" json:"userName"`
	Password string
	Email    string
	Birthday time.Time `orm:"null;type(datetime)"`
	Comment  string    `orm:"null" json:"comment"`
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
	hashPass, err := passwordHash(u.Password)
	if err != nil {
		panic(err.Error())
	}
	u.Password = hashPass
	u.IsActive = 1
	res, err := o.Insert(u)
	if err != nil {
		panic(err.Error())
	}
	return strconv.FormatInt(res, 10)
}

// GetUser returns user information from userID
func GetUser(uid int64) (u *User, err error) {
	o := orm.NewOrm()
	user := User{ID: uid}
	if err := o.Read(&user); err == nil {
		u = &user
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

func passwordHash(pw string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(pw), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), err
}
