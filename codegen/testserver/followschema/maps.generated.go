// Code generated by github.com/trevor-scheer/gqlgen, DO NOT EDIT.

package followschema

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"sync/atomic"

	"github.com/trevor-scheer/gqlgen/graphql"
	"github.com/vektah/gqlparser/v2/ast"
)

// region    ************************** generated!.gotpl **************************

// endregion ************************** generated!.gotpl **************************

// region    ***************************** args.gotpl *****************************

// endregion ***************************** args.gotpl *****************************

// region    ************************** directives.gotpl **************************

// endregion ************************** directives.gotpl **************************

// region    **************************** field.gotpl *****************************

func (ec *executionContext) _MapNested_value(ctx context.Context, field graphql.CollectedField, obj *MapNested) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_MapNested_value(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp := ec._fieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Value, nil
	})

	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(CustomScalar)
	fc.Result = res
	return ec.marshalNCustomScalar2githubᚗcomᚋ99designsᚋgqlgenᚋcodegenᚋtestserverᚋfollowschemaᚐCustomScalar(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_MapNested_value(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "MapNested",
		Field:      field,
		IsMethod:   false,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			return nil, errors.New("field of type CustomScalar does not have child fields")
		},
	}
	return fc, nil
}

func (ec *executionContext) _MapStringInterfaceType_a(ctx context.Context, field graphql.CollectedField, obj map[string]interface{}) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_MapStringInterfaceType_a(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp := ec._fieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		switch v := obj["a"].(type) {
		case *string:
			return v, nil
		case string:
			return &v, nil
		case nil:
			return (*string)(nil), nil
		default:
			return nil, fmt.Errorf("unexpected type %T for field %s", v, "a")
		}
	})

	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(*string)
	fc.Result = res
	return ec.marshalOString2ᚖstring(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_MapStringInterfaceType_a(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "MapStringInterfaceType",
		Field:      field,
		IsMethod:   false,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			return nil, errors.New("field of type String does not have child fields")
		},
	}
	return fc, nil
}

func (ec *executionContext) _MapStringInterfaceType_b(ctx context.Context, field graphql.CollectedField, obj map[string]interface{}) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_MapStringInterfaceType_b(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp := ec._fieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		switch v := obj["b"].(type) {
		case *int:
			return v, nil
		case int:
			return &v, nil
		case nil:
			return (*int)(nil), nil
		default:
			return nil, fmt.Errorf("unexpected type %T for field %s", v, "b")
		}
	})

	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(*int)
	fc.Result = res
	return ec.marshalOInt2ᚖint(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_MapStringInterfaceType_b(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "MapStringInterfaceType",
		Field:      field,
		IsMethod:   false,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			return nil, errors.New("field of type Int does not have child fields")
		},
	}
	return fc, nil
}

func (ec *executionContext) _MapStringInterfaceType_c(ctx context.Context, field graphql.CollectedField, obj map[string]interface{}) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_MapStringInterfaceType_c(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp := ec._fieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		switch v := obj["c"].(type) {
		case *CustomScalar:
			return v, nil
		case CustomScalar:
			return &v, nil
		case nil:
			return (*CustomScalar)(nil), nil
		default:
			return nil, fmt.Errorf("unexpected type %T for field %s", v, "c")
		}
	})

	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(*CustomScalar)
	fc.Result = res
	return ec.marshalOCustomScalar2ᚖgithubᚗcomᚋ99designsᚋgqlgenᚋcodegenᚋtestserverᚋfollowschemaᚐCustomScalar(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_MapStringInterfaceType_c(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "MapStringInterfaceType",
		Field:      field,
		IsMethod:   false,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			return nil, errors.New("field of type CustomScalar does not have child fields")
		},
	}
	return fc, nil
}

func (ec *executionContext) _MapStringInterfaceType_nested(ctx context.Context, field graphql.CollectedField, obj map[string]interface{}) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_MapStringInterfaceType_nested(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp := ec._fieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		switch v := obj["nested"].(type) {
		case *MapNested:
			return v, nil
		case MapNested:
			return &v, nil
		case nil:
			return (*MapNested)(nil), nil
		default:
			return nil, fmt.Errorf("unexpected type %T for field %s", v, "nested")
		}
	})

	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(*MapNested)
	fc.Result = res
	return ec.marshalOMapNested2ᚖgithubᚗcomᚋ99designsᚋgqlgenᚋcodegenᚋtestserverᚋfollowschemaᚐMapNested(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_MapStringInterfaceType_nested(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "MapStringInterfaceType",
		Field:      field,
		IsMethod:   false,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			switch field.Name {
			case "value":
				return ec.fieldContext_MapNested_value(ctx, field)
			}
			return nil, fmt.Errorf("no field named %q was found under type MapNested", field.Name)
		},
	}
	return fc, nil
}

// endregion **************************** field.gotpl *****************************

// region    **************************** input.gotpl *****************************

func (ec *executionContext) unmarshalInputMapNestedInput(ctx context.Context, obj interface{}) (MapNested, error) {
	var it MapNested
	asMap := map[string]interface{}{}
	for k, v := range obj.(map[string]interface{}) {
		asMap[k] = v
	}

	fieldsInOrder := [...]string{"value"}
	for _, k := range fieldsInOrder {
		v, ok := asMap[k]
		if !ok {
			continue
		}
		switch k {
		case "value":
			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("value"))
			data, err := ec.unmarshalNCustomScalar2githubᚗcomᚋ99designsᚋgqlgenᚋcodegenᚋtestserverᚋfollowschemaᚐCustomScalar(ctx, v)
			if err != nil {
				return it, err
			}
			it.Value = data
		}
	}

	return it, nil
}

func (ec *executionContext) unmarshalInputMapStringInterfaceInput(ctx context.Context, obj interface{}) (map[string]interface{}, error) {
	it := make(map[string]interface{}, len(obj.(map[string]interface{})))
	asMap := map[string]interface{}{}
	for k, v := range obj.(map[string]interface{}) {
		asMap[k] = v
	}

	fieldsInOrder := [...]string{"a", "b", "c", "nested"}
	for _, k := range fieldsInOrder {
		v, ok := asMap[k]
		if !ok {
			continue
		}
		switch k {
		case "a":
			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("a"))
			data, err := ec.unmarshalNString2string(ctx, v)
			if err != nil {
				return it, err
			}
			it["a"] = data
		case "b":
			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("b"))
			data, err := ec.unmarshalOInt2ᚖint(ctx, v)
			if err != nil {
				return it, err
			}
			it["b"] = data
		case "c":
			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("c"))
			data, err := ec.unmarshalOCustomScalar2ᚖgithubᚗcomᚋ99designsᚋgqlgenᚋcodegenᚋtestserverᚋfollowschemaᚐCustomScalar(ctx, v)
			if err != nil {
				return it, err
			}
			it["c"] = data
		case "nested":
			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("nested"))
			data, err := ec.unmarshalOMapNestedInput2ᚖgithubᚗcomᚋ99designsᚋgqlgenᚋcodegenᚋtestserverᚋfollowschemaᚐMapNested(ctx, v)
			if err != nil {
				return it, err
			}
			it["nested"] = data
		}
	}

	return it, nil
}

func (ec *executionContext) unmarshalInputNestedMapInput(ctx context.Context, obj interface{}) (NestedMapInput, error) {
	var it NestedMapInput
	asMap := map[string]interface{}{}
	for k, v := range obj.(map[string]interface{}) {
		asMap[k] = v
	}

	fieldsInOrder := [...]string{"map"}
	for _, k := range fieldsInOrder {
		v, ok := asMap[k]
		if !ok {
			continue
		}
		switch k {
		case "map":
			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("map"))
			data, err := ec.unmarshalOMapStringInterfaceInput2map(ctx, v)
			if err != nil {
				return it, err
			}
			it.Map = data
		}
	}

	return it, nil
}

// endregion **************************** input.gotpl *****************************

// region    ************************** interface.gotpl ***************************

// endregion ************************** interface.gotpl ***************************

// region    **************************** object.gotpl ****************************

var mapNestedImplementors = []string{"MapNested"}

func (ec *executionContext) _MapNested(ctx context.Context, sel ast.SelectionSet, obj *MapNested) graphql.Marshaler {
	fields := graphql.CollectFields(ec.OperationContext, sel, mapNestedImplementors)

	out := graphql.NewFieldSet(fields)
	deferred := make(map[string]*graphql.FieldSet)
	for i, field := range fields {
		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("MapNested")
		case "value":
			out.Values[i] = ec._MapNested_value(ctx, field, obj)
			if out.Values[i] == graphql.Null {
				out.Invalids++
			}
		default:
			panic("unknown field " + strconv.Quote(field.Name))
		}
	}
	out.Dispatch(ctx)
	if out.Invalids > 0 {
		return graphql.Null
	}

	atomic.AddInt32(&ec.deferred, int32(len(deferred)))

	for label, dfs := range deferred {
		ec.processDeferredGroup(graphql.DeferredGroup{
			Label:    label,
			Path:     graphql.GetPath(ctx),
			FieldSet: dfs,
			Context:  ctx,
		})
	}

	return out
}

var mapStringInterfaceTypeImplementors = []string{"MapStringInterfaceType"}

func (ec *executionContext) _MapStringInterfaceType(ctx context.Context, sel ast.SelectionSet, obj map[string]interface{}) graphql.Marshaler {
	fields := graphql.CollectFields(ec.OperationContext, sel, mapStringInterfaceTypeImplementors)

	out := graphql.NewFieldSet(fields)
	deferred := make(map[string]*graphql.FieldSet)
	for i, field := range fields {
		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("MapStringInterfaceType")
		case "a":
			out.Values[i] = ec._MapStringInterfaceType_a(ctx, field, obj)
		case "b":
			out.Values[i] = ec._MapStringInterfaceType_b(ctx, field, obj)
		case "c":
			out.Values[i] = ec._MapStringInterfaceType_c(ctx, field, obj)
		case "nested":
			out.Values[i] = ec._MapStringInterfaceType_nested(ctx, field, obj)
		default:
			panic("unknown field " + strconv.Quote(field.Name))
		}
	}
	out.Dispatch(ctx)
	if out.Invalids > 0 {
		return graphql.Null
	}

	atomic.AddInt32(&ec.deferred, int32(len(deferred)))

	for label, dfs := range deferred {
		ec.processDeferredGroup(graphql.DeferredGroup{
			Label:    label,
			Path:     graphql.GetPath(ctx),
			FieldSet: dfs,
			Context:  ctx,
		})
	}

	return out
}

// endregion **************************** object.gotpl ****************************

// region    ***************************** type.gotpl *****************************

func (ec *executionContext) unmarshalNCustomScalar2githubᚗcomᚋ99designsᚋgqlgenᚋcodegenᚋtestserverᚋfollowschemaᚐCustomScalar(ctx context.Context, v interface{}) (CustomScalar, error) {
	var res CustomScalar
	err := res.UnmarshalGQL(v)
	return res, graphql.ErrorOnPath(ctx, err)
}

func (ec *executionContext) marshalNCustomScalar2githubᚗcomᚋ99designsᚋgqlgenᚋcodegenᚋtestserverᚋfollowschemaᚐCustomScalar(ctx context.Context, sel ast.SelectionSet, v CustomScalar) graphql.Marshaler {
	return v
}

func (ec *executionContext) unmarshalOCustomScalar2ᚖgithubᚗcomᚋ99designsᚋgqlgenᚋcodegenᚋtestserverᚋfollowschemaᚐCustomScalar(ctx context.Context, v interface{}) (*CustomScalar, error) {
	if v == nil {
		return nil, nil
	}
	var res = new(CustomScalar)
	err := res.UnmarshalGQL(v)
	return res, graphql.ErrorOnPath(ctx, err)
}

func (ec *executionContext) marshalOCustomScalar2ᚖgithubᚗcomᚋ99designsᚋgqlgenᚋcodegenᚋtestserverᚋfollowschemaᚐCustomScalar(ctx context.Context, sel ast.SelectionSet, v *CustomScalar) graphql.Marshaler {
	if v == nil {
		return graphql.Null
	}
	return v
}

func (ec *executionContext) marshalOMapNested2ᚖgithubᚗcomᚋ99designsᚋgqlgenᚋcodegenᚋtestserverᚋfollowschemaᚐMapNested(ctx context.Context, sel ast.SelectionSet, v *MapNested) graphql.Marshaler {
	if v == nil {
		return graphql.Null
	}
	return ec._MapNested(ctx, sel, v)
}

func (ec *executionContext) unmarshalOMapNestedInput2ᚖgithubᚗcomᚋ99designsᚋgqlgenᚋcodegenᚋtestserverᚋfollowschemaᚐMapNested(ctx context.Context, v interface{}) (*MapNested, error) {
	if v == nil {
		return nil, nil
	}
	res, err := ec.unmarshalInputMapNestedInput(ctx, v)
	return &res, graphql.ErrorOnPath(ctx, err)
}

func (ec *executionContext) unmarshalOMapStringInterfaceInput2map(ctx context.Context, v interface{}) (map[string]interface{}, error) {
	if v == nil {
		return nil, nil
	}
	res, err := ec.unmarshalInputMapStringInterfaceInput(ctx, v)
	return res, graphql.ErrorOnPath(ctx, err)
}

func (ec *executionContext) marshalOMapStringInterfaceType2map(ctx context.Context, sel ast.SelectionSet, v map[string]interface{}) graphql.Marshaler {
	if v == nil {
		return graphql.Null
	}
	return ec._MapStringInterfaceType(ctx, sel, v)
}

func (ec *executionContext) unmarshalONestedMapInput2ᚖgithubᚗcomᚋ99designsᚋgqlgenᚋcodegenᚋtestserverᚋfollowschemaᚐNestedMapInput(ctx context.Context, v interface{}) (*NestedMapInput, error) {
	if v == nil {
		return nil, nil
	}
	res, err := ec.unmarshalInputNestedMapInput(ctx, v)
	return &res, graphql.ErrorOnPath(ctx, err)
}

// endregion ***************************** type.gotpl *****************************
