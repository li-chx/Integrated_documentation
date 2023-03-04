package mysql

import "fmt"

func init() {
	dbDSN := fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=%s&parseTime=%v&loc=%s")
}
