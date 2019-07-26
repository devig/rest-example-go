package repository

import (
	"database/sql"
	"log"
	"rest-example-go/entity"
	"rest-example-go/helper"

	"github.com/jmoiron/sqlx"
)

// UserDTO type is a struct for users.
type UserDTO struct {
	ID        sql.NullInt64 `db:"id"` //`json:"id,omitempty"` //если указатель на тип, то поле может быть null
	FirstName string        `db:"first_name"`
	LastName  string        `db:"last_name"`
	Email     string        `db:"email"`
	Password  string        `db:"password"`
	Sex       int           `db:"sex"`
}

// User2UserDTO conversion
func User2UserDTO(u *entity.User) *UserDTO {
	return &UserDTO{
		ID:        helper.ToNullInt64(u.ID),
		FirstName: u.FirstName,
		LastName:  u.LastName,
		Email:     u.Email,
		Password:  u.Password,
		Sex:       u.Sex,
	}
}

// Extract conversion
func Extract(u *UserDTO) *entity.User {
	//i := sql.NullInt64{Int64: 42, Valid: true}
	return &entity.User{
		ID:        u.ID.Int64,
		FirstName: u.FirstName,
		LastName:  u.LastName,
		Email:     u.Email,
		Password:  u.Password,
		Sex:       u.Sex,
	}
}

// ReadRepository type is a interface to read from DB.
type ReadRepository interface {
	FindAll() ([]*entity.User, error)
	FindOneByID(id int64) (*entity.User, error)
}

// WriteRepository type is a interface to write to DB.
type WriteRepository interface {
	Add(r *UserRepository) error
	Delete(id int64) error
	Update(id int64, r *UserRepository) error
}

// UserRepository type is a struct for users repository.
type UserRepository struct {
	db *sqlx.DB
}

// NewUserRepository initiate the service.
func NewUserRepository(db *sqlx.DB) *UserRepository {
	return &UserRepository{db}
}

// FindAll func retrieves all users.
func (r *UserRepository) FindAll() ([]*entity.User, error) {
	var users []*UserDTO
	err := r.db.Select(&users, "SELECT * FROM users")
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}

	var rez []*entity.User
	for _, v := range users {
		rez = append(rez, Extract(v))
	}

	return rez, nil
}

// FindOneByID func finds a user by a given ID.
func (r *UserRepository) FindOneByID(id int64) (*entity.User, error) {
	var user UserDTO
	//log.Println(id)
	err := r.db.Get(&user, "SELECT * FROM users WHERE id = $1 LIMIT 1", id)
	//err := r.db.QueryRowx("SELECT id, first_name, last_name, email, password, sex FROM users WHERE id = $1 LIMIT 1").StructScan(&user)
	if err != nil && err != sql.ErrNoRows {
		log.Println(err, sql.ErrNoRows)
		return nil, err
	}

	//log.Println(user)
	return Extract(&user), nil
}

// Add func add a new user.
func (r *UserRepository) Add(u *entity.User) error {
	var user *UserDTO
	user = User2UserDTO(u)
	if _, err := r.db.NamedExec("INSERT INTO users (first_name, last_name, email, password, sex) VALUES (:first_name, :last_name, :email, :password, :sex)", user); err != nil {
		return err
	}
	return nil
}

// Update func updates the given user.
func (r *UserRepository) Update(id int64, u *entity.User) error {
	var user *UserDTO
	user = User2UserDTO(u)
	user.ID = helper.ToNullInt64(id)
	if _, err := r.db.NamedExec("UPDATE users SET first_name=:first_name, last_name=:last_name, email=:email, password=:password, sex=:sex  WHERE id =:id ", user); err != nil {
		return err
	}
	return nil
}

// Delete func removes a user by a given ID.
func (r *UserRepository) Delete(id int64) error {
	//s := strconv.Itoa(id)
	if _, err := r.db.Exec("DELETE FROM users WHERE id = $1", id); err != nil {
		return err
	}
	return nil
}
