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

func CmdAcquireNama() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "acquireNama [namaId]",
		Short: "To Acquire Nama",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			namaId := args[0]
			price := args[1]
			namaIdUint64, err := strconv.ParseUint(namaId, 10, 64)
			if err != nil {
				return err
			}
			priceUint64, err := strconv.ParseUint(price, 10, 64)
			if err != nil {
				return err
			}
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			msg := types.NewAcquireNama(
				clientCtx.GetFromAddress().String(),
				namaIdUint64,
				priceUint64,
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
