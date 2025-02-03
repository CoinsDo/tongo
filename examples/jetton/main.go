package main

import (
	"context"
	"encoding/hex"
	"fmt"
	"log"
	"math/big"
	"time"

	"github.com/tonkeeper/tongo"
	"github.com/tonkeeper/tongo/contract/jetton"
	"github.com/tonkeeper/tongo/liteapi"
	"github.com/tonkeeper/tongo/wallet"
)

func main() {

	recipientAddr := tongo.MustParseAddress("0:68bdf41999c991f765c354262816683ecf4feb1ac0e61960a7ee8e413279e8f8")

	//parseAccountID := tongo.MustParseAccountID("kQCCRfKECr4uLG7Ui5ekTvIzuJeqTPjH6yMJT9wG6chsUh1H").ToRaw()
	recipientAddr = tongo.MustParseAddress(tongo.MustParseAccountID("0QCPsjO7PufrayJxR6Orh5pTq_DnNxsWks48Zou6eBFSnpaZ").ToRaw())

	client, err := liteapi.NewClientWithDefaultTestnet()
	if err != nil {
		log.Fatalf("Unable to create tongo client: %v", err)
	}

	//pk, _ := base64.StdEncoding.DecodeString("OyAWIb4FeP1bY1VhALWrU2JN9/8O1Kv8kWZ0WfXXpOM=")
	//privateKey := ed25519.NewKeyFromSeed(pk)
	//w, err := wallet.New(privateKey, wallet.V4R2, client)
	//if err != nil {
	//	panic("unable to create wallet")
	//}

	w, err := wallet.DefaultWalletFromSeed("best journey rifle scheme bamboo daring finish life have puzzle verb wagon double pencil plate parent canoe soup stable salon drift elephant border hero", client)
	if err != nil {
		return
	}
	privateKey, err := wallet.SeedToPrivateKey("best journey rifle scheme bamboo daring finish life have puzzle verb wagon double pencil plate parent canoe soup stable salon drift elephant border hero")
	if err != nil {
		return
	}

	fmt.Println(hex.EncodeToString(privateKey))
	return

	log.Printf("Jetton from address: %s", w.GetAddress().ToHuman(false, true))
	log.Printf("Jetton TO address: %s", recipientAddr.ID.ToHuman(false, true))

	master := tongo.MustParseAccountID("kQAiboDEv_qRrcEdrYdwbVLNOXBHwShFbtKGbQVJ2OKxY_Di")
	//jettonWalletAddr := tongo.MustParseAccountID("kQCCDbrrw2gVMK9ViK8flHeqZrR8_7032QgfgnbhxVwS_bRc")

	log.Printf("Jetton  address: %s", master.ToHuman(false, true))
	j := jetton.New(master, client)
	b, err := j.GetBalance(context.Background(), w.GetAddress())
	if err != nil {
		log.Fatalf("Unable to get jetton wallet balance: %v", err)
	}
	d, err := j.GetDecimals(context.Background())
	if err != nil {
		log.Fatalf("Get decimals error: %v", err)
	}
	jw, err := j.GetJettonWallet(context.Background(), w.GetAddress())
	if err != nil {
		log.Fatalf("Get jetton wallet error: %v", err)
	}

	log.Printf("Jetton balance: %v", b)
	log.Printf("Jetton decimals: %v", d)
	log.Printf("Jetton wallet owner address: %v", w.GetAddress().String())
	log.Printf("Jetton wallet address: %v", jw.String())
	log.Printf("Jetton wallet address: %v", jw.ToHuman(true, true))
	//return

	amount := big.NewInt(1000000)
	if amount.Cmp(b) == 1 {
		log.Fatalf("%v jettons needed, but only %v on balance", amount, b)
	}
	getAddress := w.GetAddress()
	jettonTransfer := jetton.TransferMessage{
		//Jetton:              j,
		JettonAmount:        amount,
		Sender:              w.GetAddress(),
		Destination:         recipientAddr.ID,
		ResponseDestination: &getAddress,
		AttachedTon:         50_000_000,
		ForwardTonAmount:    1_000,
	}
	//return
	resp, err := w.SendV2(context.Background(), 0, jettonTransfer)
	if err != nil {
		log.Fatalf("Unable to send transfer message: %v", err)
	}
	fmt.Println(resp.Base64())
	time.Sleep(time.Second * 15)
	b, err = j.GetBalance(context.Background(), w.GetAddress())
	if err != nil {
		log.Fatalf("Unable to get jetton wallet balance: %v", err)
	}
	log.Printf("New Jetton balance: %v", b)
}
