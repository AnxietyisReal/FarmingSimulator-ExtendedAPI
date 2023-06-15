# FarmingSimulator-ExtendedAPI
This extends the current GIANTS' Farming Simulator Server API by fetching XML files over FTP. 

The URL path can be found at `your.api.server.ip:8095`, you can read any file that is visible in the current directory, e.g `your.api.server.ip:8095/environment.xml` which returns the contents of that file.

## JSON Schema
The filename should be `ftp_details.json`\
The following schema will be like this:
```json
{
  "host": "IP of the FTP server that has a savegame folder",
  "path": "Current path, e.g /savegame1",
  "username": "FTP username",
  "password": "FTP password"
}
```