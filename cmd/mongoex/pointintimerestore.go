package mongoex

import (
        "mongoex/pkg/atlas"
        "github.com/spf13/cobra"
	_ "fmt"
	"strconv"
	"time"
)

var pointInTimeCmd = &cobra.Command{
    Use:   "pointintime",
    Aliases: []string{"pitr"},
    Short:  "Create a temporary cluster and do a Point-In-Time-Recovery to this cluster",
    //Args:  cobra.ExactArgs(1),
    RunE: func(cmd *cobra.Command, args []string) error {
        projectName, _               := cmd.Flags().GetString("proj")
	clusterName, _               := cmd.Flags().GetString("clusterName")
	pointInTimeSeconds, _        := cmd.Flags().GetString("time")
	sourceClusterName, _         := cmd.Flags().GetString("sourceClusterName")
	targetProjectID, _           := cmd.Flags().GetString("targetProject")

	// convert diskSize from string to float which is required
        var timeSeconds int64
	if pointInTimeSeconds == "" {
		timeSeconds = time.Now().Unix()
	} else {
		timeSeconds, _ = strconv.ParseInt(pointInTimeSeconds, 10, 64)
	}
	atlas.PointInTimeRestore(projectName, clusterName, timeSeconds, sourceClusterName, targetProjectID)
        return nil
    },
}

func init() {
        tempClusterCmd.AddCommand(pointInTimeCmd)
        pointInTimeCmd.Flags().StringP("proj", "p", "", "MongoDB Project Name")
        pointInTimeCmd.Flags().StringP("clusterName", "c", "", "Name for temporary cluster")
        pointInTimeCmd.Flags().StringP("sourceClusterName", "", "", "Source MongoDB Cluster Name")
        pointInTimeCmd.Flags().StringP("targetProject", "", "", "Target Project ID")
	pointInTimeCmd.Flags().StringP("time", "", "", "Point-in-time since epoch (defaults to current time)")

        pointInTimeCmd.MarkFlagRequired("proj")
        pointInTimeCmd.MarkFlagRequired("clusterName")
        // pointInTimeCmd.MarkFlagRequired("time") // time is optional, if omitted, it will defaul tto current time
        pointInTimeCmd.MarkFlagRequired("sourceClusterName")
        pointInTimeCmd.MarkFlagRequired("targetProject")
}
