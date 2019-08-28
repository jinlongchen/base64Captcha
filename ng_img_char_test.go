package base64Captcha

import (
	"os"
	"testing"
)

func TestEngineCharCreate(t *testing.T) {
	//os.RemoveAll("/Users/jinlongchen/work/go/src/github.com/jinlongchen/base64Captcha/test/*")
	tc := "/Users/jinlongchen/work/go/src/github.com/jinlongchen/base64Captcha/test/" // ioutil.TempDir("", "png")
	t.Log(tc)
	defer os.Remove(tc)
	config := ConfigCharacter{
		Height:             40,
		Width:              160,
		Mode:               CaptchaModeAlphabet, // CaptchaModeNumber:数字,CaptchaModeAlphabet:字母,CaptchaModeArithmetic:算术,CaptchaModeNumberAlphabet:数字字母混合.
		ComplexOfNoiseText: CaptchaComplexLower,
		ComplexOfNoiseDot:  CaptchaComplexLower,
		IsUseSimpleFont:    false,
		IsShowHollowLine:   false,
		IsShowNoiseDot:     true,
		IsShowNoiseText:    false,
		IsShowSlimeLine:    true,
		IsShowSineLine:     true,
		CaptchaLen:         4,
	}
	//configC.Mode = CaptchaModeAlphabet //i % 4
	//configC.IsUseSimpleFont = true
	//configC.CaptchaLen = 4
	//configC.IsShowSineLine = true
	//configC.IsShowNoiseDot = true
	//configC.Width = 160
	//configC.Height  =40

	im := EngineCharCreate(config)
	//fileName := strings.Trim(im.Content, "/+-+=?")
	err := CaptchaWriteToFile(im, tc, "123456", "png")
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
