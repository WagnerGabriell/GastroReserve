package repositories

import (
	"GastroReserve/internal/entities"
	"database/sql"
)

type UserRepositoryMySql struct {
	Db *sql.DB
}

func NewUserRepositoryMySql(db *sql.DB) *UserRepositoryMySql {
	return &UserRepositoryMySql{Db: db}
}

func (r *UserRepositoryMySql) GetAllUsers() ([]*entities.User, error) {
	var users []*entities.User
	rows, err := r.Db.Query("SELECT id,email,password,name,phoneNumber isAdmin FROM user")
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var user entities.User
		err := rows.Scan(&user.Id, &user.Email, &user.Password, &user.Name, &user.PhoneNumber, &user.IsAdmin)
		if err != nil {
			return nil, err
		}
		users = append(users, &user)
	}
	return users, nil
}
func (r *UserRepositoryMySql) CreateUser(user entities.User) error {
	_, err := r.Db.Exec("INSERT INTO user (id,email,password,name,phoneNumber,isAdmin) VALUES (?,?,?,?,?,?)",
		user.Id, user.Email, user.Password, user.Name, user.PhoneNumber, user.IsAdmin)
	if err != nil {
		return err
	}
	return nil
}
func (r *UserRepositoryMySql) GetUser(id string) (*entities.User, error) {
	var user entities.User
	row := r.Db.QueryRow("SELECT id,email,password,name,phoneNumber,isAdmin FROM user WHERE id = ?", id)
	err := row.Scan(&user.Id, &user.Email, &user.Password, &user.Name, &user.PhoneNumber, &user.IsAdmin)
	if err != nil {
		return &entities.User{}, err
	}

	return &user, nil
}
func (r *UserRepositoryMySql) BecomeAdm(id string) error {
	_, err := r.Db.Exec("UPDADE user SET isAdmin = true WHERE id = ?", id)
	if err != nil {
		return err
	}
	return nil
}
func (r *UserRepositoryMySql) GetUserPerEmail(email string) (*entities.User, error) {
	var user entities.User
	row := r.Db.QueryRow("SELECT id,email,password,name,phoneNumber,isAdmin FROM user WHERE email = ?", email)
	err := row.Scan(&user.Id, &user.Email, &user.Password, &user.Name, &user.PhoneNumber, &user.IsAdmin)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
