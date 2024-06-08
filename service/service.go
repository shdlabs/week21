package service

import "errors"

type User struct {
	Fname   string  `json:"fname"`
	City    string  `json:"city"`
	Phone   string  `json:"phone"`
	ID      int32   `json:"id"`
	Height  float32 `json:"height"`
	Married bool    `json:"Married"`
}

// DbMock is a mock for DB
//
//	key := user.ID
//	value := User{}
type DbMock map[int32]User

func NewDBMock() DbMock {
	return DbMock{}
}

func (db DbMock) NewUsers(users ...User) {
	id := len(db)
	for i, user := range users {
		user.ID = int32(id + i)
		db[user.ID] = user

	}
}

func NewUser(fname, city string, phone string, height float32, married bool) User {
	return User{
		Fname:   fname,
		City:    city,
		Phone:   phone,
		Height:  height,
		Married: married,
	}
}

func (db DbMock) AddUser(id int32, fname, city string, phone string, height float32, married bool) error {
	if _, ok := db[id]; ok {
		return errors.New("user ID exists")
	}
	db[id] = User{
		ID:      id,
		Fname:   fname,
		City:    city,
		Phone:   phone,
		Height:  height,
		Married: married,
	}

	return nil
}

func (db DbMock) FindUser(id int32) User {
	if _, ok := db[id]; !ok {
		return User{}
	}

	return db[id]
}
