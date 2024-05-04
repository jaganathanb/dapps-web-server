bin_name='DAppsWebServer'  # Change the binary name as desired

archs=("386" "amd64")
GOOS="windows"

# Look for signs of trouble in each log
for i in ${!archs[@]};
do
arch=${archs[$i]}
echo "Building for platform - ${arch}..."

CGO_ENABLED=1 GOOS="${GOOS}" GOARCH=${arch} \
         go build -o "./out1/${bin_name}-${arch}.exe"
done