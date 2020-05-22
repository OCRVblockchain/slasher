

### HOWTO

Build executable binary file:

    go build -o slasher

Configure your revocation or removal request in [/pkg/config/config.yaml.default](https://github.com/OCRVblockchain/slasher/blob/master/pkg/config/config.yaml.default) and rename config.yaml.default to config.yaml           
_(or use two copies, because first sample config file not recognized by Slasher)_

Start:

    ./slasher --identity admin  --secret adminpw --mode fullslash
    
--mode flag:

    revokecert - for certificate revocation 
    removeidentity - for identity removal 
    fullslash - for both options
    delall - delete all users (certs and identities)