package context
import "reflect"

type ApplicationContext struct {
	// Definition Register Bean
	definitions map[reflect.Type]BeanDefinition
	// Caching Created Bean
	singletons map[reflect.Type]any
	// circular dependency detect flag
	creating map[reflect.Type]bool
}

func NewApplicationContext() *ApplicationContext{
	return &ApplicationContext{
		definitions : map[reflect.Type]BeanDefinition{},
		singletons: map[reflect.Type]any{},
		creating: map[reflect.Type]bool{},
	}
}

/*
Register Bean Method 
*/
func (ctx *ApplicationContext) RegisterBean(beanType reflect.Type, factory func(ctx *ApplicationContext) any) {
	ctx.definitions[beanType] = BeanDefinition{BeanType: beanType, Factory: factory}
}

/*
Get Bean Method

1. 이미 생성된 Bean이면 -> 캐시를 반환합니다.
2. 지금 생성중이면 순환 의존성이므로, 즉시 실패합니다.
3. 정의가 없으면 실패합니다.
4. Factory를 실행하고 Bean을 반환받습니다.
5. Bean 결과를 Singleton 캐시에 저장합니다.
6. 생성된 Bean을 반환합니다.
*/
func (ctx *ApplicationContext) GetBean(beanType reflect.Type) any {
	// value, exsists(t/f) := map_var[key]
	bean, ok := ctx.singletons[beanType]
	if ok {
		return bean
	}

	if ctx.creating[beanType] {
		panic("circular dependency detected")
	}

	def, ok := ctx.definitions[beanType]
	if !ok {
		panic("bean not registered")
	}

	ctx.creating[beanType] = true
	bean = def.Factory(ctx)
	ctx.creating[beanType] = false

	ctx.singletons[beanType] = bean
	return bean
}