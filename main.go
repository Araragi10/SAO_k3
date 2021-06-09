// Rudeus Telegram Bot Project
// Copyright (C) 2021 wotoTeam, ALiwoto
// This file is subject to the terms and conditions defined in
// file 'LICENSE', which is part of the source code.

package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/gotd/td/telegram"
	"github.com/gotd/td/tg"

	"github.com/Araragi10/SAO_k3/wotoPacks/appSettings"
	"github.com/Araragi10/SAO_k3/wotoPacks/interfaces"
	"github.com/Araragi10/SAO_k3/wotoPacks/wotoActions"
	"github.com/Araragi10/SAO_k3/wotoPacks/wotoDB"
	ws "github.com/Araragi10/SAO_k3/wotoPacks/wotoSecurity/wotoStrings"
	wv "github.com/Araragi10/SAO_k3/wotoPacks/wotoValues"
)

// Rudeus01 Telegram Bot Project's main entry.
func main() {
	//sao, err := wotoNeko.GetTenorRandomTinywebm("Sword Art Online")
	//if err != nil {
	//	log.Fatal(err)
	//}

	//log.Println(sao)
	//cmd := exec.Command("curl https://google.com")
	//b, err := cmd.Output()
	//if err != nil {
	//	log.Println(err)
	//}
	//log.Println(string(b))
	//log.Println(string(bb))

	log.Println("Copy Right ALiwoto; All Right Reserved.")
	log.Println("=======================================")
	log.Println("Now starting settings~")

	//ra, err := wotoNeko.GetRandomTGif("hentai")
	//if err != nil {
	//	log.Fatal(err)
	//}

	//log.Println(ra)
	//er := flag.CommandLine.ErrorHandling()
	//fs := flag.NewFlagSet("/pat --txt=hello", er)
	//rf := fs.Lookup("txt")
	//log.Println(rf.DefValue)
	//test12 := strings.Replace("Hello How are you?", " ", "_", 1)
	//log.Println(test12)

	//myText := "/tr --txt \"Hello!=How Are : You??:huha -- hahaha\" --message = \"So, Can I ask- why-- are \\\"You\\\" :here??\" --txt : \"Oh yeah, I love you too.\" --grade 7890"
	//myText := "/ud hello world : = : --pv"
	//pA, err := pTools.ParseArg(myText)
	//if err != nil {
	//	log.Fatal(err)
	//}

	//t := pA.GetIndexFlag(0).GetType()
	//bf := pA.IsEmpty()
	//log.Println(pA)

	go Listen2()

	settings := appSettings.GetSettings()

	token := os.Getenv(wv.TOKEN_KEY)
	if ws.IsEmpty(&token) {
		log.Fatal(wv.INVALID_API)
	}

	apiId := os.Getenv("API_ID")
	if ws.IsEmpty(&token) {
		log.Fatal(wv.INVALID_API)
	}
	apiInt, err := strconv.Atoi(apiId)
	if err != nil {
		log.Fatal(err)
	}

	ApiHash := os.Getenv("API_HASH")
	if ws.IsEmpty(&token) {
		log.Fatal(wv.INVALID_API)
	}

	//ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	//defer cancel()

	ctx := context.Background()

	dclient := telegram.NewClient(apiInt, ApiHash, telegram.Options{})

	dclient.Run(ctx, func(ctx context.Context) error {

		log.Println("Starting dClient Run function.")
		_, err := dclient.Auth().Bot(ctx, token)

		if err != nil {
			log.Fatal(err)
		}

		g := tg.NewClient(dclient)

		if g == nil {
			log.Fatal("gClient is nil. we can't continue our work like this.")
		}

		//p := uploader.NewUploader(g)
		//iClass, err := p.FromURL(ctx, "https://myhost/mygif.gif")

		settings.SetGClient(g, &ctx)
		log.Println("Done in Run function of dClient")
		runApp(settings)

		return nil
	})

}

func Listen2() {
	port := os.Getenv(wv.APP_PORT)
	if ws.IsEmpty(&port) {
		// couldn't find port value in env, so
		// let us continue without it, there is no
		// problem at all.
		log.Println(wv.PORT_ERROR)
		return
	}
	router := gin.New()
	router.Use(gin.Logger())

	router.GET(wv.GET_SLASH, func(c *gin.Context) {
		c.String(http.StatusOK, wv.ALREADY_RUNNING)
	})

	router.Run(wv.HTTP_ADDRESS + port)
}

func runApp(settings interfaces.WSettings) {
	if !appSettings.IsRunning() {
		appSettings.App_enter()
	} else {
		log.Fatal(wv.INVALID_ENGINE)
		return
	}
	wv.DebugMode = true
	token := os.Getenv(wv.TOKEN_KEY)
	if ws.IsEmpty(&token) {
		log.Fatal(wv.INVALID_API)
	}

	log.Println("Before SetTobt")
	settings.SetTObt(token)
	wotoDB.GenerateClients(settings)
	wotoActions.RunBot(settings)

}
