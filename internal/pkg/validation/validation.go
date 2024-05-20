package validation

import (
	"context"
	"github.com/costa92/k8s-krm-go/internal/pkg/middleware/validate"
	"github.com/google/wire"
	"k8s.io/klog/v2"
	"reflect"
	"strings"
)

type validator struct {
	registry map[string]reflect.Value
}

var ProviderSet = wire.NewSet(New, wire.Bind(new(validate.IValidator), new(*validator)))

// any is an alias for interface{}
func New(cv any) *validator {
	return &validator{registry: GetValidateFuncs(cv)}
}

func (vd *validator) Validate(ctx context.Context, rq any) error {
	m, ok := vd.registry[reflect.TypeOf(rq).Elem().Name()]
	if !ok {
		return nil
	}

	val := m.Call([]reflect.Value{reflect.ValueOf(ctx), reflect.ValueOf(rq)})
	if !val[0].IsNil() {
		return val[0].Interface().(error)
	}

	return nil
}

func GetValidateFuncs(cv any) map[string]reflect.Value {
	funcs := make(map[string]reflect.Value)
	typeOf := reflect.TypeOf(cv)
	valueOf := reflect.ValueOf(cv)
	for i := 0; i < typeOf.NumMethod(); i++ {
		m := typeOf.Method(i)
		val := valueOf.MethodByName(m.Name)
		if !val.IsValid() {
			continue
		}

		if !strings.HasPrefix(m.Name, "Validate") {
			continue
		}

		typ := val.Type()
		if typ.NumIn() != 2 || typ.NumOut() != 1 {
			continue
		}

		if typ.In(0) != reflect.TypeOf((*context.Context)(nil)).Elem() {
			continue
		}

		if typ.In(1).Kind() != reflect.Pointer {
			continue
		}

		vName := typ.In(1).Elem().Name()
		if m.Name != ("Validate" + vName) {
			continue
		}

		if typ.Out(0) != reflect.TypeOf((*error)(nil)).Elem() {
			continue
		}

		klog.V(4).InfoS("Register validator", "validator", vName)
		funcs[vName] = val
	}

	return funcs
}
