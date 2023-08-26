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

func CmdSetNamaServiceCost() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "setNamaServiceCost [amount]",
		Short: "To Set New Price Of Nama",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			newCost := args[0]
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			newCostInUint64, err := strconv.ParseUint(newCost, 10, 64)
			if err != nil {
				return err
			}
			msg := types.NewMsgSetNamaServiceCost(
				clientCtx.GetFromAddress().String(),
				newCostInUint64,
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
