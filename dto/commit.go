package dto

type SimpleCommitInput struct {
	Commit string
}

type CommitInput struct {
	UserId        int
	ServicePageId int
	Commit        string
}

type CommitOutput struct {
	Login  string
	Commit string
}
