package handler

type ParentService interface {
	StartService() (string, error)
}
type ParentServiceImpl struct {
}

func (service *ParentServiceImpl) StartService() (string, error) {
	return "ok", nil
}
