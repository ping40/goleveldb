package testutil

import (
	"fmt"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func RunSuite(t GinkgoTestingT, name string) {

	fmt.Println("in RunSuite ", time.Now())
	RunDefer()

	//SynchronizedBeforeSuite must be passed two functions. The first must return []byte and the second
	// must accept []byte. When running with multiple nodes the first function is only run on node 1.
	// When this function completes, all nodes (including node 1) proceed to run the second function and
	// will receive the data returned by the first function.
	SynchronizedBeforeSuite(func() []byte {
		RunDefer("setup")
		return nil
	}, func(data []byte) {})

	//With SynchronizedAfterSuite the first function is run on all nodes (including node 1). The second function
	// is only run on node 1. Moreover, the second function is only run when all other nodes have finished running.
	// This is important, since node 1 is responsible for setting up and tearing down the singleton resources
	// it must wait for the other nodes to end before tearing down the resources they depend on.
	SynchronizedAfterSuite(func() {
		RunDefer("teardown")
	}, func() {})

	RegisterFailHandler(Fail)
	RunSpecs(t, name) // tells Ginkgo to start the test suite.
}
