package proxy

import (
	"fmt"
	log "gitlab.zfu.fintek.local/zfu/deeplink/core/logger"
	"gitlab.zfu.fintek.local/zfu/deeplink/models"
	"io/ioutil"
	"net/http"
	"os"
	"encoding/json"
)

type PokemonProxy struct {
	logger log.ILogger
}

func NewPokemonProxy(logger log.ILogger) PokemonProxy {
	return PokemonProxy{logger}
}

func (p *PokemonProxy) GetPokemons() (models.Response, error) {
	response, err := http.Get("http://pokeapi.co/api/v2/pokedex/kanto/")
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		//log.Fatal(err)
	}

	var responseObject models.Response
	json.Unmarshal(responseData, &responseObject)

	return responseObject, err
}

