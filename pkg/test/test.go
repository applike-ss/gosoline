package test

import (
	"fmt"
	"github.com/ory/dockertest"
	"github.com/ory/dockertest/docker"
	"github.com/pkg/errors"
	"github.com/twinj/uuid"
	"log"
	"sync"
)

var err error
var wait sync.WaitGroup
var dockerPool *dockertest.Pool
var network *docker.Network
var dockerResources []*dockertest.Resource

func init() {
	dockerPool, err = dockertest.NewPool("")
	dockerResources = make([]*dockertest.Resource, 0)

	if err != nil {
		log.Fatalf("Could not connect to docker: %s", err)
	}
}

func logErr(err error, msg string) {
	Shutdown()
	log.Println(msg)
	log.Fatal(err)
}

func Boot() error {
	config := readConfig()

	testNetworkName := "test_network_" + uuid.NewV4().String()
	err := createTestNetwork(testNetworkName)

	if err != nil {
		return errors.Wrapf(err, "failed to create test network %s", testNetworkName)
	}

	for name, mockConfig := range config.Mocks {
		bootComponent(name, mockConfig)
	}

	wait.Wait()

	log.Println("test environment up and running")
	fmt.Println()

	return nil
}

func bootComponent(name string, mockConfig configInput) {
	component := mockConfig["component"]

	switch component {
	case "dynamodb":
		runDynamoDb(name, mockConfig)
	case "cloudwatch":
		runCloudwatchContainer(name, mockConfig)
	case "sns":
		runSnsContainer(name, mockConfig)
	case "sqs":
		runSqsContainer(name, mockConfig)
	case "elasticsearch":
		runElasticsearch(name, mockConfig)
	case "mysql":
		runMysql(name, mockConfig)
	case "redis":
		runRedis(name, mockConfig)
	case "wiremock":
		runWiremock(name, mockConfig)
	default:
		err := fmt.Errorf("unknown component '%s'", component)
		logErr(err, err.Error())
	}
}

func Shutdown() {
	for _, res := range dockerResources {
		if err := dockerPool.Purge(res); err != nil {
			log.Fatalf("Could not purge resource: %s", err)
		}
	}

	err := removeTestNetwork(network)

	if err != nil {
		log.Println("failed to remove test network! " + err.Error())
	}
}
