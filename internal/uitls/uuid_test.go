package uitls

import (
	"fmt"
	"testing"

	"github.com/google/uuid"
)

func TestName(t *testing.T) {
	fmt.Print(uuid.New().String())
}
