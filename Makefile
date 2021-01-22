SHELL=powershell.exe

include mk/build-linux.mk
include mk/build-windows.mk
include mk/clean.mk
include mk/gen-protobuf.mk