package service

/*
import (
	"fmt"
	log "gitlab.zfu.fintek.local/zfu/deeplink/core/logger"
	"gitlab.zfu.fintek.local/zfu/deeplink/models"
	"gitlab.zfu.fintek.local/zfu/deeplink/proxy"
	"time"
)

type PokemonService struct {
	logger log.ILogger
}

func NewPokemonService(logger log.ILogger) PokemonService {
	return PokemonService{logger}
}

type pokeChanType struct {
	order        int
	startTime    time.Time
	pokeResponse models.Response
	err          error
}

var response []models.Response

func (p *PokemonService) GetPokemons() ([]models.Response, error) {
	response = make([]models.Response, 0)

	chanError := make(chan error, 1)
	chanPokemon := make(chan pokeChanType)

	p.getPokes(0, chanError, chanPokemon)
	p.getPokes(1, chanError, chanPokemon)

	waitingCount := 2
	for i := 0; i < waitingCount; i++ {
		select {
		case err := <-chanError:
			return response, err
		case poke := <-chanPokemon:
			p.setPokes(poke)
		}
	}
	return response, nil
}

func (p *PokemonService) getPokes(order int, chanError chan<- error, chanPoke chan<- pokeChanType) {
	fmt.Println("asdasd frt: %+v", order)
	go func(order int) {
		proxy := proxy.NewPokemonProxy(p.logger)
		response, err := proxy.GetPokemons()

		if err != nil {
			chanError <- err
		}

		chanPoke <- pokeChanType{order: order, startTime: time.Now(), pokeResponse: response, err: nil}
	}(order)
}

func (p *PokemonService) setPokes(poke pokeChanType) {
	response = append(response, poke.pokeResponse)
}
*/
