const schema = [
    {
        "data": "chain",
        "title": "Chain"
    },

    {
        "data": "poolAddress",
        "title": "Pool Address",
        "customRender": data => `${data} 123!`
    },

    {
        "data": "extraData.attribute1",
        "title": "Attribute #1"
    },

    {
        "data": "extraData.list.1",
        "title": "List 2nd Value"
    }

];