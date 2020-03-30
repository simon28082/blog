package infrastructure

import (
    "github.com/go-playground/locales/zh"
    "github.com/go-playground/universal-translator"
    "github.com/go-playground/validator/v10"
    translations "github.com/go-playground/validator/v10/translations/zh"
    "reflect"
    "regexp"
    //"time"
)

var (
    //app           *firmeve2.Firmeve
    validate      *validator.Validate
    validateTrans ut.Translator
)

//func Make(abstract interface{}, params ...interface{}) interface{} {
//    return app.Make(abstract, params...)
//}
//
//func DB() *database.DB {
//    return app.Get(`db`).(*database.DB)
//}
//
//func Cache() *cache.Cache {
//    return app.Get(`cache`).(*cache.Cache)
//}
//
//func Connection() *gorm.DB {
//    return app.Get(`db.connection`).(*gorm.DB)
//}
//
//func Dispatch(name string, params event.InParams) {
//    app.Get(`event`).(event.IDispatcher).Dispatch(name, params)
//}

//func FormDecoder()  {
//    parser.FormDecoder.RegisterCustomTypeFunc(func(strings []string) (i interface{}, e error) {
//        local,_ := time.LoadLocation("Local")
//        return time.ParseInLocation("2006-01-02 15:04:05", strings[0],local)
//    },time.Time{})
//}

// Singleton
func Validate(structData interface{}) {
    if validate == nil {
        validate = validator.New()
        uni := ut.New(zh.New())
        if validateTrans == nil {
            validateTrans, _ = uni.GetTranslator("zh")
        }

        // RegisterTagNameFunc
        validate.RegisterTagNameFunc(func(field reflect.StructField) string {
            return field.Tag.Get(`alias`)
        })

        // RegisterValidation
        err := validate.RegisterValidation(`mobile`, func(fl validator.FieldLevel) bool {
            return regexp.MustCompile(`^1[\d]{10}$`).MatchString(fl.Field().String())
        })
        if err != nil {
            panic(err)
        }

        // RegisterTranslation
        err = validate.RegisterTranslation(`mobile`, validateTrans, func(ut ut.Translator) error {
            return ut.Add(`mobile`, `手机号格式不正确`, true)
        }, func(ut ut.Translator, fe validator.FieldError) string {
            t, _ := ut.T(`mobile`, fe.Field())
            return t
        })
        if err != nil {
            panic(err)
        }

        err = translations.RegisterDefaultTranslations(validate, validateTrans)
        if err != nil {
            panic(err)
        }
    }

    err := validate.Struct(structData)

    if err != nil {
        panic(err)
        //panic(http.Error422(err.(validator.ValidationErrors)[0].Translate(validateTrans)))
    }
}
