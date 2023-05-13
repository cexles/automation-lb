package service

import (
	"context"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/event"
	"github.com/rs/zerolog/log"
	"time"
)

type ChainService struct {
	ctx         context.Context
	client      *ethclient.Client
	moduleName  string
	headerCh    chan *types.Header
	SubChannels []chan *types.Header
}

func NewChainService(ctx context.Context, client *ethclient.Client, listeners []func(chan *types.Header)) *ChainService {
	s := &ChainService{
		ctx:         ctx,
		client:      client,
		headerCh:    make(chan *types.Header),
		SubChannels: make([]chan *types.Header, 0, len(listeners)),
		moduleName:  "service.ChainService",
	}
	for range listeners {
		s.SubChannels = append(s.SubChannels, make(chan *types.Header))
	}
	return s
}

func (s *ChainService) StartBlockListener() error {
	_ = event.Resubscribe(time.Minute*2, func(ctx context.Context) (event.Subscription, error) {
		log.Info().Str("module", s.moduleName).
			Str("func", "StartBlockListener").
			Msg("subscribed to NewHead event")
		e, err := s.client.SubscribeNewHead(ctx, s.headerCh)
		if err != nil {
			log.Err(err).Str("module", s.moduleName).
				Str("func", "StartBlockListener").
				Msg("failed to reconnect to node. Trying to reconnect...")
			return nil, err
		}
		return e, nil
	})

	return nil
}

func (s *ChainService) StartBlockBroadcaster() {
	defer func() {
		for _, ch := range s.SubChannels {
			close(ch)
		}
	}()

	for {
		select {
		case <-s.ctx.Done():
			return
		case newBlock, ok := <-s.headerCh:
			if !ok {
				return
			}
			log.Debug().Str("module", s.moduleName).
				Str("func", "StartBlockBroadcaster").
				Fields(map[string]any{
					"height":    newBlock.Number,
					"timestamp": newBlock.Time,
					"hash":      newBlock.Hash(),
				}).
				Msg("new block height")
			for _, ch := range s.SubChannels {
				ch <- newBlock
			}
		}
	}
}
