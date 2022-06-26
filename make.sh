mkdir -p out
cd out
makeindex book.idx
cd ..
go run build.go $@