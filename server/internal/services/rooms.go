package services

import (
	"database/sql"
	"errors"
	"github.com/iarsham/websocket-chat/internal/domain"
	"github.com/iarsham/websocket-chat/internal/entities"
	"github.com/iarsham/websocket-chat/internal/models"
	"go.uber.org/zap"
)

type RoomsService struct {
	db  *sql.DB
	log *zap.Logger
}

func NewRoomsService(db *sql.DB, log *zap.Logger) domain.RoomsRepository {
	return &RoomsService{
		db:  db,
		log: log,
	}
}

func (r *RoomsService) GetRoomByID(id string) (*models.Rooms, error) {
	query := "SELECT * FROM rooms WHERE id=$1;"
	row := r.db.QueryRow(query, id)
	return r.collectRow(row)
}

func (r *RoomsService) GetRoomByName(name string) (*models.Rooms, error) {
	query := "SELECT * FROM rooms WHERE name=$1;"
	row := r.db.QueryRow(query, name)
	return r.collectRow(row)
}

func (r *RoomsService) GetAllRooms() (*[]models.Rooms, error) {
	query := "SELECT * FROM rooms;"
	rows, err := r.db.Query(query)
	if err != nil {
		r.log.Warn(err.Error())
		return nil, err
	}
	defer rows.Close()
	return r.collectRows(rows)
}

func (r *RoomsService) CreateRoom(req *entities.RoomRequest) (*models.Rooms, error) {
	query := "INSERT INTO rooms (name) VALUES ($1) RETURNING *"
	stat, err := r.db.Prepare(query)
	if err != nil {
		r.log.Warn(err.Error())
		return nil, err
	}
	defer stat.Close()
	row := stat.QueryRow(req.Name)
	return r.collectRow(row)
}

func (r *RoomsService) UpdateRoom(req *entities.RoomRequest, roomID string) (*models.Rooms, error) {
	query := "UPDATE rooms SET name=$1 WHERE id=$2 RETURNING *"
	stat, err := r.db.Prepare(query)
	if err != nil {
		r.log.Warn(err.Error())
		return nil, err
	}
	defer stat.Close()
	row := stat.QueryRow(req.Name, roomID)
	return r.collectRow(row)
}

func (r *RoomsService) DeleteRoom(roomID string) error {
	query := "DELETE FROM rooms where id=$1;"
	if _, err := r.db.Exec(query, roomID); err != nil {
		r.log.Warn(err.Error())
		return err
	}
	return nil
}

func (r *RoomsService) collectRow(row *sql.Row) (*models.Rooms, error) {
	var room models.Rooms
	err := row.Scan(&room.ID, &room.Name)
	if err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			r.log.Warn(err.Error())
		}
		return nil, err
	}
	return &room, nil
}

func (r *RoomsService) collectRows(rows *sql.Rows) (*[]models.Rooms, error) {
	var rooms []models.Rooms
	for rows.Next() {
		var room models.Rooms
		if err := rows.Scan(&room.ID, &room.Name); err != nil {
			r.log.Warn(err.Error())
			return nil, err
		}
		rooms = append(rooms, room)
	}
	if err := rows.Err(); err != nil {
		r.log.Warn(err.Error())
		return nil, err
	}
	return &rooms, nil
}
