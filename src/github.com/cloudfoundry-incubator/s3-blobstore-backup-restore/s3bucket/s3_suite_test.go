package s3bucket_test

import (
	"os"

	. "github.com/cloudfoundry-incubator/backup-and-restore-sdk-release-system-tests"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

var (
	S3Endpoint                   = MustHaveEnvOrBeEmpty("S3_ENDPOINT")
	LiveRegion                   = MustHaveEnv("S3_LIVE_REGION")
	BackupRegion                 = MustHaveEnv("S3_BACKUP_REGION")
	AccessKey                    = MustHaveEnv("S3_ACCESS_KEY_ID")
	SecretKey                    = MustHaveEnv("S3_SECRET_ACCESS_KEY")
	PreExistingBigFileBucketName = MustHaveEnv("S3_BIG_FILE_BUCKET")
	EmptyBucketName              = MustHaveEnv("S3_EMPTY_BUCKET")
)

func TestS3(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "S3 Suite")
}

func MustHaveEnvOrBeEmpty(keyname string) string {
	val, exist := os.LookupEnv(keyname)
	if !exist {
		panic("Need " + keyname + " for the test")
	}
	return val
}
