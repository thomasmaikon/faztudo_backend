package dto

type LoginDTO struct {
	Login    string
	Password string
}

func (l LoginDTO) IsEmpty() bool {
	if l.Login == "" && l.Password == "" {
		return true
	}
	return false
}
