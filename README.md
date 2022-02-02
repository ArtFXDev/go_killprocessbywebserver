# Silex GO Kill Process
[![made-with-Go](https://img.shields.io/badge/Made%20with-Go-1f425f.svg)](http://golang.org)
[![GitHub go.mod Go version of a Go module](https://img.shields.io/github/go-mod/go-version/artfxdev/silex_GoKillProcess)](https://github.com/ArtFXDev/silex_GoKillProcess)

Webserver running locally, listenning on port 5119 to kill process by pid.

## Why ?
In our pipeline, when user whant to kill job runnings on her computer, sometimes tractor cant kill remotely this jobs.
So Jobs are running as user system on computer and user don't have enough access to kill mannually the jobs.

This app is deployed as windows services in user system, with this setup, we just have to send a post webrequest with the pid as param and we can kill the jobs.


## Usage

```
http://localhost:1159/kill/{YOUR_PID}
```
Using postman 
<p align="center">
  <img src="https://github.com/ArtFXDev/silex_GoKillProcess/blob/main/screenshots/postrequest.png?raw=true">
</p>

## Usage in our pipeline
Using [NSSM](https://nssm.cc) to convert .exe to service

You can find the deployement script [here](https://github.com/ArtFXDev/silex_fog_snapin/blob/main/gokillprocess/go-killprocess.ps1)
