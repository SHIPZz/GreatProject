package Factory

import (
	"GreatProject/Data/Repository"
	"GreatProject/Infrastructure/Configs"
	"log"
)

type RepositoryType int

type IRepositoryFactory interface {
	GetTaskRepository(repoType RepositoryType) Repository.ITaskRepository
}

const (
	InMemory RepositoryType = iota
	Postgres
	File
)

type RepositoryFactory struct {
	ConnectionString string
}

func NewRepositoryFactory() IRepositoryFactory {
	return &RepositoryFactory{
		ConnectionString: Configs.ConnectionString,
	}
}

func (rf *RepositoryFactory) GetTaskRepository(repoType RepositoryType) Repository.ITaskRepository {
	switch repoType {
	case InMemory:
		return Repository.NewInMemoryTaskRepository()
	case Postgres:
		repo, err := Repository.NewPostgresTaskRepository(rf.ConnectionString)
		if err != nil {
			log.Fatalf("Failed to create Postgres repository: %v", err)
		}
		return repo
	case File:
		return Repository.NewInMemoryTaskRepository()
	default:
		return Repository.NewInMemoryTaskRepository()
	}
}

type MockRepositoryFactory struct{}

func NewMockRepositoryFactory() IRepositoryFactory {
	return &MockRepositoryFactory{}
}

func (mrf *MockRepositoryFactory) GetTaskRepository(RepositoryType) Repository.ITaskRepository {
	return Repository.NewInMemoryTaskRepository()
}
