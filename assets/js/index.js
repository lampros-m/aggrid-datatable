let gridApi;

// Grid options
const gridOptions = {
    defaultColDef: {
        flex: 1,
        columnChooserParams: {
            // Prevent the layout from updating when columns are rearranged
            suppressSyncLayoutWithGrid: true,
        },
        // Columns filtering
        filter: true,
    },
    // If catch a significant value, mark red the whole row
    getRowStyle: params => {
        if (params.data.logIndex > 40) {
            return { background: 'red' };
        }
    },
    // Pagination
    pagination: true,
    rowSelection: { mode: "multiRow", headerCheckbox: false }
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

            function convert(sourceData) {
                const convertedData = {};
                for (const column of schema) {
                    let current = sourceData;
                    for (const part of column.data.split('.')) {
                        current = current?.[part];
                    }
                    convertedData[column.title] = column.customRender 
                        ? column.customRender(current) 
                        : current;
                }
                return convertedData;
            }

            if (schema) data = data.map(convert);            

            // Collect all unique keys from the data, excluding the predefined ones
            const existingFields = gridOptions.columnDefs?.map(def => def.field) ?? [];

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

            // Update gridOptions with new column definitions
            gridOptions.columnDefs = [
                ...(gridOptions.columnDefs ?? []), // Predefined columns
                ...newColumnDefs // Dynamically added columns
            ];

            // Sort the columns alphabetically
            const sortedColumnDefs = gridOptions.columnDefs.sort((a, b) => a.field.localeCompare(b.field));
            gridOptions.columnDefs = sortedColumnDefs;

            // Update the grid data
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

