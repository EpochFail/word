package word

import (
	"encoding/json"
	"log"
	"net"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

func CreateRouter(server *HTTPServer) (*mux.Router, error) {
	r := mux.NewRouter()
	m := map[string]map[string]HttpApiFunc{
		"GET": {
			"/api/word":     server.GetRandomWord,
			"/api/history":  server.GetHistory,
			"/api/top10":    server.GetTop10,
			"/api/bottom10": server.GetBottom10,
			"/api/random10": server.GetRandom10,
		},
		"POST": {
			"/api/vote/{word}/up":   server.UpVoteWord,
			"/api/vote/{word}/down": server.DownVoteWord,
		},
		"OPTIONS": {
			"": options,
		},
	}

	for method, routes := range m {
		for route, handler := range routes {
			localRoute := route
			localHandler := handler
			localMethod := method
			f := makeHttpHandler(localMethod, localRoute, localHandler)

			if localRoute == "" {
				r.Methods(localMethod).HandlerFunc(f)
			} else {
				r.Path(localRoute).Methods(localMethod).HandlerFunc(f)
			}
		}
	}

	return r, nil
}

func makeHttpHandler(localMethod string, localRoute string, handlerFunc HttpApiFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		writeCorsHeaders(w, r)
		if err := handlerFunc(w, r, mux.Vars(r)); err != nil {
			httpError(w, err)
		}
	}
}

func writeCorsHeaders(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept")
	w.Header().Add("Access-Control-Allow-Methods", "GET, POST, DELETE, PUT, OPTIONS")
}

type HttpApiFunc func(w http.ResponseWriter, r *http.Request, vars map[string]string) error

type HTTPServer struct {
	DB *DB
}

func NewHTTPServer(db *DB) *HTTPServer {
	s := &HTTPServer{
		DB: db,
	}

	return s
}

func writeJSON(w http.ResponseWriter, code int, thing interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	val, err := json.Marshal(thing)
	w.Write(val)
	return err
}

func httpError(w http.ResponseWriter, err error) {
	statusCode := http.StatusInternalServerError

	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), statusCode)
	}
}

func options(w http.ResponseWriter, r *http.Request, vars map[string]string) error {
	w.WriteHeader(http.StatusOK)
	return nil
}

func (s *HTTPServer) GetRandomWord(w http.ResponseWriter, r *http.Request, vars map[string]string) error {
	word, err := s.DB.GetRandomWord()

	if err != nil {
		return err
	}

	writeJSON(w, http.StatusOK, word)

	return nil
}

func (s *HTTPServer) GetHistory(w http.ResponseWriter, r *http.Request, vars map[string]string) error {
	words, err := s.DB.GetHistory()

	if err != nil {
		return err
	}

	writeJSON(w, http.StatusOK, words)

	return nil
}

func (s *HTTPServer) GetTop10(w http.ResponseWriter, r *http.Request, vars map[string]string) error {
	words, err := s.DB.GetTop10()

	if err != nil {
		return err
	}

	writeJSON(w, http.StatusOK, words)

	return nil
}

func (s *HTTPServer) GetBottom10(w http.ResponseWriter, r *http.Request, vars map[string]string) error {
	words, err := s.DB.GetBottom10()

	if err != nil {
		return err
	}

	writeJSON(w, http.StatusOK, words)

	return nil
}

func (s *HTTPServer) GetRandom10(w http.ResponseWriter, r *http.Request, vars map[string]string) error {
	words, err := s.DB.GetRandom10()

	if err != nil {
		return err
	}

	writeJSON(w, http.StatusOK, words)

	return nil
}

func (s *HTTPServer) UpVoteWord(w http.ResponseWriter, r *http.Request, vars map[string]string) error {
	word := vars["word"]
	ip, err := ipForRequest(r)
	if err != nil {
		return err
	}

	err = s.DB.UpVoteWord(word, ip)
	if err != nil {
		return err
	}
	w.WriteHeader(http.StatusOK)
	return nil
}

func (s *HTTPServer) DownVoteWord(w http.ResponseWriter, r *http.Request, vars map[string]string) error {
	word := vars["word"]
	ip, err := ipForRequest(r)
	if err != nil {
		return err
	}

	err = s.DB.UpVoteWord(word, ip)
	if err != nil {
		return err
	}

	w.WriteHeader(http.StatusOK)
	return nil
}

func ipForRequest(r *http.Request) (string, error) {
	host, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		return "", err
	}

	realIp := r.Header.Get("X-Real-Ip")
	forwardedFor := r.Header.Get("X-Forwarded-For")

	if realIp != "" {
		host = realIp
	}

	if forwardedFor != "" {
		addresses := strings.Split(forwardedFor, ",")
		for _, ele := range addresses {
			if ele != "127.0.0.1" {
				host = strings.TrimSpace(ele)
				break
			}
		}
	}

	return host, nil
}
