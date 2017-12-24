# hue

A driver for controlling Phillips Hue lights.

## Requirements

1. Follow the guide at https://www.developers.meethue.com/documentation/getting-started up until you can successfully send commands via the web interface
2. Add your bridge's IP address and your user ID to your environment. The easiest way for testing is to create a *.env file like the following:
    ```
    export HUE_BRIDGE_IP=192.168.1.XX
    export HUE_USER_ID=YOUR_ID_HERE
    ```
    and then run `source your_file.env`.
3. Run `go install ./cmd/client` from the project root
4. Run client to run the demo program