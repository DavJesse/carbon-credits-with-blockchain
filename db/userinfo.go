package db

import (
    "log"
    "time"

    "github.com/couchbase/gocb/v2"
)

var Cluster *gocb.Cluster
var Collection *gocb.Collection

func InitDB() {
    var err error
    Cluster, err = gocb.Connect("couchbase://localhost", gocb.ClusterOptions{
        Username: "yourusername",
        Password: "yourpassword",
    })
    if err != nil {
        log.Fatalf("Error connecting to Couchbase: %v", err)
    }

    bucket := Cluster.Bucket("yourbucket")

    // Wait until the bucket is ready
    err = bucket.WaitUntilReady(5 * time.Second, nil)
    if err != nil {
        log.Fatalf("Error connecting to bucket: %v", err)
    }

    Collection = bucket.DefaultCollection()
}
