package cli

import (
	"fmt"

	"github.com/LMF709268224/titan-vps/api/types"
	"github.com/urfave/cli/v2"
)

// BasisCMDs Basis cmd
var BasisCMDs = []*cli.Command{
	WithCategory("order", orderCmds),
	WithCategory("user", userCmds),
	describeRegionsCmd,
	mintCmd,
}

var orderCmds = &cli.Command{
	Name:  "order",
	Usage: "Manage order",
	Subcommands: []*cli.Command{
		createOrderCmd,
		cancelOrderCmd,
		paymentCompletedCmd,
	},
}

var userCmds = &cli.Command{
	Name:  "user",
	Usage: "Manage user",
	Subcommands: []*cli.Command{
		signCodeCmd,
		loginCmd,
		logoutCmd,
		getBalanceCmd,
		rechargeCmd,
	},
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
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:  "user",
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

		user := cctx.String("user")

		address, err := api.CreateOrder(ctx, types.CreateOrderReq{Vps: types.CreateInstanceReq{
			RegionId:     "cn-qingdao",
			ImageId:      "aliyun_2_1903_x64_20G_alibase_20230704.vhd",
			PeriodUnit:   "week",
			Period:       1,
			InstanceType: "ecs.t5-lc1m1.small",
			DryRun:       true,
		}, User: user})
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
			Name:  "order_id",
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

		orderID := cctx.String("order_id")

		return api.CancelOrder(ctx, orderID)
	},
}

var paymentCompletedCmd = &cli.Command{
	Name:  "pc",
	Usage: "payment completed",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:  "order_id",
			Usage: "node type: edge 1, update 6",
			Value: "",
		},
		&cli.StringFlag{
			Name:  "tr_id",
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

		orderID := cctx.String("order_id")
		trID := cctx.String("tr_id")

		str, err := api.PaymentCompleted(ctx, types.PaymentCompletedReq{OrderID: orderID, TransactionID: trID})
		if err != nil {
			return err
		}

		fmt.Println(str)
		return nil
	},
}

var getBalanceCmd = &cli.Command{
	Name:  "gb",
	Usage: "get balance",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:  "address",
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

		address := cctx.String("address")

		str, err := api.GetBalance(ctx, address)
		if err != nil {
			return err
		}

		fmt.Println(str)
		return nil
	},
}

var rechargeCmd = &cli.Command{
	Name:  "recharge",
	Usage: "recharge",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:  "address",
			Usage: "node type: edge 1, update 6",
			Value: "",
		},
		&cli.StringFlag{
			Name:  "ra",
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

		address := cctx.String("address")
		rechargeAddr := cctx.String("ra")

		str, err := api.Recharge(ctx, address, rechargeAddr)
		if err != nil {
			return err
		}

		fmt.Println(str)
		return nil
	},
}

var mintCmd = &cli.Command{
	Name:  "mint",
	Usage: "mint token",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:  "address",
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

		address := cctx.String("address")

		info, err := api.MintToken(ctx, address)
		if err != nil {
			return err
		}

		fmt.Println(info)
		return nil
	},
}

var signCodeCmd = &cli.Command{
	Name:  "sc",
	Usage: "user get code",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:  "",
			Usage: "",
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

		userId := cctx.String("user_id")

		str, err := api.SignCode(ctx, userId)
		if err != nil {
			return err
		}

		fmt.Println(str)
		return nil
	},
}

var loginCmd = &cli.Command{
	Name:  "login",
	Usage: "user login",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:  "",
			Usage: "",
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

		userId := cctx.String("user_id")
		publicKey := cctx.String("public_key")
		signature := cctx.String("signature")

		str, err := api.Login(ctx, &types.UserReq{
			UserId:    userId,
			PublicKey: publicKey,
			Signature: signature,
		})
		if err != nil {
			return err
		}

		fmt.Println(str)
		return nil
	},
}

var logoutCmd = &cli.Command{
	Name:  "logout",
	Usage: "user logout",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:  "",
			Usage: "",
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

		userId := cctx.String("user_id")
		token := cctx.String("token")

		err = api.Logout(ctx, &types.UserReq{
			UserId: userId,
			Token:  token,
		})
		if err != nil {
			return err
		}
		return nil
	},
}
