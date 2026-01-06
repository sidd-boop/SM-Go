package thought

import "gorm.io/gorm"

type Repository interface{
	Create(thought *Thought) error
	FindAll() ([]Thought, error)
	FindByID(id uint) (*Thought, error)
}

type repo struct{
	db *gorm.DB
}

func NewRepo(db *gorm.DB) *repo{
	if db==nil{
		panic("db is nil")
	}
	return &repo{db: db}
}

//create
func (r *repo) Create(thought *Thought) error{
	return r.db.Create(thought).Error
}

//findall
func (r *repo) FindAll()([]Thought, error){
	var thoughts []Thought
	err :=r.db.Order("created_at desc").Find(&thoughts).Error
	return thoughts, err
}

//find by id
func (r *repo) FindByID(id uint) (*Thought, error){
	var thought Thought
	err:=r.db.First(&thought, id).Error
	if err!=nil{
		return nil,err
	}
	return &thought,nil
}

