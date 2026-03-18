package usecases

type IHealthUseCase interface {
	Apply() error
}

type HealthUseCase struct{}

func (u *HealthUseCase) Apply() error {
	return nil
}

func NewHealthUseCase() *HealthUseCase {
	return &HealthUseCase{}
}
