package thought

type Service struct{
	repo Repository
}

func NewService(repo Repository) *Service{
	return &Service{repo: repo}
}

//create thought
func (s *Service) Create(content, tag string, userID uint) error{
	thought:=Thought{
		Content:content,
		Tag: tag,
		UserID: userID,
	}
	return s.repo.Create(&thought)
}

//list thoughts
func (s *Service) List() ([]Thought, error){
	return s.repo.FindAll()
}

//get thought
func (s *Service) Get(id uint) (*Thought, error){
	return s.repo.FindByID(id)
}