package cli

import (
	"strconv"

	"nama/x/nama/types"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdReserveNama() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "reserveNama [newNama]",
		Short: "To Set New Admin Of Nama",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			newNama := args[0]
			price := args[1]
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			priceInUint, err := strconv.ParseUint(price, 10, 64)
			if err != nil {
				return err
			}
			msg := types.NewReserveNama(
				clientCtx.GetFromAddress().String(),
				newNama,
				priceInUint,
			)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
