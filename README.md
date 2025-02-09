# AG-Grid DataTable

1. [ Overview. ](#overview)
2. [ How to run. ](#run)
3. [ Configuration. ](#configuration)
4. [ Distribution - Version. ](#distribution)
5. [ TODO ](#todo)

<a name="overview"></a>
## 1. Overview

Creates a polling based datatable grid using JS package: https://www.ag-grid.com/

<a name="run"></a>
## 2. How to run

Set in file `dataendpoints.join`, the URLs from where the JSON data are retrieved.

Example:
```
{
    "1": "http://localhost:8090/mockdataalpha",
    "2": "http://localhost:8090/mockdatabeta",
    "3": "http://localhost:8090/mockdatagamma"
}
```
Set under `.env` file, the environmental variable that indicates the server's port.
```
HTTP_PORT=8090
```
Optional: In `assets/js/index.js` you can change the polling interval update in milliseconds:
```
// Interval update in milliseconds
const pollingInterval = 3 * 1000;
```

As <u>REQUIREMENT</u>, the endpoint should return a JSON array. Example:
```
{ "id": 1, "name": "Alice", "age": 25 },
{ "id": 2, "name": "Bob", "age": 30 },
{ "id": 3, "name": "Charlie", "age": 35 }
```
<u>**Run:**</u> `make build run`

Now the grids are available using the URLs based on ID accordingly

Example: `http://localhost:8090?id=3`

<a name="configuration"></a>
## 3. How to configure the grid

In order to configure the `behavior` of the grid, enhance the: 
```
const gridOptions = {};
```

In order to configure the style of cells based on specific conditions, enhance the:
```
const cellClassRules = {};
```

<a name="distribution"></a>
## 4. Distribution - Version

In order to change between `community` or `enterprise` distribution or change package version alter the `.html` file. Example:
```
<!--Community package-->
<script src="https://cdn.jsdelivr.net/npm/ag-grid-community@32.3.3/dist/ag-grid-community.js?t=1731498596308"></script>
    
<!--Enterprise package-->
<script src="https://cdn.jsdelivr.net/npm/ag-grid-charts-enterprise@32.3.3/dist/ag-grid-charts-enterprise.js"></script>
```
<a name="todo"></a>
## 5. TODO

- Move Pagination from FE to server side
- Place the option of a horizontal scroll bar
- Custome color and rendering per cell or column etc.
- Custom schema rendering
- The user has to save its rendering options

