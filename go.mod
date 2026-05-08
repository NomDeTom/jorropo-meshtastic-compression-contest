module github.com/Jorropo/meshtastic-compression-contest

go 1.25

toolchain go1.25.4

require (
	github.com/cespare/go-smaz v1.0.0
	github.com/cloudflare/golz4 v0.0.0-20240916140612-caecf3c00c06
	github.com/egonelbre/exp-protobuf-compression v0.0.0-00010101000000-000000000000
	github.com/inkyblackness/res v0.0.0-20180728072643-e78e2ff1969d
	github.com/klauspost/compress v1.18.2
	github.com/pierrec/lz4/v4 v4.1.23
	github.com/tmthrgd/shoco v0.0.0-20190904060532-13bc643e0d0b
	golang.org/x/sys v0.36.0
	gonum.org/v1/plot v0.16.0
	google.golang.org/protobuf v1.36.11
	modernc.org/sqlite v1.41.0
)

require (
	codeberg.org/go-fonts/liberation v0.5.0 // indirect
	codeberg.org/go-latex/latex v0.1.0 // indirect
	codeberg.org/go-pdf/fpdf v0.10.0 // indirect
	git.sr.ht/~sbinet/gg v0.6.0 // indirect
	github.com/ajstarks/svgo v0.0.0-20211024235047-1546f124cd8b // indirect
	github.com/campoy/embedmd v1.0.0 // indirect
	github.com/dustin/go-humanize v1.0.1 // indirect
	github.com/golang/freetype v0.0.0-20170609003504-e2365dfdc4a0 // indirect
	github.com/google/uuid v1.6.0 // indirect
	github.com/mattn/go-isatty v0.0.20 // indirect
	github.com/ncruces/go-strftime v0.1.9 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/remyoudompheng/bigfft v0.0.0-20230129092748-24d4a6f8daec // indirect
	golang.org/x/exp v0.0.0-20250620022241-b7579e27df2b // indirect
	golang.org/x/image v0.25.0 // indirect
	golang.org/x/text v0.23.0 // indirect
	gopkg.in/check.v1 v1.0.0-20201130134442-10cb98267c6c // indirect
	modernc.org/libc v1.66.10 // indirect
	modernc.org/mathutil v1.7.1 // indirect
	modernc.org/memory v1.11.0 // indirect
)

replace github.com/egonelbre/exp-protobuf-compression => github.com/Jorropo/exp-protobuf-compression v0.0.0-20251230095524-70f04914e5b9 // https://github.com/egonelbre/exp-protobuf-compression/pull/1
