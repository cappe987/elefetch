#!/bin/bash

RED="\033[31m"
GREEN="\033[32m"
YELLOW="\033[33m"
CYAN="\033[36m"
ENDCOLOR="\033[0m"

echo -e "${CYAN}"
echo "elefetch is cross-platform, neofetch alternative for fetching system info."
echo "-----------------------"
echo "Fetching latest version"
version=$(curl -s https://api.github.com/repos/burntcarrot/elefetch/releases/latest | grep '"tag_name"' | cut -d'"' -f 4)
echo "Installing elefetch $version"
echo "-----------------------"
echo -n -e "${ENDCOLOR}"


if [[ "$OSTYPE" == "linux-gnu"* ]]; then
        ARCH=$(uname -m)
        OS="Linux"
        
elif [[ "$OSTYPE" == "darwin"* ]]; then
        ARCH=$(uname -m)
        OS="Darwin"
else
        echo -e "${RED}"
        echo "Not Supported OS"
        echo -n -e "${ENDCOLOR}"
        exit 1
fi

echo -e "${YELLOW}"
echo -e "OS: $OS\tARCH: $ARCH"
echo "Fetching the latest release for $OS"
echo -n -e "${ENDCOLOR}"

URL=$(curl -s https://api.github.com/repos/burntcarrot/elefetch/releases/latest \
    | grep "browser_download_url" | grep -i "$OS.*$ARCH" \
    | cut -d '"' -f 4)

if [ -z "$URL" ]; then
    echo -e "${RED}"
    echo "Release not found for your OS and ARCH"
    echo -n -e "${ENDCOLOR}"
    exit 1
fi

curl -sL $URL -o /tmp/elefetch.tar.gz
tar -xzf /tmp/elefetch.tar.gz -C /tmp elefetch

echo -e "${YELLOW}"
echo "Installing binary...."
echo -n -e "${ENDCOLOR}"

mv /tmp/elefetch /usr/local/bin
if [ $? -ne 0 ]; then
        echo -e  "${YELLOW}"
        echo "Runing command again with sudo"
        echo -n -e "${ENDCOLOR}"
        sudo mv /tmp/elefetch /usr/local/bin
fi
chmod +x /usr/local/bin/elefetch

rm /tmp/elefetch*

echo -e "${GREEN}"
echo "elefetch has been successfully installed!"
echo "Check your installation by running elefetch --help."
echo "-----------------------------------------------"
echo -n -e "${ENDCOLOR}"
