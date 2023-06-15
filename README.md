# FarmingSimulator-ExtendedAPI
This extends the current GIANTS' Farming Simulator Server API by fetching XML files over FTP. 

## JSON Schema
The filename should be `ftp_details.json`\
The following schema will be like this:
```json
{
  "host": "IP of the FTP server that has a savegame folder",
  "path": "Current path, e.g /savegame1",
  "file": "File you want the API to fetch the data and use it in your projects, e.g careerSavegame.xml",
  "username": "FTP username",
  "password": "FTP password"
}
```