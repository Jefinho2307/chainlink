package evm

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"

	"github.com/smartcontractkit/chainlink-common/pkg/codec"

	commontypes "github.com/smartcontractkit/chainlink-common/pkg/types"

	"github.com/smartcontractkit/chainlink/v2/core/services/relay/evm/types"
)

type codecEntry struct {
	Args           abi.Arguments
	encodingPrefix []byte
	checkedType    reflect.Type
	nativeType     reflect.Type
	mod            codec.Modifier
}

func (entry *codecEntry) Init() error {
	if entry.checkedType != nil {
		return nil
	}

	args := UnwrapArgs(entry.Args)
	argLen := len(args)
	native := make([]reflect.StructField, argLen)
	checked := make([]reflect.StructField, argLen)

	// Single returns that aren't named will return that type
	// whereas named parameters will return a struct with the fields
	// Eg: function foo() returns (int256) ... will return a *big.Int for the native type
	// function foo() returns (int256 i) ... will return a struct { I *big.Int } for the native type
	// function foo() returns (int256 i1, int256 i2) ... will return a struct { I1 *big.Int, I2 *big.Int } for the native type
	if len(args) == 1 && args[0].Name == "" {
		nativeArg, checkedArg, err := getNativeAndCheckedTypesForArg(&args[0])
		if err != nil {
			return err
		}
		entry.nativeType = nativeArg
		entry.checkedType = checkedArg
		return nil
	}

	numIndices := 0
	seenNames := map[string]bool{}
	for i, arg := range args {
		if arg.Indexed {
			if numIndices == maxTopicFields {
				return fmt.Errorf("%w: too many indexed arguments", commontypes.ErrInvalidConfig)
			}
			numIndices++
		}

		tmp := arg
		nativeArg, checkedArg, err := getNativeAndCheckedTypesForArg(&tmp)
		if err != nil {
			return err
		}
		if len(arg.Name) == 0 {
			return fmt.Errorf("%w: empty field names are not supported for multiple returns", commontypes.ErrInvalidType)
		}

		name := strings.ToUpper(arg.Name[:1]) + arg.Name[1:]
		if seenNames[name] {
			return fmt.Errorf("%w: duplicate field name %s, after ToCamelCase", commontypes.ErrInvalidConfig, name)
		}
		seenNames[name] = true
		native[i] = reflect.StructField{Name: name, Type: nativeArg}
		checked[i] = reflect.StructField{Name: name, Type: checkedArg}
	}

	entry.nativeType = reflect.StructOf(native)
	entry.checkedType = reflect.StructOf(checked)
	return nil
}

func (entry *codecEntry) GetMaxSize(n int) (int, error) {
	if entry == nil {
		return 0, fmt.Errorf("%w: nil entry", commontypes.ErrInvalidType)
	}
	return GetMaxSize(n, entry.Args)
}

func UnwrapArgs(args abi.Arguments) abi.Arguments {
	// Unwrap an unnamed tuple so that callers don't need to wrap it
	// Eg: If you have struct Foo { ... } and return an unnamed Foo, you should be able ot decode to a go Foo{} directly
	if len(args) != 1 || args[0].Name != "" {
		return args
	}

	elms := args[0].Type.TupleElems
	if len(elms) != 0 {
		names := args[0].Type.TupleRawNames
		args = make(abi.Arguments, len(elms))
		for i, elm := range elms {
			args[i] = abi.Argument{
				Name: names[i],
				Type: *elm,
			}
		}
	}
	return args
}

func getNativeAndCheckedTypesForArg(arg *abi.Argument) (reflect.Type, reflect.Type, error) {
	tmp := arg.Type
	if arg.Indexed {
		switch arg.Type.T {
		case abi.StringTy:
			return reflect.TypeOf(common.Hash{}), reflect.TypeOf(common.Hash{}), nil
		case abi.ArrayTy:
			u8, _ := types.GetAbiEncodingType("uint8")
			if arg.Type.Elem.GetType() == u8.Native {
				return reflect.TypeOf(common.Hash{}), reflect.TypeOf(common.Hash{}), nil
			}
			fallthrough
		case abi.SliceTy, abi.TupleTy, abi.FixedBytesTy, abi.FixedPointTy, abi.FunctionTy:
			// https://github.com/ethereum/go-ethereum/blob/release/1.12/accounts/abi/topics.go#L78
			return nil, nil, fmt.Errorf("%w: unsupported indexed type: %v", commontypes.ErrInvalidConfig, arg.Type)
		default:
		}
	}

	return getNativeAndCheckedTypes(&tmp)
}

func getNativeAndCheckedTypes(curType *abi.Type) (reflect.Type, reflect.Type, error) {
	converter := func(t reflect.Type) reflect.Type { return t }
	for curType.Elem != nil {
		prior := converter
		switch curType.GetType().Kind() {
		case reflect.Slice:
			converter = func(t reflect.Type) reflect.Type {
				return prior(reflect.SliceOf(t))
			}
			curType = curType.Elem
		case reflect.Array:
			tmp := curType
			converter = func(t reflect.Type) reflect.Type {
				return prior(reflect.ArrayOf(tmp.Size, t))
			}
			curType = curType.Elem
		default:
			return nil, nil, fmt.Errorf(
				"%w: cannot create type for kind %v", commontypes.ErrInvalidType, curType.GetType().Kind())
		}
	}
	base, ok := types.GetAbiEncodingType(curType.String())
	if ok {
		return converter(base.Native), converter(base.Checked), nil
	}

	return createTupleType(curType, converter)
}

func createTupleType(curType *abi.Type, converter func(reflect.Type) reflect.Type) (reflect.Type, reflect.Type, error) {
	if len(curType.TupleElems) == 0 {
		if curType.TupleType == nil {
			return nil, nil, fmt.Errorf("%w: unsupported solitidy type: %v", commontypes.ErrInvalidType, curType.String())
		}
		return curType.TupleType, curType.TupleType, nil
	}

	// Create native type ourselves to assure that it'll always have the exact memory layout of checked types
	// Otherwise, the "unsafe" casting that will be done to convert from checked to native won't be safe.
	// At the time of writing, the way the TupleType is built it will be the same, but I don't want to rely on that
	// If they ever add private fields for internal tracking
	// or anything it would break us if we don't build the native type.
	// As an example of how it could possibly change in the future, I've seen struct{}
	// added with tags to the top of generated structs to allow metadata exploration.
	nativeFields := make([]reflect.StructField, len(curType.TupleElems))
	checkedFields := make([]reflect.StructField, len(curType.TupleElems))
	for i, elm := range curType.TupleElems {
		name := curType.TupleRawNames[i]
		nativeFields[i].Name = name
		checkedFields[i].Name = name
		nativeArgType, checkedArgType, err := getNativeAndCheckedTypes(elm)
		if err != nil {
			return nil, nil, err
		}
		nativeFields[i].Type = nativeArgType
		checkedFields[i].Type = checkedArgType
	}
	return converter(reflect.StructOf(nativeFields)), converter(reflect.StructOf(checkedFields)), nil
}
