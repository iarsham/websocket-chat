package domain

import (
	"github.com/iarsham/websocket-chat/internal/entities"
	"github.com/iarsham/websocket-chat/internal/models"
)

type RoomsRepository interface {
	GetAllRooms() (*[]models.Rooms, error)
	GetRoomByID(id string) (*models.Rooms, error)
	GetRoomByName(name string) (*models.Rooms, error)
	CreateRoom(req *entities.RoomRequest) (*models.Rooms, error)
	UpdateRoom(req *entities.RoomRequest, roomID string) (*models.Rooms, error)
	DeleteRoom(roomID string) error
}
