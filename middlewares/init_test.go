package middlewares

import (
	"os"
	"testing"

	"github.com/Amazeful/Amazeful-Backend/util"
	"github.com/Amazeful/Amazeful-Backend/util/mocks"
)

var mockRepo *mocks.Repository

func TestMain(m *testing.M) {

	mockRepo = new(mocks.Repository)
	util.SetMockRepoGetter(mockRepo)

	os.Exit(m.Run())
}
