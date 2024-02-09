package services

import (
	"database/sql"
	"errors"
	"github.com/google/uuid"
	"github.com/gorilla/sessions"
	"github.com/iarsham/websocket-chat/internal/domain"
	"github.com/iarsham/websocket-chat/internal/entites"
	"github.com/iarsham/websocket-chat/internal/models"
	"github.com/iarsham/websocket-chat/pkg/constans"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

type UserService struct {
	db    *sql.DB
	log   *zap.Logger
	store *sessions.CookieStore
}

func NewUserService(db *sql.DB, log *zap.Logger, store *sessions.CookieStore) domain.UserRepository {
	return &UserService{
		db:    db,
		log:   log,
		store: store,
	}
}

func (u *UserService) GetUserByID(id int64) (*models.Users, error) {
	query := "SELECT * FROM users WHERE id=$1;"
	row := u.db.QueryRow(query, id)
	return u.collectRow(row)
}

func (u *UserService) GetUserByUsername(userName string) (*models.Users, error) {
	query := "SELECT * FROM users WHERE username=$1;"
	row := u.db.QueryRow(query, userName)
	return u.collectRow(row)
}

func (u *UserService) CreateUser(req *entites.UserRequest) (*models.Users, error) {
	encryptPass, _ := bcrypt.GenerateFromPassword([]byte(req.Password), 10)
	query := "INSERT INTO users (username, password) VALUES ($1,$2) RETURNING *"
	stmt, err := u.db.Prepare(query)
	if err != nil {
		u.log.Warn(err.Error())
		return nil, err
	}
	defer stmt.Close()
	row := stmt.QueryRow(req.Username, encryptPass)
	return u.collectRow(row)
}

func (u *UserService) DeleteUser(userName string) error {
	query := "DELETE FROM users where username=$1;"
	if _, err := u.db.Exec(query, userName); err != nil {
		u.log.Warn(err.Error())
		return err
	}
	return nil
}

func (u *UserService) collectRow(row *sql.Row) (*models.Users, error) {
	var user models.Users
	err := row.Scan(&user.ID, &user.Username, &user.Password, &user.JoinedAt, &user.LastSeen, &user.Verified)
	if err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			u.log.Warn(err.Error())
		}
		return nil, err
	}
	return &user, nil
}

func (u *UserService) Authenticate(w http.ResponseWriter, r *http.Request, userID uuid.UUID, auth bool) error {
	session, err := u.store.Get(r, constans.Session)
	if err != nil {
		u.log.Warn(err.Error())
		return err
	}
	session.Values["authenticated"] = auth
	session.Values["user_id"] = userID
	if err := session.Save(r, w); err != nil {
		u.log.Warn(err.Error())
		return err
	}
	return nil
}
