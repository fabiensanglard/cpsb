mkdir -p out
cd out
touch book.idx
makeindex book.idx
cd ..
go run build.go $@