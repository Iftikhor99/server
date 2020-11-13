package banners

import (
	"context"
	"errors"
	"sync"
)

// Service npenctasnset co6oi cepsuc no ynpasnenwo OaHHepamn.
type Service struct {
	mu sync.RWMutex

	items []*Banner
}

// NewService co3qaét cepsuc.
func NewService() *Service {
	return &Service{items: make([]*Banner, 0)}

}

// Banner npenctasnaet codoi GaHHep.
type Banner struct {
	ID int64

	Title string

	Content string

	Button string

	Link string
}

// ByID Bo3BpawaeT OaHHep no upeHTHOuKaTopy.
func (s *Service) ByID(ctx context.Context, id int64) (*Banner, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	for _, banner := range s.items {
		if banner.ID == id {
			return banner, nil
		}
	}

	return nil, errors.New("item not found")
}

// All for
func (s *Service) All(ctx context.Context) ([]*Banner, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	// for _, banner := range s.items {
	// 	if banner.ID == id {
	// 		return banner, nil
	// 	}
	// }
	banners := s.items

	return banners, nil
	//panic("not implemented")
}

// Save for
func (s *Service) Save(ctx context.Context, item *Banner) (*Banner, error) {
	panic("not implemented")
}

// RemoveByID for
func (s *Service) RemoveByID(ctx context.Context, id int64) (*Banner, error) {
	panic("not implemented")
}

// Init for
func (s *Service) Init() {
	banner := Banner{
		ID: 1,

		Title: "string Title",

		Content: "string Content",

		Button: "string Button",

		Link: "string Link",
	}

	//item := s.items
	s.items = append(s.items, &banner)
	//item[1] = &banner
	//	panic("not implemented")
}
