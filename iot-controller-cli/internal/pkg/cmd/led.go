package cmd

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/raykrishardi/iot-controller-cli/internal/pkg/iot"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	grpcHost          = os.Getenv("GRPC_HOST")
	grpcPort          = os.Getenv("GRPC_PORT")
	defaultCtxTimeout = 3 * time.Second
)

type ledCmdFlag struct {
	pin    int
	active bool
}

var lcf ledCmdFlag

var ledCmd = &cobra.Command{
	Use:   "led",
	Short: "LED Controller",
	Run:   func(cmd *cobra.Command, args []string) { runLEDCmd(cmd, args) },
}

func init() {
	rootCmd.AddCommand(ledCmd)

	ledCmd.Flags().IntVar(&lcf.pin, "pin", 0, "led pin")
	ledCmd.MarkFlagRequired("pin")
	ledCmd.Flags().BoolVar(&lcf.active, "active", false, "led status (on or off)")
	ledCmd.MarkFlagRequired("active")
}

func runLEDCmd(cmd *cobra.Command, args []string) {
	grpcURI := fmt.Sprintf("%s:%s", grpcHost, grpcPort)

	conn, err := grpc.Dial(grpcURI, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
	if err != nil {
		log.Fatalf("unable to connect to grpc server %s: %v", grpcURI, err)
	}
	defer conn.Close()

	c := iot.NewIotServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), defaultCtxTimeout)
	defer cancel()

	req := iot.IotRequest{
		Led: &iot.Led{
			Pin:    int32(lcf.pin),
			Active: lcf.active,
		},
	}

	res, err := c.ProcessIot(ctx, &req)
	if err != nil {
		log.Fatalf("unable to process IoT request %v: %v", req.Led, err)
	}

	fmt.Println(res.Result)
}
