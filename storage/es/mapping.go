package es

var mappings = `
{
    "properties": {
        "title": {
            "type": "string"
        },
        "author_url": {
            "type": "string",
            "index": "not_analyzed"
        },
        "author": {
            "type": "string"
        },
        "reply": {
            "type": "integer"
        },
        "last_reply_time": {
            "type": "date"
        },
        "topic_content": {
            "properties": {
                "update_time": {
                    "type": "date"
                },
                "content": {
                    "type": "string"
                },
                "with_pic": {
                    "type": "boolean"
                },
                "pic_urls": {
                    "type": "string",
                    "index": "not_analyzed"
                },
                "like": {
                    "type": "integer"
                }
            }
        }
    }
}
`
