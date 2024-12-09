let gridApi;

// Grid options
const gridOptions = {
    defaultColDef: {
        flex: 1,
        columnChooserParams: {
            suppressSyncLayoutWithGrid: true, // Prevent the layout from updating when columns are rearranged
        },
        filter: true, // Columns filtering
    },
    pagination: true, // Pagination
    rowSelection: { mode: "multiRow", headerCheckbox: false },
};

// Cell class rules
const cellClassRules = {
    "int-positive": params => Number.isInteger(params.value) && params.value > 0,
    "int-negative": params => Number.isInteger(params.value) && params.value < 0,
    "int-zero": params => Number.isInteger(params.value) && params.value === 0,    
};

// Interval update in milliseconds
const pollingInterval = 3 * 1000;

// Endpoint const
const endpoint = 'http://localhost:8090/data';

document.addEventListener("DOMContentLoaded", () => {
    // Initial fetch
    fetchDataAndUpdateGrid();

    // Set interval for polling
    setInterval(fetchDataAndUpdateGrid, pollingInterval);
});

const fetchDataAndUpdateGrid = () => {
    fetch(endpoint)
        .then(response => response.json())
        .then(data => {
            // Ensure the data is an array and not empty
            if (!Array.isArray(data) || data.length === 0) {
                console.error('No data received or data is not an array.');
                return;
            }

            // Collect all unique keys from the data to ensure all columns are covered
            const allKeys = new Set();
            data.forEach(item => {
                Object.keys(item).forEach(key => allKeys.add(key));
            });

            // Create column definitions from the keys
            const columnDefs = Array.from(allKeys).map(key => ({ field: key }));

            // Apply format rules to columns
            columnDefs.forEach((colDef) => {
                colDef.cellClassRules = cellClassRules
            });

            // Update the grid options with the data and column definitions
            gridOptions.columnDefs = columnDefs;
            gridOptions.rowData = data;

            // Create the grid
            var gridDiv = document.querySelector("#myGrid");

            // If gridApi doesnt exist, create the grid. else update the grid with new data
            if (gridApi) {
                gridApi.setGridOption('rowData', data);
            } else {
                gridApi = agGrid.createGrid(gridDiv, gridOptions);
            }
        })
        .catch(error => {
            console.error('Error fetching data:', error);
        });
};

