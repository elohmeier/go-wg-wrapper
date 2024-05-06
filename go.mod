module github.com/aschmidt75/go-wg-wrapper

go 1.15

replace github.com/aschmidt75/go-wg-wgrapper/wgwrapper => ./pkg/wgwrapper

require (
	github.com/google/go-cmp v0.6.0 // indirect
	github.com/jsimonetti/rtnetlink v0.0.0-20200117123717-f846d4f6c1f4 // indirect
	github.com/mdlayher/socket v0.5.1 // indirect
	golang.org/x/net v0.24.0 // indirect
	golang.org/x/sync v0.7.0 // indirect
	golang.org/x/sys v0.20.0 // indirect
	golang.zx2c4.com/wireguard v0.0.20200121 // indirect
	golang.zx2c4.com/wireguard/wgctrl v0.0.0-20230429144221-925a1e7659e6
)
