# AG-Grid DataTable

1. [ Overview. ](#overview)
2. [ How to run. ](#run)
3. [ Configuration. ](#configuration)
4. [ Distribution - Version. ](#distribution)
5. [ TODO ](#todo)

<a name="overview"></a>
## 1. Overview

Creates a polling based grid using JS package: https://www.ag-grid.com/

<a name="run"></a>
## 2. How to run

Replace under `assets/js/index.js` the GET endpoint from which the JSON data are retrieved.
```
const endpoint = 'http://localhost:8090/data';
```

As requirement, the endpoint should return a JSON array. Example:
```
{ "id": 1, "name": "Alice", "age": 25 },
{ "id": 2, "name": "Bob", "age": 30 },
{ "id": 3, "name": "Charlie", "age": 35 }
```

Enhance your `.html` file with the script:
```
<script src="/js/index.js"></script>
```

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

- Move Pagination from FE to Server Side
- Place the option of a horizontal scroll bar
- Create the opion of custom render function per cell

