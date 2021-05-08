package main

import (
	ecc "exchange/crypto/ECC"
	"exchange/utils"
	"fmt"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/urfave/cli"
	"net/http"
	"os"
	"time"
)

var (
	cmp       = 1
	app       = cli.NewApp()
	baseFlags = []cli.Flag{
		utils.PortFlag,
		utils.KeyFlag,
		utils.EthAccountFlag,
		utils.EthKeyFlag,
	}
	ethaccount    string
	usrpub        = ecc.PublicKey{}
	publisherpub  = ecc.PublicKey{}
	publisherpriv = ecc.PrivateKey{}
	regulatorpub  = ecc.PublicKey{}
	cm_and_r      = ecc.Commitment{}
	elgamal_info  = ecc.CypherText{}
	elgamal_r     = ecc.CypherText{}
	signature     = ecc.Signature{}
	ea            string
	ek            string
	gk            string
)

func init() {
	app.Name = "exchange"
	app.Usage = "user exchange from there"
	app.Action = exchange
	app.Flags = append(app.Flags, baseFlags...)

}

func main() {
	if err := app.Run(os.Args); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func exchange(ctx *cli.Context) {
	gk = ctx.String("generatekey")
	ea = ctx.String("ethaccount")
	ek = ctx.String("ethkey")
	ethaccount = ctx.String("ethaccount")
	publisherpub, publisherpriv, _ = utils.GenerateKey(gk)
	regulatorpub = utils.SetRegulator()
	//if utils.UnlockAccount(ea, ek) == true {
	startNetwork(ctx)
	//} else {
	//	fmt.Println("erro unlock exchanger eth_account")
	//	return
	//}
}

func startNetwork(ctx *cli.Context) error {
	e := echo.New()
	port := ctx.String("port")

	// Root level middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.POST("/buy", buy)
	e.GET("/pubpub", pubpub)

	e.Logger.Fatal(e.Start(":" + port))
	return nil
}

func buy(c echo.Context) error {
	u := new(utils.Purchase)
	if err := c.Bind(u); err != nil {
		return err
	}
	if u.G1 == "" || u.G2 == "" || u.P == "" || u.H == "" || u.Amount == "" {
		return c.JSON(http.StatusCreated, "err params lack")
	}
	if utils.Verify(u.H) == false {
		return c.JSON(http.StatusCreated, "error publickey, please check again or registe now")
	} else {
		utils.UnlockAccount(ea, ek)
		usrpub = utils.CreateUsrPub(u.G1, u.G2, u.P, u.H)
		//cm_and_r = utils.CreateCM_v(regulatorpub, u.Amount)
		//elgamal_info = utils.CreateElgamalInfo(regulatorpub, u.Amount, u.H)
		elgamal_info, cm_and_r = utils.CreateDE_CM(regulatorpub, u.Amount)
		elgamal_r = utils.CreateElgamalR(usrpub, cm_and_r.R)
		fmt.Println("you want this one",utils.Byteto0xstring(cm_and_r.R))
		signature = utils.CreateSign(publisherpriv, u.Amount)
		//sendTranscation
		if succ, hash := utils.SendTransaction(elgamal_info, elgamal_r, signature, cm_and_r, ethaccount); succ == true {
			result := utils.Toreceipt(cm_and_r.Commitment, elgamal_r.C1, elgamal_r.C2, hash)
			return c.JSON(http.StatusOK, result)
		} else {
			return c.JSON(http.StatusCreated, "err send transaction")
		}
	}
}

func pubpub(c echo.Context) error {
	if cmp == 1 {
		go unlock()
	}
	return c.JSON(http.StatusCreated, publisherpub)
}
func unlock() {
	time.Sleep(3 * time.Second)
	utils.UnlockAccount(ea, ek)
	cmp = 0
}
