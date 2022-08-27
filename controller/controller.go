package controller

type Controller interface {
	Start(string) error
	Stop()
}
