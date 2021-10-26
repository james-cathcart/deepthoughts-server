package user

import "os/user"

type UserRepo interface {
    GetUserById(id string) user.User
    UpdateUser(user user.User) user.User
}