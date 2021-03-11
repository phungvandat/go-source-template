package usecase

import "github.com/phungvandat/source-template/usecase/authen"

// Usecase holds all usecase
type Usecase struct {
	Authen authen.UseCase
}
