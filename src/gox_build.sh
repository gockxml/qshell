export GOPATH=$GOPATH:/Users/gock/github/qshell
gox -os="darwin" -arch="amd64" -output="../bin/qshell_darwin_amd64" 
gox -os="darwin" -arch="386" -output="../bin/qshell_darwin_386"
gox -os="windows" -arch="amd64" -output="../bin/qshell_windows_amd64"
gox -os="windows" -arch="386" -output="../bin/qshell_windows_386"
gox -os="linux" -arch="amd64" -output="../bin/qshell_linux_amd64"
gox -os="linux" -arch="386" -output="../bin/qshell_linux_386"
gox -os="linux" -arch="arm" -output="../bin/qshell_linux_arm"
