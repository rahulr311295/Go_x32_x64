Installation steps for golang 32bit and 64bit

32bit
-----


wget https://storage.googleapis.com/golang/go1.8.3.linux-386.tar.gz && \
sudo tar -C /usr/local -xzf go1.8.3.linux-386.tar.gz && \
rm go1.8.3.linux-386.tar.gz && \
echo 'export PATH=$PATH:/usr/local/go/bin' | sudo tee -a /etc/profile && \
echo 'export GOPATH=$HOME/go' | tee -a $HOME/.bashrc && \
echo 'export PATH=$PATH:$GOROOT/bin:$GOPATH/bin' | tee -a $HOME/.bashrc && \
mkdir -p $HOME/go/{src,pkg,bin}



64bit
-----


wget https://storage.googleapis.com/golang/go1.8.3.linux-amd64.tar.gz && \
sudo tar -C /usr/local -xzf go1.8.3.linux-amd64.tar.gz && \
rm go1.8.3.linux-amd64.tar.gz && \
echo 'export PATH=$PATH:/usr/local/go/bin' | sudo tee -a /etc/profile && \
echo 'export GOPATH=$HOME/go' | tee -a $HOME/.bashrc && \
echo 'export PATH=$PATH:$GOROOT/bin:$GOPATH/bin' | tee -a $HOME/.bashrc && \
mkdir -p $HOME/go/{src,pkg,bin}


You may need to restart any open terminals for the change to take effect. 

Windows x64 and x32
-------------------

First download the below executable file

https://dl.google.com/go/go1.10.windows-amd64.msi

After installation your GROOT directory will be C:\Go 

Setting environmental variables 

Windows 7 ,8 
-------------

Go to My Computer --> System Properties --> Advanced Settings --> Environmental Variables


Windows 10 

Go to control panel and in the search box type Environmental Variable 

Rest of the steps are shown in the image

and copy these locations

C:\Go\
C:\Go\bin

You may need to restart any open command prompts for the change to take effect. 



ETHEREUM WALLET
---------------

Windows
-------

https://github.com/ethereum/mist/releases/download/v0.9.3/Ethereum-Wallet-installer-0-9-3.exe


Linux
-----
32bit
-----
https://github.com/ethereum/mist/releases/download/v0.9.3/Ethereum-Wallet-linux32-0-9-3.deb


64bit
-----
https://github.com/ethereum/mist/releases/download/v0.9.3/Ethereum-Wallet-linux64-0-9-3.deb
