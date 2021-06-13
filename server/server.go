package server

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/peter9207/unischeme/interpreter"
	"github.com/peter9207/unischeme/lexer"
	"github.com/rs/zerolog/log"
	"io/ioutil"
	"net/http"
	"time"
)

type Server struct {
	Name       string
	URL        string
	CheckAlive time.Duration
	Router     *gin.Engine
	signals    chan bool
	port       int
}

var nodes = map[string]string{}

type AddNodeRequest struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type PingResponse struct {
	Name string `json:"string"`
}

type Program struct {
	main         interpreter.Expression
	declarations []interpreter.FunctionDeclaration
}

func Interpret(c *gin.Context) {

	var data string
	err := c.Bind(&data)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	lexed, err := lexer.Parse(data)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	ast, err := interpreter.ToAST(lexed.Expressions)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	var p Program

	for _, v := range ast {

		fn, ok := v.(*interpreter.FunctionDeclaration)
		if ok {
			p.declarations = append(p.declarations, *fn)
		}

		exp, ok := v.(interpreter.Expression)
		if ok {
			if p.main != nil {
				c.JSON(400, gin.H{
					"error": "can only have 1 main method",
				})
				return
			}
			p.main = exp
		}
	}

	varScope := make(map[string]interpreter.Expression)
	fnScope := make(map[string]interpreter.FunctionDeclaration)

	for _, d := range p.declarations {
		err = d.Perform(varScope, fnScope)
		if err != nil {
			c.JSON(400, gin.H{
				"error": err.Error(),
			})
			return
		}
	}

	v, err := p.main.Resolve(varScope, fnScope)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(200, v)
	return

}

func New(name, url string) (server *Server) {
	r := gin.Default()
	server = &Server{
		Router: r,
	}

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, PingResponse{Name: name})
	})

	r.GET("/nodes", func(c *gin.Context) {
		c.JSON(200, nodes)
	})

	r.POST("/nodes", func(c *gin.Context) {
		message := AddNodeRequest{}
		c.Bind(&message)

		err := server.register(message.URL)
		if err != nil {
			log.Info().Str("url", url).Msgf("adding broker failed: %s", err)
			c.JSON(400, "request failed")
			return
		}

		nodes[message.Name] = message.URL
		c.JSON(200, nodes)
	})
	r.POST("/interpret", Interpret)

	r.POST("/do", func(c *gin.Context) {

		r := InterpretRequest{}

		var input map[string]interface{}
		data, err := ioutil.ReadAll(c.Request.Body)
		fmt.Println(string(data))

		err = json.Unmarshal(data, &input)

		if err != nil {
			c.JSON(400, gin.H{
				"Error": err.Error(),
			})
			return
		}

		fmt.Println("input", input)

		results, err := interpreter.Eval([]interpreter.ASTNode{r.Body})
		if err != nil {
			log.Info().Str("url", r.URL).Msgf("do failed: %s", err)
			c.JSON(400, gin.H{
				"Error": err.Error(),
			})
			return
		}
		c.JSON(200, results)
	})

	return
}

func (s *Server) register(node string) (err error) {

	addNodeRequest := AddNodeRequest{
		Name: s.Name,
		URL:  s.URL,
	}

	data, err := json.Marshal(addNodeRequest)
	if err != nil {
		return
	}

	url := fmt.Sprintf("%s/nodes", node)

	var resp *http.Response
	resp, err = http.Post(url, "application/json", bytes.NewBuffer(data))
	if err != nil {
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		log.Info().Str("url", url).Msg("failed to register with node")
		return
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Info().Str("url", url).Msg("failed to read body")
		return
	}
	result := map[string]string{}
	err = json.Unmarshal(body, &result)
	if err != nil {
		log.Info().Str("url", url).Msg("failed to parse response")
		return
	}

	for k, v := range result {
		nodes[k] = v
	}

	return
}

func (s *Server) Start() {

	duration := s.CheckAlive
	if duration == 0 {
		duration = time.Minute
	}
	go checkAlive(duration, s.signals)
	s.Router.Run()
}

func (s *Server) Stop() {
	s.signals <- true
}

func checkAlive(duration time.Duration, signalCh chan bool) {

	for {
		select {
		case <-signalCh:
			return
		case <-time.After(duration):
		}

		total := 0
		for name, url := range nodes {
			log.Debug().Str(name, url).Msg("checking... ")
			_, err := http.Get(fmt.Sprintf("%s/ping", url))
			if err != nil {
				log.Info().Str(name, url).Msgf("failed healthcheck removing from list of nodes: %s", err)
				delete(nodes, name)
				continue
			}
			total = total + 1
		}

		log.Info().Int("count", total).Msg("check alive finished")

	}
}
