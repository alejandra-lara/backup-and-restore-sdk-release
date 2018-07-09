package gcp_test

import (
	"testing"

	"time"

	"os/exec"

	. "github.com/cloudfoundry-incubator/backup-and-restore-sdk-release-system-tests"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gexec"
	. "github.com/onsi/gomega/gexec"
)

func TestGcp(t *testing.T) {
	RegisterFailHandler(Fail)
	SetDefaultEventuallyTimeout(15 * time.Minute)
	RunSpecs(t, "GCP System Tests Suite")
}

var _ = BeforeSuite(func() {
	MustRunSuccessfully("gcloud", "auth", "activate-service-account",
		"--key-file", MustHaveEnv("GCP_SERVICE_KEY_PATH"))
	MustRunSuccessfully("gcloud", "config", "set", "project", MustHaveEnv("GCP_PROJECT_NAME"))
})

func MustRunSuccessfully(command string, args ...string) {
	cmd := exec.Command(command, args...)

	session, err := gexec.Start(cmd, GinkgoWriter, GinkgoWriter)
	Expect(err).NotTo(HaveOccurred())
	Eventually(session).Should(Exit(0))
}