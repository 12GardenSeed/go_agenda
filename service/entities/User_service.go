package entities

// UserServiceProvider provide series of op to Entity users

type UserServiceProvider struct{

}

//an instance og UserServiceProvider
var UserService = UserServiceProvider{}

//Fuctions
func (*UserServiceProvider) Insert(user *User) error {
    dao := userDao{mydb}
    err := dao.Insert(user)
    CheckErr(err)
    return nil
}

func (*UserServiceProvider) FindAll() ([]Users, error) {
  dao := userDao{mydb}
  users, err := dao.FindAll()
  CheckErr(err)
  return users, nil
}

func (*UserServiceProvider) DeleteByKey(key string) error {
    dao := userDao{mydb}
    err := dao.DeleteByKey(key)
    if err != nil {
        return err
    }
    return nil
}
