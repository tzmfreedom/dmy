#!/bin/bash

NAME="dmy"
VERSION="0.1.0"
PREFIX="/usr/local"
GITHUB_USER="tzmfreedom"
TMP_DIR="/tmp"

set -ue

UNAME=$(uname -s)
if [ "$UNAME" != "Linux" -a "$UNAME" != "Darwin" ] ; then
    echo "Sorry, windows user download from release page to install dmy."
    exit 1
fi


UNAME_P=$(uname -p)
if [ "$UNAME" = "Darwin" ] ; then
  OS="darwin"
elif [ "$UNAME" = "Linux" ] ; then
  OS="linux"
fi

if [ "i386" == "${UNAME_P}" -o "i686" == "${UNAME_P}" ] ; then
  ARCH="386"
else
  ARCH="amd64"
fi

ARCHIVE_FILE=${NAME}-${VERSION}-${OS}-${ARCH}.tar.gz
BINARY="https://github.com/${GITHUB_USER}/${NAME}/releases/download/v${VERSION}/${ARCHIVE_FILE}"

cd $TMP_DIR
curl -L -O ${BINARY}

tar xzf ${ARCHIVE_FILE}
mv ${OS}-${ARCH}/${NAME} ${PREFIX}/bin/${NAME} 
rm -rf ${OS}-${ARCH}
rm -rf ${ARCHIVE_FILE}
