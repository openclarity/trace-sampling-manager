package main

import (
	"os"
	"os/signal"
	"syscall"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/urfave/cli"
	"github.com/Portshift/go-utils/healthz"
	"github.com/Portshift/go-utils/k8s"
	logutils "github.com/Portshift/go-utils/log"

	"github.com/apiclarity/trace-sampling-manager/manager/pkg/config"
	"github.com/apiclarity/trace-sampling-manager/manager/pkg/manager"
)

const defaultChanSize = 100

func run(c *cli.Context) {
	logutils.InitLogs(c, os.Stdout)
	conf := config.LoadConfig()

	errChan := make(chan struct{}, defaultChanSize)

	healthServer := healthz.NewHealthServer(conf.HealthCheckAddress)
	healthServer.Start()
	healthServer.SetIsReady(false)

	clientset, _, err := k8s.CreateK8sClientset(nil /*InClusterConfig*/, k8s.KubeOptions{})
	if err != nil {
		log.Fatalf("Failed to create K8s clientset: %v", err)
	}

	m, err := manager.Create(clientset, &manager.Config{
		RestServerPort: conf.RestServerPort,
		GRPCServerPort: conf.GRPCServerPort,
	})
	if err != nil {
		log.Fatalf("Failed to create manager: %v", err)
	}

	if err := m.Start(errChan); err != nil {
		log.Fatalf("Failed to start manager: %v", err)
	}
	defer m.Stop()

	healthServer.SetIsReady(true)

	// Wait for deactivation
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	select {
	case <-errChan:
		log.Errorf("Received an error - shutting down")
	case s := <-sig:
		log.Warningf("Received a termination signal: %v", s)
	}
}

func main() {
	viper.SetDefault(config.RestServerPort, "9990")
	viper.SetDefault(config.GRPCServerPort, "9991")
	viper.SetDefault(config.HealthCheckAddress, ":8080")
	viper.AutomaticEnv()

	app := cli.NewApp()
	app.Usage = ""
	app.Name = "Trace Sampling Manager"
	app.Version = "1.0.0"

	runCommand := cli.Command{
		Name:   "run",
		Usage:  "Starts Trace Sampling Manager",
		Action: run,
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  logutils.LogLevelFlag,
				Value: logutils.LogLevelDefaultValue,
				Usage: logutils.LogLevelFlagUsage,
			},
		},
	}
	runCommand.UsageText = runCommand.Name

	app.Commands = []cli.Command{
		runCommand,
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
