package main

import (
	"errors"
	"flag"
	"fmt"
)

const Default_cluster_size =  5

type ClusterMode int

const (
	ClusterOff ClusterMode = iota
	ClusterOn
	ClusterError
)

var clusterStatus = map[ClusterMode]string{
	ClusterOff: "off",
	ClusterOn: "on",
	ClusterError: "error",
}

func main() {

	// cluster size can be set using a flag
	cluster_size := flag.Int("numserve", Default_cluster_size, "an int")
	flag.Parse()

	currcluster := cluster{ClusterOff, cluster_size}

	// Run the cluster

}


type Mode struct {


}

type cluster struct {
	mode Mode
	members	[]string
}


while len

leader := ""
if leader == "" {
	// run election
	// return leader
} else {

}

func (currcluster *cluster) init(numserve int) -> cluster {

}

func cluster run() {
	while leader {
		// handle inputs
	} 
	
	// hold election, recursively call Cluster.run()

}