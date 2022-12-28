package users

type UserService struct {
	UserRepository *UserRepository
}

func ProvideUserService(u *UserRepository) *UserService {
	return &UserService{UserRepository: u}
}
func (u *UserService) Find(id uint) (User, error) {
	return u.UserRepository.Find(id)
}

func (u *UserService) Create(dto CreateUserDTO) (User, error) {
	return u.UserRepository.Create(dto)
}

func (u *UserService) Update(id uint, dto UpdateUserDTO) (User, error) {
	return u.UserRepository.Update(id, dto)
}

func (u *UserService) Delete(id uint) (int64, error) {
	return u.UserRepository.Delete(id)
}
