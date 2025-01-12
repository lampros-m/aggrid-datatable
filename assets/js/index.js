let gridApi;

// Grid options
const gridOptions = {
    initialState: initialState,
    defaultColDef: {
        flex: 1,
        columnChooserParams: {
            // Prevent the layout from updating when columns are rearranged
            suppressSyncLayoutWithGrid: true,
        },
        // Columns filtering
        filter: true,
    },
    // Pagination
    pagination: true,
    rowSelection: { mode: "multiRow", headerCheckbox: false }
};

// Interval update in milliseconds
const pollingInterval = 3 * 1000;

// Endpoint const
const endpoint = 'http://localhost:8090/data';

async function delayAsync(ms) {
    return new Promise(resolve => setTimeout(resolve, ms));
}

document.addEventListener("DOMContentLoaded", async () => {
    while (true) {
        await fetchDataAndUpdateGrid();
        await delayAsync(pollingInterval);
    }
});

function saveState() {

    console.log(gridApi.getState());
}

async function fetchDataAndUpdateGrid() {

    try {
        const response = await fetch(endpoint);

        const data = await response.json();

        // Ensure the data is an array and not empty
        if (!Array.isArray(data) || data.length === 0) {
            console.error('No data received or data is not an array.');
            return;
        }

        // Create the grid
        var gridDiv = document.querySelector("#myGrid");

        // If gridApi doesnt exist, create the grid. else update the grid with new data
        if (gridApi) {
            gridApi.setGridOption('rowData', data);
        } else {

            if (!schema) {

                // Get all existing fields from all rows
                const newKeys = new Set();
                data.forEach(item => {
                    Object.keys(item).forEach(key => {
                        if (!existingFields.includes(key)) {
                            newKeys.add(key);
                        }
                    });
                });
    
                // Create additional column definitions for new keys
                const newColumnDefs = Array.from(newKeys).map(key => ({
                    field: key,
                }));
    
                // Sort the columns alphabetically
                const sortedColumnDefs = newColumnDefs.sort(
                    (a, b) => a.field.localeCompare(b.field));
    
                // Update gridOptions with new column definitions
                gridOptions.columnDefs = sortedColumnDefs;              
            }
            else {
    
                // Simply use whatever our user gave us without any changes
                gridOptions.columnDefs = schema.columnDefs;
            }
    
            // Update the grid data
            gridOptions.rowData = data;

            gridApi = agGrid.createGrid(gridDiv, gridOptions);
        }
    }
    catch (error) {    
        console.error('Error fetching data:', error);
    }
}
