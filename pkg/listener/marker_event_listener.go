package listener

//import (
//	"ethsyncer/types"
//	ethCommon "github.com/ethereum/go-ethereum/common"
//	"github.com/ethereum/go-ethereum/ethclient"
//	log "github.com/sirupsen/logrus"
//)
//
//type MarkerEventListener struct {
//	Web3Client      *ethclient.Client
//	Address         string
//	OnOrderExecuted func(order *types.VERC20MarketVERC20OrderExecuted)
//	OnOrderCanceled func(order *types.VERC20MarketVERC20OrderCanceled)
//	contract        *types.VERC20Market
//}
//
//func NewMarkerEventListener(
//	web3Client *ethclient.Client,
//	address string,
//	OnOrderExecuted func(order *types.VERC20MarketVERC20OrderExecuted),
//	OnOrderCanceled func(order *types.VERC20MarketVERC20OrderCanceled),
//) (*MarkerEventListener, error) {
//	contractAddress := ethCommon.HexToAddress(address)
//	contract, err := types.NewVERC20Market(contractAddress, web3Client)
//	if err != nil {
//		return nil, err
//	}
//
//	return &MarkerEventListener{
//		Web3Client: web3Client,
//		Address:    address,
//		contract:   contract,
//		OnOrderExecuted: OnOrderExecuted,
//		OnOrderCanceled: OnOrderCanceled,
//	}, nil
//}
//
//func (mel *MarkerEventListener) ListenOrderExecuted() {
//	eventChannel := make(chan *types.VERC20MarketVERC20OrderExecuted)
//	subscription, err := mel.contract.WatchVERC20OrderExecuted(nil, eventChannel)
//	if err != nil {
//		log.Fatal(err)
//	}
//	for {
//		select {
//		case err := <-subscription.Err():
//			log.Error(err)
//		case event := <-eventChannel:
//			mel.OnOrderExecuted(event)
//		}
//	}
//}
//
//func (mel *MarkerEventListener) ListenOrderCanceled() {
//	eventChannel := make(chan *types.VERC20MarketVERC20OrderCanceled)
//	subscription, err := mel.contract.WatchVERC20OrderCanceled(nil, eventChannel)
//	if err != nil {
//		log.Fatal(err)
//	}
//	for {
//		select {
//		case err := <-subscription.Err():
//			log.Error(err)
//		case event := <-eventChannel:
//			mel.OnOrderCanceled(event)
//		}
//	}
//}
