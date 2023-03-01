package resources

import (
	"os"
	"testing"

	"github.com/selefra/selefra-provider-sdk/test_helper"
)

func Test_Provider(t *testing.T) {
	provider := GetSelefraProvider()
	// split := strings.Split(os.Getenv("SELEFRA_TEST_TABLES"), ",")
	os.Setenv("AUTH0_DOMAIN", "dev-pdzw1hcstv43ci6d.us.auth0.com")
	os.Setenv("AUTH0_CLIENT_ID", "Uv2prhCe436y5E9Utzq8RtTCGqnwVbjL")
	os.Setenv("AUTH0_CLIENT_SECRET", "GK_LLekK_CLxphRigTl0inMh9hNp10t5zfHDiVLAUQcwENRtKDrAskKtx3BJfGlt")

	os.Setenv("SELEFRA_DATABASE_DSN", "host=127.0.0.1 user=postgres password=postgres port=5432 dbname=postgres sslmode=disable")
	test_helper.RunProviderPullTables(provider, "log: info", "./", "auth0_organization")
}
