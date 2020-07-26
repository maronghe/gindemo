package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/en_CA"
	"github.com/go-playground/locales/fr"
	"github.com/go-playground/locales/nl"
	"github.com/go-playground/locales/zh"
	"github.com/go-playground/locales/zh_Hant_TW"
	"github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_transaltions "github.com/go-playground/validator/v10/translations/en"
	zh_transaltions "github.com/go-playground/validator/v10/translations/zh"

	"gindemo/global"
)

// TODO don't know how to use
func Translations() gin.HandlerFunc {
	return func(c *gin.Context) {
		e := en.New()
		uni := ut.New(e, e, en_CA.New(), nl.New(), fr.New(), zh.New(), zh_Hant_TW.New())
		locale := c.GetHeader("locale")
		trans, found := uni.GetTranslator(locale)
		global.Logger.Debug(found, trans)
		v, ok := binding.Validator.Engine().(*validator.Validate)
		var err error
		if ok {
			switch locale {
			case "zh":
				err = zh_transaltions.RegisterDefaultTranslations(v, trans)
			case "en":
				err = en_transaltions.RegisterDefaultTranslations(v, trans)
			default:
				err = zh_transaltions.RegisterDefaultTranslations(v, trans)
			}

			global.Logger.Debugf("Translations error %+v", err)

			c.Set("trans", trans)
		}

		c.Next()
	}

}
