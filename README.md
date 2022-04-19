# CS5970_Lab2

Hello, we will talk about generating vanity address in this lab. 

Vanity address are created by cycling through addresses and finding the matching private key for it.

We will borrow code from https://github.com/MarinX/btc-vanity. MarinX has done most of the work for us, so we will just use his code and change a few things to get it working like how we want!

When we first start the lab in VS Code, we want to make sure eveything is updated. To do that we will press Ctrl+Shift+P and type the command ">Go: Install/Update Tools". Check all the boxes and let it update the tools necessary for this project.

Once we have our tools updated, we will open an empty directory. We will first create our go.mod file. In order to create a new go.mod file, open up a new session in the terimal. Once in terminal, type in "go mod init 'filename'". After that, type in the command "go get github.com/MarinX/btc-vanity". That will create a go.mod and go.sum files for us. I have uploaded those files to the github for reference.

We will create a new file with .go extension. We will name it same as the filename when we initiated the go mod file. I named my file vanityAddress.go, you can name it whatever you want. Once the file is created. We will use the code provided by Marinx. Once we type the code in, we will want to go into Find function located in between lines 20-25. Once the Find is located we will want to press Ctrl and click on the Find function. It will take us to a file named btcvanity.go. Within that file we want to scroll down to line 47 and find the isMatch function. We will once again press Ctrl and click on the isMatch function. That will take us to a file called matcher.go. In that file, we will remove the specified length used in the ToLower function, so it should look something like: addr = strings.ToLower(addr). Which will allow us to get the pattern we want at the end of the address. Then we will change the functoin HasPrefix to HasSuffix, so it should look something like: return strings.HasSuffix(addr, pattern). After that save all the files. Overide and save if required. 

Now, we will come back to the file we created. Within that file, we will set the TestNet variable to true. The variable will be found in between lines 14-19. Then we will change the string to the to what we desire. I wanted to find "Bhav" since that is the first four letters of my name. Once you type in your pattern, please save the file.

We will come to terminal and type in the command "go run .", which will run the base directory. It will generate bunch of addresses, but will only display the one that matches for your pattern. It will give your public key and the private key associated with the public key. For me I got my pair to be the following:

PUBLIC KEY: mjNm3SearukpQRqQjSCQk4iNjD7AVpBhaV

PRIVATE KEY: 92mXu1zZcK7KUkfYeQtLfu6gwerZFN2GHoHsHFKhPSqS6cxvdjk

Once that is done. We will want to upload the private key in the Bitcoin Core software. We will open up the console by going to Windows and then Clicking on console. Once in the console, type in the command "importprivkey <key> "" false", where <key> is your private key. Once you have that uploaded, send some testnet bitcoins to your address. Once that is done, we will use another function called "getblockchaininfo". When you type that in and press enter it will tell you where you pruned your height. Therefore, we will take that height, for me it was block number 1902678. Then I tpye in the console "rescanblockchain (1902678 2196234)" so it will rescan the blocks from the time I had it pruned. Once the rescan is complete, the bitcoins will show up. From there, you can send the bitcoins to Batman! I have uploaded a few images in the images folder showing bitcoins going to batman from my public key. I have also uploaded a couple images of VS Code and BitCoin Core Software.
  
Thank you!
  Bhaven Patel
