package keeper

import (
	"fmt"

	"cosmossdk.io/core/store"
	"cosmossdk.io/log"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/charles08200/cosmos/x/cosmos/types"
)

type (
	Keeper struct {
		cdc          codec.BinaryCodec
		storeService store.KVStoreService
		logger       log.Logger

		// the address capable of executing a MsgUpdateParams message. Typically, this
		// should be the x/gov module account.
		authority string
	}
)

func NewKeeper(
	cdc codec.BinaryCodec,
	storeService store.KVStoreService,
	logger log.Logger,
	authority string,

) Keeper {
	if _, err := sdk.AccAddressFromBech32(authority); err != nil {
		panic(fmt.Sprintf("invalid authority address: %s", authority))
	}

	return Keeper{
		cdc:          cdc,
		storeService: storeService,
		authority:    authority,
		logger:       logger,
	}
}

// GetAuthority returns the module's authority.
func (k Keeper) GetAuthority() string {
	return k.authority
}

// Logger returns a module-specific logger.
func (k Keeper) Logger() log.Logger {
	return k.logger.With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

// HandleBitcoinMessage processes a message sent from the Bitcoin relayer.
func (k Keeper) HandleBitcoinMessage(ctx sdk.Context, bitcoinMessage string) error {
	// Log the incoming Bitcoin message
	k.Logger().Info("Handling Bitcoin message", "message", bitcoinMessage)

	// Example of storing the message in the module's state
	store := k.storeService.OpenKVStore(ctx)
	messageKey := []byte("last_bitcoin_message") // Key to store the latest Bitcoin message
	err := store.Set(messageKey, []byte(bitcoinMessage))
	if err != nil {
		return fmt.Errorf("failed to store Bitcoin message: %w", err)
	}

	// Placeholder for further processing logic
	k.Logger().Info("Bitcoin message stored successfully", "message", bitcoinMessage)

	return nil
}

// GetLastBitcoinMessage retrieves the last message received from Bitcoin.
func (k Keeper) GetLastBitcoinMessage(ctx sdk.Context) (string, error) {
	store := k.storeService.OpenKVStore(ctx)
	messageKey := []byte("last_bitcoin_message") // Key used to store the message

	messageBytes, err := store.Get(messageKey)
	if err != nil {
		return "", fmt.Errorf("failed to retrieve Bitcoin message: %w", err)
	}
	if messageBytes == nil {
		return "", nil // No message found
	}

	return string(messageBytes), nil
}
