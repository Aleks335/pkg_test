module github.com/Aleks335/pkg_test

go 1.22.1

require github.com/go-chi/chi/v5 v5.0.12

// учтановка зависимости counter )  go get github.com/Aleks335/pkg_test/pkg/counter
require github.com/Aleks335/pkg_test/pkg/counter v0.0.0-00010101000000-000000000000 // indirect

// обязательно прописать эту строчку чтобы библиотека бралась по локальному пути
replace github.com/Aleks335/pkg_test/pkg/counter => ./pkg/counter
