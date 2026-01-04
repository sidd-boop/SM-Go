/*
In database management, "migration" refers to a formal process for managing changes to the database schema over time in a controlled and versioned manner [2]. This process ensures that changes—such as adding a table, modifying a column, or creating an index—are applied consistently across different environments (e.g., development, testing, production) [2]. 
Key Concepts
Version Control for Databases: Database migrations bring the principles of source code version control to the database schema [1]. Each change is treated as a "version" or "migration script," often stored alongside the application's source code [1, 3].
Schema Evolution: The primary goal is to allow the database schema to evolve in sync with the application's development. This is crucial for collaborative development, ensuring all developers and environments are working with the correct, up-to-date schema [1, 2].
Automation and Tooling: Specialized tools and frameworks (like Rails Migrations, Entity Framework Migrations, Flyway, or Liquibase) automate the application and tracking of these changes [3, 4, 5]. These tools typically track which migrations have been applied to a specific database instance [3].
Forward and Backward Compatibility: Each migration script often includes instructions for how to apply a change ("up") and how to reverse it ("down") [2]. The ability to reverse a change (rollback) is a critical safety feature, allowing developers to revert to a previous, stable state if issues arise [2]. 
*/
package auth

import (
	"errors"
	"gorm.io/gorm"
)

type Repository interface {
	CreateUser(user *User) error
	FindByEmail(email string) (*User, error)
}
type Repo struct {
	db *gorm.DB
}

func NewRepo(db *gorm.DB) *Repo {
	return &Repo{db: db}
}

//CreateUser
func (r *Repo) CreateUser(user *User) error {
	return r.db.Create(user).Error
}

//FindByEmail
func (r *Repo) FindByEmail(email string) (*User, error) {
	var user User
	err :=r.db.Where("email = ?",email).First(&user).Error
	if err !=nil{
		return nil, errors.New("user not found")
	}
	return &user, nil
}