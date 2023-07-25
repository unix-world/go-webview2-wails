//go:build windows

package webview2

import (
	"golang.org/x/sys/windows"
	"syscall"
	"unsafe"
)

type ICoreWebView2_13Vtbl struct {
	IUnknownVtbl
	GetProfile ComProc
}

type ICoreWebView2_13 struct {
	Vtbl *ICoreWebView2_13Vtbl
}

func (i *ICoreWebView2_13) AddRef() uintptr {
	refCounter, _, _ := i.Vtbl.AddRef.Call(uintptr(unsafe.Pointer(i)))
	return refCounter
}

func (i *ICoreWebView2) GetICoreWebView2_13() *ICoreWebView2_13 {
	var result *ICoreWebView2_13

	iidICoreWebView2_13 := NewGUID("{F75F09A8-667E-4983-88D6-C8773F315E84}")
	_, _, _ = i.Vtbl.QueryInterface.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(iidICoreWebView2_13)),
		uintptr(unsafe.Pointer(&result)))

	return result
}

func (i *ICoreWebView2_13) GetProfile() (*ICoreWebView2Profile, error) {

	var value ICoreWebView2Profile

	hr, _, err := i.Vtbl.GetProfile.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&value)),
	)
	if windows.Handle(hr) != windows.S_OK {
		return nil, syscall.Errno(hr)
	}
	return &value, err
}
