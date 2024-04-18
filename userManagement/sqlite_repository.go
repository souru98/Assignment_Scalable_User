package userManagement

import (
	"database/sql"
	"errors"

	"github.com/mattn/go-sqlite3"
)

var (
	ErrDuplicate    = errors.New("record already exists")
	ErrNotExists    = errors.New("row not exists")
	ErrUpdateFailed = errors.New("update failed")
	ErrDeleteFailed = errors.New("delete failed")
)

type SQLiteRepository struct {
	db *sql.DB
}

func NewSQLiteRepository(db *sql.DB) *SQLiteRepository {
	return &SQLiteRepository{
		db: db,
	}
}

func (r *SQLiteRepository) Migrate() error {
	query := `
    CREATE TABLE IF NOT EXISTS users(
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        name TEXT NOT NULL,
        UserName TEXT NOT NULL UNIQUE,
        Password TEXT NOT NULL,
		IsActive BOOLEAN NOT NULL,
		IsInternal BOOLEAN NOT NULL
    );
    `

	_, err := r.db.Exec(query)
	return err
}

func (r *SQLiteRepository) Create(user User) (*User, error) {
	res, err := r.db.Exec("INSERT INTO users(name, UserName, Password, IsActive, IsInternal) values(?,?,?,?,?)", user.Name, user.UserName, user.Password, user.IsActive, user.IsInternal)
	if err != nil {
		var sqliteErr sqlite3.Error
		if errors.As(err, &sqliteErr) {
			if errors.Is(sqliteErr.ExtendedCode, sqlite3.ErrConstraintUnique) {
				return nil, ErrDuplicate
			}
		}
		return nil, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}
	user.ID = id

	return &user, nil
}

func (r *SQLiteRepository) All() ([]UserResponse, error) {
	rows, err := r.db.Query("SELECT id, Name, UserName, IsActive FROM users where IsInternal = ?", false)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var all []UserResponse
	for rows.Next() {
		var user UserResponse
		if err := rows.Scan(&user.ID, &user.Name, &user.UserName, &user.IsActive); err != nil {
			return nil, err
		}
		all = append(all, user)
	}
	return all, nil
}

func (r *SQLiteRepository) GetByName(name string) (*UserResponse, error) {
	row := r.db.QueryRow("SELECT id, Name, UserName, IsActive FROM users WHERE name = ?", name)

	var user UserResponse
	if err := row.Scan(&user.ID, &user.Name, &user.UserName, &user.IsActive); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrNotExists
		}
		return nil, err
	}
	return &user, nil
}

func (r *SQLiteRepository) GetByID(id string) (*UserResponse, error) {
	row := r.db.QueryRow("SELECT id, Name, UserName, IsActive FROM users WHERE id = ?", id)

	var user UserResponse
	if err := row.Scan(&user.ID, &user.Name, &user.UserName, &user.IsActive); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrNotExists
		}
		return nil, err
	}
	return &user, nil
}

func (r *SQLiteRepository) Validate(un string, pwd string) (*UserResponse, error) {
	if un == "" || pwd == "" {
		return nil, errors.New("invalid input parameters")
	}
	row := r.db.QueryRow("select id, Name, UserName, IsActive from users where username =? and password =? and IsActive =?", un, pwd, true)

	var user UserResponse
	if err := row.Scan(&user.ID, &user.Name, &user.UserName, &user.IsActive); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrNotExists
		}
		return nil, err
	}
	return &user, nil
}

func (r *SQLiteRepository) Update(id string, updated User) (*User, error) {
	if id == "0" {
		return nil, errors.New("invalid updated ID")
	}
	res, err := r.db.Exec("UPDATE users SET Name = ?, Password = ?, IsActive = ? WHERE id = ?", updated.Name, updated.Password, updated.IsActive, id)
	if err != nil {
		return nil, err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return nil, err
	}

	if rowsAffected == 0 {
		return nil, ErrUpdateFailed
	}

	return &updated, nil
}

func (r *SQLiteRepository) Delete(id int64) error {
	res, err := r.db.Exec("DELETE FROM users WHERE id = ?", id)
	if err != nil {
		return err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return ErrDeleteFailed
	}

	return err
}
