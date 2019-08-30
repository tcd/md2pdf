date="$(date +%F)"
version=$(git describe --always)
hash=$(git rev-parse --short HEAD)

major=$(echo "$version" | cut -d. -f1 | sed 's/v//g')
minor=$(echo "$version" | cut -d. -f2)
patch=$(echo "$version" | cut -d. -f3 | cut -d- -f1)

filepath="github.com/tcd/md2pdf/internal/md2pdf"

flags="
  -X $filepath.VersionMajor=$major
  -X $filepath.VersionMinor=$minor
  -X $filepath.VersionPatch=$patch
  -X $filepath.VersionHash=$hash
  -X $filepath.VersionDate=$date
"

GO111MODULE=on
go install -ldflags="$flags"
