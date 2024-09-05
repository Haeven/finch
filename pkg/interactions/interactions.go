// pkg/interactions/interactions.go
package interactions

import (
	"context"
	"time"

	"flux/internal/db"

	socketio "github.com/googollee/go-socket.io"
	"github.com/uptrace/bun"
)

type InteractionService struct {
	db     *bun.DB
	socket *socketio.Server
}

type Interaction struct {
	bun.BaseModel `bun:"table:interactions"`

	ID        int64     `bun:"id,pk,autoincrement"`
	VideoID   string    `bun:"video_id,notnull"`
	UserID    string    `bun:"user_id,notnull"`
	Type      string    `bun:"type,notnull"`
	Content   string    `bun:"content"`
	CreatedAt time.Time `bun:"created_at,notnull"`
	UpdatedAt time.Time `bun:"updated_at,notnull"`
}

func NewInteractionService(socket *socketio.Server) *InteractionService {
	return &InteractionService{
		db:     db.Initialize(),
		socket: socket,
	}
}

func (s *InteractionService) HandleLike(videoID, userID string) error {
	interaction := &Interaction{
		VideoID:   videoID,
		UserID:    userID,
		Type:      "like",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	_, err := s.db.NewInsert().Model(interaction).Exec(context.Background())
	if err != nil {
		return err
	}

	s.socket.BroadcastToRoom("", interaction.VideoID, "new_like", interaction)
	return nil
}

func (s *InteractionService) HandleDislike(videoID, userID string) error {
	interaction := &Interaction{
		VideoID:   videoID,
		UserID:    userID,
		Type:      "dislike",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	_, err := s.db.NewInsert().Model(interaction).Exec(context.Background())
	if err != nil {
		return err
	}

	s.socket.BroadcastToRoom("", interaction.VideoID, "new_dislike", interaction)
	return nil
}

func (s *InteractionService) HandleComment(videoID, userID, content string) error {
	interaction := &Interaction{
		VideoID:   videoID,
		UserID:    userID,
		Type:      "comment",
		Content:   content,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	_, err := s.db.NewInsert().Model(interaction).Exec(context.Background())
	if err != nil {
		return err
	}

	s.socket.BroadcastToRoom("", interaction.VideoID, "new_comment", interaction)
	return nil
}

func (s *InteractionService) GetInteractions(videoID string) ([]Interaction, error) {
	var interactions []Interaction
	err := s.db.NewSelect().
		Model(&interactions).
		Where("video_id = ?", videoID).
		Order("created_at DESC").
		Scan(context.Background())

	return interactions, err
}

// Additional helper functions for database operations can be added here
