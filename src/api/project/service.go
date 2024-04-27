package project

type ProjectInteractor struct {
	ProjectRepositry ProjectRepositry
}

func NewProjectInteractor(repo ProjectRepositry) *ProjectInteractor {
	return &ProjectInteractor{repo}
}
