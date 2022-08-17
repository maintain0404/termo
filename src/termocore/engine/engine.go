package engine

import (
	"github.com/maintain0404/termocore/storage"
)

type Engine interface {
	//Constructors
	NewEngine()

	//properties
	GetStorage() storage.Storage
}
