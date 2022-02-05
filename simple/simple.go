package simple

type SimpleRepository struct {
}

func NewSimpleRepository() *SimpleRepository {
	return &SimpleRepository{}
}

type StructService struct {
	*SimpleRepository
}

func NewStructService(repository *SimpleRepository) *StructService {
	return &StructService{
		SimpleRepository: repository,
	}
}
