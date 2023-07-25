//go:build windows

package webview2

import (
	"golang.org/x/sys/windows"
	"syscall"
	"unsafe"
)

type ICoreWebView2FrameInfoCollectionVtbl struct {
	IUnknownVtbl
	GetIterator ComProc
}

type ICoreWebView2FrameInfoCollection struct {
	Vtbl *ICoreWebView2FrameInfoCollectionVtbl
}

func (i *ICoreWebView2FrameInfoCollection) AddRef() uintptr {
	refCounter, _, _ := i.Vtbl.AddRef.Call(uintptr(unsafe.Pointer(i)))
	return refCounter
}

func (i *ICoreWebView2FrameInfoCollection) GetIterator() (*ICoreWebView2FrameInfoCollectionIterator, error) {

	var iterator ICoreWebView2FrameInfoCollectionIterator

	hr, _, err := i.Vtbl.GetIterator.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&iterator)),
	)
	if windows.Handle(hr) != windows.S_OK {
		return nil, syscall.Errno(hr)
	}
	return &iterator, err
}
