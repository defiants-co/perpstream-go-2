package synthetix

import (
	"github.com/defiants-co/perpstream-go-2/clients/common"
	"github.com/defiants-co/perpstream-go-2/clients/common/models"
)

type SynthetixClient struct{}

func (s *SynthetixClient) GetLeaderboard(limit int) ([]models.User, error) {
	// TODO: Implement this method
	return nil, nil
}

func (s *SynthetixClient) FetchPositions(userId string) ([]models.Future, error) {
	// TODO: Implement this method
	return nil, nil
}

func (s *SynthetixClient) StartPositionStream(userId string, callback common.StreamPositionsCallback[models.Future], waitSeconds float64) error {
	// TODO: Implement this method
	return nil
}

func (s *SynthetixClient) CancelPositionStream(userId string) error {
	// TODO: Implement this method
	return nil
}

func (s *SynthetixClient) StartOrderFlowStream(callback common.StreamOrdersCallback[models.Future]) error {
	// TODO: Implement this method
	return nil
}

func (s *SynthetixClient) CancelOrderFlowStream() error {
	// TODO: Implement this method
	return nil
}
