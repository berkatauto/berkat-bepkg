# berkat-bepkg - Dokumentasi

Package Backend khusus untuk aplikasi website Berkat Auto.

## Update library untuk backend ini

```
go get -u all
go mod tidy
git tag                                                 # Cek version untuk package ini
git tag v1.0.0                                          # Set version untuk package ini
git push origin --tags                                  # Push Package dengan Version
go list -m github.com/berkatauto/berkat-bepkg@v1.0.0    # Publish Package ke pkg.go.dev
```

## Cara Install Package ini

Untuk menginstal package ini cukup mudah. Buat Project Golangmu dan masukkan command berikut.

```
go mod init "init kemana nih"
go get github.com/berkatauto/berkat-bepkg
go mod tidy
```

Package siap langsung pakai.
