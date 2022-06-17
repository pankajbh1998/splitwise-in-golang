package repo

import (
	"fmt"
	"splitwise-in-golang/models"
	"strings"
)

type UserMap struct {
	mp map[int64]*models.User
	strMap map[string]*models.User
}

func NewUserMap()*UserMap{
	return &UserMap{
		mp: make(map[int64]*models.User),
		strMap: make(map[string]*models.User),
	}
}

func (u *UserMap) AddUser(user *models.User){
	user.ID = models.GetUniqueID()
	u.mp[user.ID] = user
	u.strMap[strings.ToLower(user.Name)] = user
}

func (u *UserMap) GetUserByName(name string)(*models.User,error){
	val, ok := u.strMap[strings.ToLower(name)]
	if !ok {
		return nil, fmt.Errorf("User not found")
	}

	return val, nil
}

func (u *UserMap) GetUserByID(id int64)*models.User{
	val, _ := u.mp[id]
	return val
}
