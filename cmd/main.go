package main

import (
	"github.com/eden-framework/context"
	"github.com/eden-framework/eden-framework/pkg/application"
	"github.com/eden-framework/sqlx/migration"
	"github.com/robotic-framework/srv-robotic-platform/internal/routers/admin"
	"github.com/robotic-framework/srv-robotic-platform/internal/routers/client"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"github.com/robotic-framework/srv-robotic-platform/internal/databases"
	"github.com/robotic-framework/srv-robotic-platform/internal/global"
)

var cmdMigrationDryRun bool

func main() {
	app := application.NewApplication(runner,
		application.WithConfig(&global.Config),
		application.WithApollo(&global.ApolloConfig),
		application.WithConfig(&databases.Config),
		application.WithConfig(&global.ProxyConfig))

	cmdMigrate := &cobra.Command{
		Use: "migrate",
		Run: func(cmd *cobra.Command, args []string) {
			migrate(&migration.MigrationOpts{
				DryRun: cmdMigrationDryRun,
			})
		},
	}
	cmdMigrate.Flags().BoolVarP(&cmdMigrationDryRun, "dry", "d", false, "migrate --dry")
	app.AddCommand(cmdMigrate)

	app.Start()
}

func runner(ctx *context.WaitStopContext) error {
	logrus.SetLevel(global.Config.LogLevel)
	go global.ProxyConfig.Server.Start(ctx)
	go global.Config.GRPCServer.Serve(ctx, client.Router)
	return global.Config.HTTPServer.Serve(ctx, admin.Router)
}

func migrate(opts *migration.MigrationOpts) {
	if err := migration.Migrate(global.Config.MasterDB, opts); err != nil {
		panic(err)
	}
}
