package ioc

import (
	"reflect"
)

// Global Bean instance container
var beanCache = make(map[string]interface{})

// Bean
// Declare struct instances to ioc control
type Bean interface {
	// CreateBean
	// Initializes bean
	CreateBean() Bean
}

// Inject Given an object instance and complete all property injections
func Inject(v interface{}) {
	proxy := reflect.ValueOf(v)
	if proxy.Kind() != reflect.Ptr {
		panic("injected bean instance must be a pointer")
	}
	proxy = proxy.Elem()
	typeProxy := reflect.TypeOf(v).Elem()
	for i := 0; i < typeProxy.NumField(); i++ {
		filed := typeProxy.Field(i)
		if !filed.IsExported() {
			continue
		}
		if filed.Type.Kind() == reflect.Interface {
			for _, b := range beanCache {
				btp := reflect.TypeOf(b)
				if btp.Implements(filed.Type) {
					primary := filed.Tag.Get("@autowired")
					if primary == "" || btp.Elem().String() == primary {
						proxy.Field(i).Set(reflect.ValueOf(b))
						break
					}
					continue
				}
			}
		} else {
			if filed.Type.Kind() != reflect.Ptr {
				continue
			}
			bean, ok := beanCache[filed.Type.Elem().String()]
			if !ok {
				if filed.Type.Implements(reflect.TypeOf((*Bean)(nil)).Elem()) {
					filedValue := proxy.Field(i)
					bean = filedValue.Interface().(Bean).CreateBean()
					if bean != nil {
						Inject(bean)
						beanCache[filed.Type.Elem().String()] = bean
						filedValue.Set(reflect.ValueOf(bean))
					}
				}
				continue
			}
			proxy.Field(i).Set(reflect.ValueOf(bean))
		}
	}
}

// SetBeans
// Manually set multiple instances to enter the IOC container. The bean name defaults to package name.Structure name.
// such as：hello.Message.
// If the same package name and structure appear, the existing bean will be overwritten
func SetBeans(beans ...interface{}) {
	for _, bean := range beans {
		typeof := reflect.TypeOf(bean)
		if typeof.Kind() != reflect.Ptr {
			panic("bean must be a pointer")
		}
		beanCache[typeof.Elem().String()] = bean
	}
}

// GetBean query according to the specified structure
func GetBean(beanStruct interface{}) interface{} {
	typeOf := reflect.TypeOf(beanStruct)
	if typeOf.Kind() == reflect.Ptr {
		typeOf = typeOf.Elem()
	}
	bean, ok := beanCache[typeOf.String()]
	if ok {
		return bean
	}
	return nil
}

// GetBeanByName query bean by name
func GetBeanByName(beanName string) interface{} {
	bean, ok := beanCache[beanName]
	if ok {
		return bean
	}
	return nil
}
