package dto

type SimpleCommitInput struct {
	Commit string
}

type CommitInput struct {
	IdLogin       int
	IdServicePage int
	Commit        string
}

type CommitOutput struct {
	Login  string
	Commit string
}
