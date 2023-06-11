package dto

type LoginDTO struct {
	Login    string
	Password string
	User     UserDTO
}

func (l LoginDTO) IsEmpty() bool {
	if l.Login == "" && l.Password == "" {
		return true
	}
	return false
}
