package cli

import (
	"fmt"

	"github.com/LMF709268224/titan-vps/api/types"
	"github.com/urfave/cli/v2"
)

// BasisCMDs Basis cmd
var BasisCMDs = []*cli.Command{
	describeRegionsCmd,
	createOrderCmd,
	cancelOrderCmd,
	paymentCompletedCmd,
}

var describeRegionsCmd = &cli.Command{
	Name:  "dr",
	Usage: "describe regions",

	Before: func(cctx *cli.Context) error {
		return nil
	},
	Action: func(cctx *cli.Context) error {
		ctx := ReqContext(cctx)

		api, closer, err := GetBasisAPI(cctx)
		if err != nil {
			return err
		}

		defer closer()

		list, err := api.DescribeRegions(ctx)
		if err != nil {
			return err
		}

		fmt.Println(list)
		return nil
	},
}

var createOrderCmd = &cli.Command{
	Name:  "create",
	Usage: "create order",

	Before: func(cctx *cli.Context) error {
		return nil
	},
	Action: func(cctx *cli.Context) error {
		ctx := ReqContext(cctx)

		api, closer, err := GetBasisAPI(cctx)
		if err != nil {
			return err
		}

		defer closer()

		address, err := api.CreateOrder(ctx, types.CreateOrderReq{Vps: "123456"})
		if err != nil {
			return err
		}

		fmt.Println(address)
		return nil
	},
}

var cancelOrderCmd = &cli.Command{
	Name:  "cancel",
	Usage: "cancel order",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:  "hash",
			Usage: "node type: edge 1, update 6",
			Value: "",
		},
	},
	Action: func(cctx *cli.Context) error {
		ctx := ReqContext(cctx)

		api, closer, err := GetBasisAPI(cctx)
		if err != nil {
			return err
		}
		defer closer()

		hash := cctx.String("hash")

		return api.CancelOrder(ctx, hash)
	},
}

var paymentCompletedCmd = &cli.Command{
	Name:  "pc",
	Usage: "payment completed",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:  "hash",
			Usage: "node type: edge 1, update 6",
			Value: "",
		},
	},
	Action: func(cctx *cli.Context) error {
		ctx := ReqContext(cctx)

		api, closer, err := GetBasisAPI(cctx)
		if err != nil {
			return err
		}

		defer closer()

		hash := cctx.String("hash")

		str, err := api.PaymentCompleted(ctx, types.PaymentCompletedReq{Hash: hash})
		if err != nil {
			return err
		}

		fmt.Println(str)
		return nil
	},
}
