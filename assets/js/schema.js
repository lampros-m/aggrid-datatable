const schema = {

    columnDefs: [
        {
            "headerName": "Chain",
            "field": "chain"        
        },
        {
            "headerName": "Pool Address",
            "valueGetter": params => `${params.data.poolAddress} 123!`          
        },
        {
            "headerName": "Attribute #1",
            "field": "extraData.attribute1"
        },
        {
            "headerName": "List 2nd Value",
            "field": "extraData.list.1"
        },
        {
            "headerName": "Log Index",
            "field": "logIndex",
            "cellStyle": params => {
                if (params.value > 40) {
                    return { background: 'red' };
                }
            },
        }

    ]

};

