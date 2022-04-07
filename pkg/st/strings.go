package st

import (
	"regexp"
	"strings"
)

func RemoveHTML(str string) string {
	if len(str) > 0 {
		//删除脚本
		reg := regexp.MustCompile(`([\r\n])[\s]+`)
		str = reg.ReplaceAllString(str, "")
		reg = regexp.MustCompile(`<script[^>]*?>.*?</script>`)
		str = reg.ReplaceAllString(str, "")
		//删除HTML
		reg = regexp.MustCompile(`<(.[^>]*)>`)
		str = reg.ReplaceAllString(str, "")
		reg = regexp.MustCompile(`([\r\n])[\s]+`)
		str = reg.ReplaceAllString(str, "")
		reg = regexp.MustCompile(`-->`)
		str = reg.ReplaceAllString(str, "")
		reg = regexp.MustCompile(`<!--.*`)
		str = reg.ReplaceAllString(str, "")
		reg = regexp.MustCompile(`&(quot|#34);`)
		str = reg.ReplaceAllString(str, "")
		reg = regexp.MustCompile(`&(amp|#38);`)
		str = reg.ReplaceAllString(str, "")
		reg = regexp.MustCompile(`&(lt|#60);`)
		str = reg.ReplaceAllString(str, "")
		reg = regexp.MustCompile(`&(gt|#62);`)
		str = reg.ReplaceAllString(str, "")
		reg = regexp.MustCompile(`&(nbsp|#160);`)
		str = reg.ReplaceAllString(str, "")
		reg = regexp.MustCompile(`&(iexcl|#161);`)
		str = reg.ReplaceAllString(str, "")
		reg = regexp.MustCompile(`&(cent|#162);`)
		str = reg.ReplaceAllString(str, "")
		reg = regexp.MustCompile(`&(pound|#163);`)
		str = reg.ReplaceAllString(str, "")
		reg = regexp.MustCompile(`&(copy|#169);`)
		str = reg.ReplaceAllString(str, "")
		reg = regexp.MustCompile(`&#(\d+);`)
		str = reg.ReplaceAllString(str, "")

		str = strings.ReplaceAll(str, "<", "")
		str = strings.ReplaceAll(str, ">", "")
		str = strings.ReplaceAll(str, "\n", "")
		str = strings.ReplaceAll(str, " ", "")
		str = strings.ReplaceAll(str, "　", "")

		return str
	}
	return ""
}
