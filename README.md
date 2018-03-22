# My Gin Tonic Benchmark

Simple web aplication that uses Gin is a web framework written in Go (Golang).

## Services Benchmark

### API requirements

The API must have a series of four endpoints to be tested. All of them must be fully implemented and without bugs. Please respect the path of each endpoint otherwise the tests will fail.

#### This are the required endpoints

##### Simple throughput endpoint: `/throughput`
This endpoint should respond with the following JSON response:
```javascript
{"throughput": "Lorem ipsum dolor sit amet, consectetur adipiscing elit. In in ipsum a velit faucibus tempor vel nec odio. Interdum et malesuada fames ac ante ipsum primis in faucibus. Curabitur quis orci eget purus tempus aliquet eu eu risus. Ut velit elit, viverra et ex vel, scelerisque rhoncus odio. Donec vitae diam pellentesque, commodo velit et, lacinia leo. In vel pharetra purus, sed eleifend nunc. Maecenas porta rhoncus consectetur. In hac habitasse platea dictumst. Duis sed erat nibh. Morbi imperdiet lorem purus, vitae facilisis enim maximus et. Phasellus ullamcorper sapien eget neque eleifend malesuada."}
```
##### CPU intensive endpoint: `/cpu`
This endpoint must hash 256 times with SHA512 using the means  provided by the standard library the el string “Sparkers doing some benchmarking” and produce the following JSON response 
```javascript
{"Hashed": "..."}
```

##### RAM intensive endpoint : `/ram`
Read `support/ram_test.txt` and respond with the number of vowels present on the text. A regular expression must be used. Expected JSON  response
```javascript 
{ "n_vowels": 1234 }
```

##### Disk intensive endpoint : `/disk`
Read the `support/disk_test.csv` file, dump everything to a new file, store the file on the temp directory of the OS and then delete it. Respond with the number of read bytes from the original file. Expected output JSON 
```javascript
{"bytes": 1234}
```

#### Container requirements
* Should serve the api on the exported port 80
* Should be based on an Alpine image
