package types_test

import (
	"fmt"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	"github.com/umee-network/umee/app"
	"github.com/umee-network/umee/x/leverage/types"
)

func TestAddressFromKey(t *testing.T) {
	address := sdk.AccAddress([]byte("addr________________"))
	key := types.CreateLoanKey(address, app.BondDenom)
	expectedAddress := types.AddressFromKey(key, types.KeyPrefixLoanToken)

	require.Equal(t, address, expectedAddress)

	address = sdk.AccAddress([]byte("anotherAddr________________"))
	key = types.CreateCollateralAmountKeyNoDenom(address)
	expectedAddress = types.AddressFromKey(key, types.KeyPrefixLoanToken)

	require.Equal(t, address, expectedAddress)
}

func TestDenomFromKeyWithAddress(t *testing.T) {
	address := sdk.AccAddress([]byte("addr________________"))
	denom := app.BondDenom
	key := types.CreateLoanKey(address, denom)
	expectedDenom := types.DenomFromKeyWithAddress(key, types.KeyPrefixLoanToken)

	require.Equal(t, denom, expectedDenom)

	uDenom := fmt.Sprintf("u%s", denom)
	key = types.CreateCollateralSettingKey(address, uDenom)
	expectedDenom = types.DenomFromKeyWithAddress(key, types.KeyPrefixCollateralSetting)

	require.Equal(t, uDenom, expectedDenom)
}