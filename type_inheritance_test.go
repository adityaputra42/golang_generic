package golanggeneric

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Employe interface {
	GetName() string
}

func GetName[T Employe](param T) string {
	return param.GetName()
}

type Manager interface {
	GetName() string
	GetManagerName() string
}

type MyManager struct {
	Name string
}

func (m *MyManager) GetName() string {
	return m.Name
}

func (m *MyManager) GetManagerName() string {
	return m.Name
}

type VisePresident interface {
	GetName() string
	GetVisePresidentName() string
}

type MyVisePresident struct {
	Name string
}

func (m *MyVisePresident) GetName() string {
	return m.Name
}

func (m *MyVisePresident) GetVisePresidentName() string {
	return m.Name
}

func TestTypeInheritance(t *testing.T) {
	assert.Equal(t, "Aditya", GetName[Manager](&MyManager{Name: "Aditya"}))
	assert.Equal(t, "Aditya", GetName[VisePresident](&MyVisePresident{Name: "Aditya"}))
}
