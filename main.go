package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"

	"github.com/miyazi777/go-todo-app2/config"
)

func main() {
	if err := run(context.Background()); err != nil {
		log.Printf("failed to terminate server: %v", err)
		os.Exit(1)
	}
}

func run(ctx context.Context) error {
	cfg, err := config.New()
	if err != nil {
		return err
	}
	l, err := net.Listen("tcp", fmt.Sprintf(":%d", cfg.Port))
	if err != nil {
		log.Fatalf("failed to listen port %d: %v", cfg.Port, err)
	}

	url := fmt.Sprintf("http://%s", l.Addr().String())
	log.Printf("start with: %v", url)
	mux := NewMux()
	s := NewServer(l, mux)
	return s.Run(ctx)

	// // SIGTERMを受け取ったら、現在の処理中の処理の終了を待ってから終了する
	// ctx, stop := signal.NotifyContext(ctx, os.Interrupt, syscall.SIGTERM)
	// defer stop()

	// cfg, err := config.New()
	// if err != nil {
	// 	return err
	// }
	// l, err := net.Listen("tcp", fmt.Sprintf(":%d", cfg.Port))
	// if err != nil {
	// 	log.Fatalf("failed to listen port %d: %v", cfg.Port, err)
	// }
	// url := fmt.Sprintf("http://%s", l.Addr().String())
	// log.Printf("start with: %v", url)

	// // 起動するサーバとハンドラーを用意する
	// s := &http.Server{
	// 	// Addr: ":18080",	引数で受け取ったListenerを使うので、Addrは指定しない
	// 	Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	// 		time.Sleep(5 * time.Second)
	// 		fmt.Fprintf(w, "Hello, %s!", r.URL.Path[1:])
	// 	}),
	// }
	// eg, ctx := errgroup.WithContext(ctx)

	// // 別ゴルーチンでHTTPサーバを起動する
	// eg.Go(func() error {
	// 	// http.ErrServerClosedはhttp.Server.Shutdown()が正常に終了したことを示すので異常ではない
	// 	if err := s.Serve(l); err != nil && err != http.ErrServerClosed {
	// 		log.Printf("failed to close: %+v", err)
	// 		return err
	// 	}
	// 	return nil
	// })

	// // チャネルからの通知（終了通知）を待機する
	// <-ctx.Done()

	// // http.Serverのshutdownを呼び出すのでグレースフルシャットダウンを開始する
	// if err := s.Shutdown(context.Background()); err != nil {
	// 	log.Printf("failed to shutdown: %+v", err)
	// }

	// // Goメソッドで起動した別ゴルーチンの終了を待つ
	// return eg.Wait()
}
