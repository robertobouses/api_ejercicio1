package user

type Repository interface {
	InsertRectangle()
}

type userAppService struct {
	repo Repository
}

func NewUserAppService(repo Repository) (*userAppService, error) {
	return &userAppService{
		repo: repo,
	}, nil
}
