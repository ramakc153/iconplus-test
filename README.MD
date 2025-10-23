# REPOSITORY FOR ICONPLUS TEST



<h1>How to clone</h1>


```bash
git clone https://github.com/ramakc153/iconplus-test/

#on windows
go build . && ./iconplus-test.exe

#on linux
go build . && ./iconplus-test

```

<summary><code>GET</code> <code><b>/</b></code> <code>to check if the server is running</code></summary>
<summary><code>GET</code> <code><b>/api</b></code></summary>

### Response fields

| Name  | Type | 
|-------|-----|
| officeName | string |
| roomName| string  | 
| usagePercentage | float    |
| ConsumptionPrice | int  |
| Consumption List | array of objects |

<h1>How to run the test file </h1>

```bash
#make sure you're already on iconplus-test directory
cd logicaltest
go test

#or to make it verbose
go test -v
```
<p> or you can just call the function by yourself and passing appropriate parameter</p>

<h1>Notes</h1>


<p>Due to lack of information from the base API, usagePercentage field is calculate by hours used in X room / total hours used across all the room</p>

