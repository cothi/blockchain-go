package unit

import (
	"fmt"
	"testing"
)

func TestGetInstance(t *testing.T) {

	GetInstance().add()
	fmt.Println(GetInstance().getCount())

	GetInstance().add()
	fmt.Println(GetInstance().getCount())

	GetInstance().add()
	fmt.Println(GetInstance().getCount())

}
