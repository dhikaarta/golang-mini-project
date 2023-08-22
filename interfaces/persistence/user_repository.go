package persistence

import(
	"database/sql"
	"fmt"
	"task-management-app/internal/entity"
)

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *userRepository {
	return &userRepository{
		db: db,
	}
}

func (u *userRepository) CreateUser(user entity.User) (string, error) {
	result, err := u.db.Exec("INSERT INTO users (username, email, password) VALUES (?, ?, ?)", user.Username, user.Email, user.Password)
	if err != nil {
		return "", err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return "", nil
	}

	return fmt.Sprintf("%d", id), nil
}

func (u *userRepository) GetUserByUsername(username string) (entity.User, error) {
	var user entity.User
	err := u.db.QueryRow("SELECT * FROM users WHERE username = ?", username).Scan(&user.ID, &user.Username, &user.Email, &user.Password)
	if err != nil {
		return user, err
	}

	return user, nil
}

func (u *userRepository) ListUsers() ([]entity.User, error) {
	rows, err := u.db.Query("SELECT * FROM users")
	if err != nil {
		return nil, err
	}

	var users []entity.User
	for rows.Next() {
		var user entity.User
		if err := rows.Scan(&user.ID, &user.Username, &user.Email, &user.Password); err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}

func (u *userRepository) UpdateUser(user entity.User) error {
	_, err := u.db.Exec("UPDATE users SET username = ?, email = ?, password = ? WHERE id = ?", user.Username, user.Email, user.Password, user.ID)
	if err != nil {
		return err
	}

	return nil
}

func (u *userRepository) DeleteUser(userID string) error {
	_, err := u.db.Exec("DELETE FROM users WHERE id = ?", userID)
	if err != nil {
		return err
	}

	return nil
}

func (u *userRepository) CheckUserExist(username string) (bool, error) {
	var user entity.User
	err := u.db.QueryRow("SELECT * FROM users WHERE username = ?", username).Scan(&user.ID, &user.Username, &user.Email, &user.Password)
	if err != nil {
		return false, err
	}

	return true, nil
}

func (u *userRepository) CheckEmailExist(email string) (bool, error) {
	var user entity.User
	err := u.db.QueryRow("SELECT * FROM users WHERE email = ?", email).Scan(&user.ID, &user.Username, &user.Email, &user.Password)
	if err != nil {
		return false, err
	}

	return true, nil
}