package binder

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/schema"
	"net/http"
	"reflect"
)

var decoder = schema.NewDecoder()
var validate = validator.New()

func init() {
	decoder.IgnoreUnknownKeys(true)
	decoder.ZeroEmpty(true)

	// 默认使用 form 标签
	decoder.SetAliasTag("form")

	decoder.RegisterConverter(map[string]string{}, func(s string) reflect.Value {
		var m map[string]string
		_ = json.Unmarshal([]byte(s), &m)
		return reflect.ValueOf(m)
	})
}

// 支持多标签映射：uri / form / query
func decodeWithMultiTag(obj interface{}, values map[string][]string) error {
	// 复制一份结构体，使用反射找到标签
	v := reflect.ValueOf(obj)
	if v.Kind() != reflect.Ptr || v.IsNil() {
		return fmt.Errorf("obj must be a non-nil pointer")
	}
	v = v.Elem()
	t := v.Type()

	formValues := map[string][]string{}

	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		if !f.IsExported() {
			continue
		}

		// 优先级：uri > form
		tag := f.Tag.Get("uri")
		if tag == "" {
			tag = f.Tag.Get("form")
		}
		if tag == "" {
			continue
		}

		if val, ok := values[tag]; ok {
			formValues[tag] = val
		}
	}

	return decoder.Decode(obj, formValues)
}

func bindPath(c *gin.Context, obj interface{}) error {
	values := map[string][]string{}
	for _, p := range c.Params {
		values[p.Key] = []string{p.Value}
	}
	return decodeWithMultiTag(obj, values)
}

func bindQuery(c *gin.Context, obj interface{}) error {
	return decodeWithMultiTag(obj, c.Request.URL.Query())
}

func bindForm(c *gin.Context, obj interface{}) error {
	if err := c.Request.ParseForm(); err != nil {
		return err
	}
	return decodeWithMultiTag(obj, c.Request.Form)
}

// Bind 通用 Bind
func Bind(c *gin.Context, obj interface{}) error {
	ct := c.ContentType()

	// 优先绑定 Path 参数（RESTful）
	_ = bindPath(c, obj)

	// 绑定 Query 参数
	_ = bindQuery(c, obj)

	switch ct {
	case "application/json":
		if err := c.ShouldBindJSON(obj); err != nil {
			return err
		}
	case "application/x-www-form-urlencoded", "multipart/form-data":
		if err := bindForm(c, obj); err != nil {
			return err
		}
	default:
		if c.Request.Method == http.MethodGet {
			// GET 请求默认只解析 Query
			return nil
		}
		return fmt.Errorf("unsupported Content-Type: %s", ct)
	}

	return nil
}

// BindAndValidate 绑定并校验
func BindAndValidate(c *gin.Context, obj interface{}) error {
	err := Bind(c, obj) // 你已有的 Bind 负责填充字段
	if err != nil {
		return err
	}
	// 调用 validator 验证结构体
	return validate.Struct(obj)
}
