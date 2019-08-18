package base64Captcha

import (
	"os"
	"strings"
	"testing"
)

func TestEngineCharCreate(t *testing.T) {
	//os.RemoveAll("/Users/jinlongchen/work/go/src/github.com/jinlongchen/base64Captcha/test/*")
	tc := "/Users/jinlongchen/work/go/src/github.com/jinlongchen/base64Captcha/test/" // ioutil.TempDir("", "png")
	t.Log(tc)
	defer os.Remove(tc)
	//for i := 0; i < 16; i++ {
	//	configC.Mode = CaptchaModeAlphabet //i % 4
	//	boooo := i%2 == 0
	//	configC.IsUseSimpleFont = false
	//	configC.IsShowSlimeLine = boooo
	//	configC.IsShowNoiseText = boooo
	//	configC.IsShowHollowLine = boooo
	//	configC.IsShowSineLine = boooo
	//	configC.IsShowNoiseDot = boooo
	//
	//	im := EngineCharCreate(configC)
	//	fileName := strings.Trim(im.Content, "/+-+=?")
	//	err := CaptchaWriteToFile(im, tc, fileName, "png")
	//	if err != nil {
	//		t.Error(err)
	//	}
	//}
	configC.Mode = CaptchaModeAlphabet //i % 4
	//boooo := false
	configC.IsUseSimpleFont = true
	//configC.IsShowSlimeLine = true
	//configC.IsShowNoiseText = true
	//configC.IsShowHollowLine = true
	configC.IsShowSineLine = true
	configC.IsShowNoiseDot = true

	im := EngineCharCreate(configC)
	fileName := strings.Trim(im.Content, "/+-+=?")
	err := CaptchaWriteToFile(im, tc, fileName, "png")
	if err != nil {
		t.Error(err)
	}
}
func TestMath(t *testing.T) {
	for i := 0; i < 100; i++ {
		q, r := randArithmetic()
		t.Log(q, "--->", r)
	}
}
