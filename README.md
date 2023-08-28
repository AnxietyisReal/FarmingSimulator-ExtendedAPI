<h1 align="center">FarmingSimulator-ExtendedAPI - JSON version</h1>
<p align="center">This extends the current GIANTS' Farming Simulator Server API by fetching XML files over FTP.</p>
<p align="center">The URL path can be found at <code>your.api.server.ip:8095</code>, you can read any file that is visible in the current directory, e.g <code>your.api.server.ip:8095/environment.xml</code> which returns the contents of that file.</p>

## JSON Schema
The filename should be `ftp_details.json` in the root directory  
The following schema will be like this:
```json
{
  "host": "1.2.3.4:21",
  "path": "/path/to/savegame_or_docs",
  "username": "FTP username",
  "password": "FTP password"
}
```
