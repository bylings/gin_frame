package models

import (
	"fmt"
	"testing"
)

func TestRegions(t *testing.T) {
	where:=map[string]interface{}{
		"id":1,
	}
	// Unscoped 关闭软删除
	var regions Regions
	DB().Debug().Where(where).Unscoped().Find(&regions)
	fmt.Println("reg   --  ",regions)
}
