# :globe_with_meridians:Network
Wrapper for network management on MacOS using go language.  

### Install
```
go get github.com/informeai/network
```
### Usage
Scanning with Airport.
```
airport := NewAirport()
wifis, err := airport.Scan()
if err != nil{
	log.Fatal(err)
}
for _, wifi := range wifis{
	fmt.Printf("SSID: %v - BSSID: %v\n",wifi.SSID,wifi.BSSID)
}
#OUTPUT:
# SSID: Rede 1 - BSSID: 50:d4:f7:65:6f:24
# SSID: WifiPlus - BSSID: bc:c8:0f:3b:4b:30
# ...
```
### List of Utility.
- [x] Airport
- [ ] Ping
- [ ] Traceroute
- [ ] Networksetup 
