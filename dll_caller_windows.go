package dll_caller

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"syscall"
	"unsafe"
)

var (
	kernel32, _ = syscall.LoadLibrary("kernel32.dll")
)

type Dll struct {
	FileName   string
	dllHandler syscall.Handle
	funcProcs  map[string]uintptr
}

type FuncCallResult struct {
	Ret1  uintptr
	Ret2  uintptr
	Errno syscall.Errno
}

func NewDll(fileName string) (dll *Dll, err error) {
	newDll := new(Dll)
	if newDll.funcProcs == nil {
		newDll.funcProcs = make(map[string]uintptr)
	}

	if err = newDll.LoadLibrary(fileName); err != nil {
		return
	}

	return newDll, nil
}

func (p *Dll) LoadLibrary(fileName string) error {
	if handler, e := syscall.LoadLibrary(fileName); e != nil {
		return e
	} else {
		p.dllHandler = handler
	}
	return nil
}

func (p *Dll) FreeLibrary() error {
	if p.IsDllLoaded() {
		p.dllHandler = 0
		return syscall.FreeLibrary(p.dllHandler)
	}
	return nil
}

func (p *Dll) IsDllLoaded() bool {
	if uintptr(p.dllHandler) == 0 {
		return false
	}
	return true
}

func (p *Dll) InitalFunctions(funcNames ...string) error {
	if funcNames == nil {
		return nil
	}

	if !p.IsDllLoaded() {
		return errors.New("dll should loaded befor inital functions")
	}

	if p.funcProcs == nil {
		p.funcProcs = make(map[string]uintptr)
	}

	for _, funcName := range funcNames {
		funcName = strings.TrimSpace(funcName)
		if funcName == "" {
			return errors.New("function name could not be empty")
		}
		if proc, e := syscall.GetProcAddress(p.dllHandler, funcName); e != nil {
			return e
		} else {
			p.funcProcs[funcName] = proc
		}
	}
	return nil
}

func (p *Dll) Call(funcName string, funcParams ...interface{}) (result FuncCallResult, err error) {
	var lenParam uintptr = uintptr(len(funcParams))

	if p.funcProcs == nil {
		err = errors.New("function address not initaled")
		return
	}

	var funcAddress uintptr
	if addr, exist := p.funcProcs[funcName]; !exist {
		err = errors.New("function address not exist")
		return
	} else {
		funcAddress = addr
	}

	var r1, r2 uintptr
	var errno syscall.Errno

	var a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15 uintptr

	for parmIndex, param := range funcParams {
		var vPtr uintptr = 0

		switch v := param.(type) {
		case string:
			vPtr = uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(v)))
		case *string:
			vPtr = uintptr(unsafe.Pointer(syscall.StringBytePtr(*v)))
		case bool:
			vPtr = uintptr(unsafe.Pointer(&v))
		case int:
			vPtr = uintptr(v)
		case int8:
			vPtr = uintptr(v)
		case uint8:
			vPtr = uintptr(v)
		case *uint8:
			vPtr = uintptr(unsafe.Pointer(v))
		case int16:
			vPtr = uintptr(v)
		case uint16:
			vPtr = uintptr(v)
		case *uint16:
			vPtr = uintptr(unsafe.Pointer(v))
		case int32:
			vPtr = uintptr(v)
		case uint32:
			vPtr = uintptr(v)
		case *uint32:
			vPtr = uintptr(unsafe.Pointer(v))
		case int64:
			vPtr = uintptr(v)
		case uint64:
			vPtr = uintptr(v)
		case *uint64:
			vPtr = uintptr(unsafe.Pointer(v))
		case float32:
			vPtr = uintptr(v)
		case float64:
			vPtr = uintptr(v)
		case []byte:
			vPtr = uintptr(unsafe.Pointer(&v[0]))
		case uintptr:
			ptr, _ := param.(uintptr)
			vPtr = ptr
		default:
			err = fmt.Errorf("unsupport convert type %v to uintptr", reflect.TypeOf(param))
			return
		}

		switch parmIndex + 1 {
		case 1:
			a1 = vPtr
		case 2:
			a2 = vPtr
		case 3:
			a3 = vPtr
		case 4:
			a4 = vPtr
		case 5:
			a5 = vPtr
		case 6:
			a6 = vPtr
		case 7:
			a7 = vPtr
		case 8:
			a8 = vPtr
		case 9:
			a9 = vPtr
		case 10:
			a10 = vPtr
		case 11:
			a11 = vPtr
		case 12:
			a12 = vPtr
		case 13:
			a13 = vPtr
		case 14:
			a14 = vPtr
		case 15:
			a15 = vPtr
		}
	}

	r1, r2, errno = syscall.Syscall15(funcAddress, lenParam, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15)
	result.Ret1 = r1
	result.Ret2 = r2
	result.Errno = errno

	return
}
