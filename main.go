package main

import (
	"fmt"
	"github.com/jacklove/go-gin-example/models"
	"github.com/jacklove/go-gin-example/pkg/logging"
	"github.com/jacklove/go-gin-example/pkg/setting"
	"github.com/jacklove/go-gin-example/routers"
	"net/http"
)

func main()  {
	setting.SetUp()
	models.SetUp()
	logging.SetUp()


	//endless.DefaultReadTimeOut = setting.ServerSetting.ReadTimeout
	//endless.DefaultWriteTimeOut = setting.ServerSetting.WriteTimeout
	//endless.DefaultMaxHeaderBytes = 1 << 20
	//endPoint := fmt.Sprintf(":%d", setting.ServerSetting.HttpPort)
	//
	//server := endless.NewServer(endPoint,routers.InitRouter())
	//server.BeforeBegin = func(add string) {
	//	log.Printf("Actual pid is %d", syscall.Getpid())
	//}
	//
	//err := server.ListenAndServe()
	//if err != nil {
	//	log.Printf("Server err: %v", err)
	//}

	//router := routers.InitRouter()
	//
	//s := &http.Server{
	//	Addr:           fmt.Sprintf(":%d", setting.HTTPPort),
	//	Handler:        router,
	//	ReadTimeout:    setting.ReadTimeout,
	//	WriteTimeout:   setting.WriteTimeout,
	//	MaxHeaderBytes: 1 << 20,
	//}
	//
	//go func() {
	//	if err := s.ListenAndServe(); err != nil {
	//		log.Printf("Listen: %s\n", err)
	//	}
	//}()
	//
	//quit := make(chan os.Signal)
	//signal.Notify(quit, os.Interrupt)
	//<- quit
	//log.Println("Shutdown Server ...")
	//
	//ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
	//defer cancel()
	//if err := s.Shutdown(ctx); err != nil {
	//	log.Fatal("Server Shutdown:", err)
	//}
	//
	//log.Println("Server exiting")




	router := routers.InitRouter()
	s := &http.Server{
		Addr: fmt.Sprintf(":%d",setting.ServerSetting.HttpPort),
		Handler: router,
		ReadTimeout: setting.ServerSetting.ReadTimeout,
		WriteTimeout: setting.ServerSetting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	s.ListenAndServe()


}
