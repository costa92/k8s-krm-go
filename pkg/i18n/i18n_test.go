package i18n

import (
	"fmt"
	"golang.org/x/text/language"
	"testing"
)

func TestNew(t *testing.T) {
	i := New()
	// 1. add dir
	i.Add("./locales")

	// 2. add file
	i.Add("./locales/en.yml")
	i.Add("./locales/zh.yml")

	// 3. add embed fs
	//i.AddFS(fs)

	fmt.Println(i.T("common.hello"))
	fmt.Println(i.Select(language.Chinese).T("common.hello"))
}
