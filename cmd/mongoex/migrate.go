package mongoex

import (
        "mongoex/pkg/migrator"
        "github.com/spf13/cobra"
        "fmt"
        "strings"
)

var migrateCmd = &cobra.Command{
    Use:   "migrate",
    Aliases: []string{"mig"},
    Short:  "Migrates data from a source MongoDB to destination MongoDB",
    //Args:  cobra.ExactArgs(1),
    RunE: func(cmd *cobra.Command, args []string) error {
	source, _               := cmd.Flags().GetString("source")
        destination, _          := cmd.Flags().GetString("destination")
        collections, _          := cmd.Flags().GetString("collections")
        databaseSource, _       := cmd.Flags().GetString("dbsrc")
        databaseDestination, _  := cmd.Flags().GetString("dbdest")
 
        coll := strings.Split(collections, ",")

        fmt.Println("\nSource: ", source)
        fmt.Println("Destination: ", destination)
        // only if collections is specified
        if collections != "" {
                fmt.Println("Collections: ", collections)
        }  
        fmt.Println("Database Source: ", databaseSource)
        fmt.Println("Database Destination: ", databaseDestination)

        if collections != "" {
                res := migrator.MigrateCollections(source, destination, databaseSource, databaseDestination, coll)
                if res {
                        fmt.Println("Selective collection copy has been successful!")
                }
                return nil
        }

        res := migrator.MigrateAll(source, destination, databaseSource, databaseDestination)
        if res {
                fmt.Println("Database copy is successful")
        }
        return nil
    },
}

func init() {
        rootCmd.AddCommand(migrateCmd)
        migrateCmd.PersistentFlags().StringP("source", "s", "mongodb://localhost:27017", "Source MongoDB Host. Example: (\"mongodb://username:password@localhost:27017\")")
        migrateCmd.PersistentFlags().StringP("destination", "d", "mongodb://localhost:27017", "Destination MongoDB Host: Example: (\"mongodb://username:password@localhost:27017\")")
        migrateCmd.PersistentFlags().StringP("dbsrc", "", "", "Source Database (optional), else will migrate all if this is omitted")
        migrateCmd.PersistentFlags().StringP("dbdest", "", "", "Destination Database (optional), else will migrate all if this is omitted")
        migrateCmd.PersistentFlags().StringP("collections", "c", "", "List of collections (optional), else will migrate all if this is omitted")
}
