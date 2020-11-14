package app

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/Iftikhor99/server/pkg/banners"
)

// Server npegctasnseT coOow normyeckwi CepBep Hawero npunomeHna.
type Server struct {
	mux *http.ServeMux

	bannersSvc *banners.Service
}

// NewServer - OyHKUMA-KOHCTpykTOp pina co3maHna cepsepa.
func NewServer(mux *http.ServeMux, bannersSvc *banners.Service) *Server {

	return &Server{mux: mux, bannersSvc: bannersSvc}

}

//ServeHTTP for
func (s *Server) ServeHTTP(writer http.ResponseWriter, request *http.Request) {

	s.mux.ServeHTTP(writer, request)

}

// Init wHmunannsupyet cepsep (permctpupyet sce Handler's)
func (s *Server) Init() {
	s.mux.HandleFunc("/banners.getAll", s.handleGetAllBanners)
	//	s.mux.HandleFunc("/banners.getById", s.handleGetBannerByID)
	s.mux.HandleFunc("/banners.save", s.handleSaveBanner)

	s.mux.HandleFunc("/banners.removeById", s.handleRemoveByID)
	s.mux.HandleFunc("/banners.getById", s.handleGetPostByID)

}

func (s *Server) handleGetPostByID(writer http.ResponseWriter, request *http.Request) {

	idParam := request.URL.Query().Get("id")

	id, err := strconv.ParseInt(idParam, 10, 64)

	if err != nil {
		log.Print(err)
		http.Error(writer, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	item, err := s.bannersSvc.ByID(request.Context(), id)

	if err != nil {
		log.Print(err)
		http.Error(writer, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	data, err := json.Marshal(item)

	if err != nil {
		log.Print(err)
		http.Error(writer, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	writer.Header().Set("Content-Type", "application/json")
	_, err = writer.Write(data)

	if err != nil {
		log.Print(err)
	}

}

func (s *Server) handleRemoveByID(writer http.ResponseWriter, request *http.Request) {
	log.Print(request)
}

func (s *Server) handleSaveBanner(writer http.ResponseWriter, request *http.Request) {
	log.Print(request)
	log.Print(request.Header)
	log.Print(request.Body)

	banner := s.bannersSvc.Initial(request)

	//idParam := request.URL.Query().Get("id")

	//id, err := strconv.ParseInt(idParam, 10, 64)

	// if err != nil {
	// 	log.Print(err)
	// 	http.Error(writer, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
	// 	return
	// }

	item, err := s.bannersSvc.Save(request.Context(), &banner)

	if err != nil {
		log.Print(err)
		http.Error(writer, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	data, err := json.Marshal(item)

	if err != nil {
		log.Print(err)
		http.Error(writer, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	writer.Header().Set("Content-Type", "application/json")
	_, err = writer.Write(data)

	if err != nil {
		log.Print(err)
	}

}

func (s *Server) handleGetBannerByID(writer http.ResponseWriter, request *http.Request) {
	log.Print(request)
}

func (s *Server) handleGetAllBanners(writer http.ResponseWriter, request *http.Request) {
	log.Print(request)
	log.Print(request.Header)
	log.Print(request.Body)
	item, err := s.bannersSvc.All(request.Context())

	if err != nil {
		log.Print(err)
		http.Error(writer, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	data, err := json.Marshal(item)

	if err != nil {
		log.Print(err)
		http.Error(writer, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	writer.Header().Set("Content-Type", "application/json")
	_, err = writer.Write(data)

	if err != nil {
		log.Print(err)
	}

	log.Print(item)
}
