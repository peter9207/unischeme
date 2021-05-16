package server

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

type Server struct {
	Name string
	URL  string

	router  *gin.Engine
	signals chan bool
	port    int
}

var nodes = map[string]string{}

type AddNodeRequest struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type PingResponse struct {
	Name string `json:"string"`
}

func New(name, url, existingNodes string) (server *Server) {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, PingResponse{Name: name})
	})

	r.GET("/nodes", func(c *gin.Context) {
		c.JSON(200, nodes)
	})

	r.POST("/node", func(c *gin.Context) {
		message := AddNodeRequest{}
		c.Bind(&message)
		nodes[message.Name] = message.URL
		c.JSON(200, nodes)
	})
	server = &Server{
		router: r,
	}
	return
}

func (s *Server) register(existingNodes string) (err error) {

	addNodeRequest := AddNodeRequest{
		Name: s.Name,
		URL:  s.URL,
	}

	data, err := json.Marshal(addNodeRequest)
	if err != nil {
		return
	}

	split := strings.Split(existingNodes, ",")
	for _, url := range split {
		var resp *http.Response
		resp, err = http.Post(url, "application/json", bytes.NewBuffer(data))
		if err != nil {
			return
		}
		if resp.StatusCode != 200 {
			log.Info().Str("url", url).Msg("failed to register with node")
			continue
		}

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Info().Str("url", url).Msg("failed to read body")
			continue
		}
		result := PingResponse{}
		err = json.Unmarshal(body, &result)
		if err != nil {
			log.Info().Str("url", url).Msg("failed to parse response")
			continue
		}

		nodes[result.Name] = url
		resp.Body.Close()
	}

	return
}

func (s *Server) Start() {

	go checkAlive(s.signals)
	s.router.Run()
}

func (s *Server) Stop() {
	s.signals <- true
}

func checkAlive(signalCh chan bool) {

	for {
		select {
		case <-signalCh:
			return
		case <-time.After(time.Minute):
		}
		for name, url := range nodes {
			_, err := http.Get(fmt.Sprintf("%s/ping", url))
			if err != nil {
				log.Info().Str(name, url).Msgf("failed healthcheck removing from list of nodes")

				delete(nodes, name)
				continue
			}
		}

	}
}
