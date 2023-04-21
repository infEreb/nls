package internal

type Controller struct {
	repo Repository
}

func NewController(repo Repository) *Controller {
	return &Controller{
		repo: repo,
	}
}