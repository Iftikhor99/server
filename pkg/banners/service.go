package banners

import (
	"context"
	"errors"
	"log"
	"net/http"
	"strconv"
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
	// banners := s.items
	// if len(s.items) == 0 {
	// 	return nil, errors.New("no items found")
	// }
	return s.items, nil
	//panic("not implemented")
}

// Save for
func (s *Service) Save(ctx context.Context, item *Banner) (*Banner, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	//	for _, banner := range s.items {
	if item.ID == 0 {

		item.ID = int64(len(s.items)) + 1
		s.items = append(s.items, item)
		return item, nil
	}
	if item.ID != 0 {
		for _, banner := range s.items {
			if banner.ID == item.ID {
				banner.Button = item.Button
				banner.Content = item.Content
				banner.Link = item.Link
				banner.Title = item.Title
				return item, nil
			}
		}
	}
	return nil, errors.New("item not found")

}

// RemoveByID for
func (s *Service) RemoveByID(ctx context.Context, id int64) (*Banner, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	for _, banner := range s.items {
		if banner.ID == id {
			retunBanner := banner
			banner.ID = 0
			banner.Button = ""
			banner.Content = ""
			banner.Link = ""
			banner.Title = ""

			return retunBanner, nil
		}
	}

	return nil, errors.New("item not found")
}

// Initial for
func (s *Service) Initial(request *http.Request) Banner {

	idParam := request.URL.Query().Get("id")
	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		log.Print(err)

	}

	titleParam := request.URL.Query().Get("title")
	contentParam := request.URL.Query().Get("content")
	buttonParam := request.URL.Query().Get("button")
	linkParam := request.URL.Query().Get("link")

	banner := Banner{
		ID: id,

		Title: titleParam,

		Content: contentParam,

		Button: buttonParam,

		Link: linkParam,
	}

	// banner2 := Banner{
	// 	ID: 2,

	// 	Title: "Title New",

	// 	Content: "Content New",

	// 	Button: "Button New",

	// 	Link: "Link New",
	// }

	//item := s.items
	//	s.items = append(s.items, &banner)
	//s.items = append(s.items, &banner2)
	//item[1] = &banner
	//	panic("not implemented")

	return banner
}
