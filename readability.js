var DOMPurify = require('dompurify');
var { Readability } = require('@mozilla/readability');
var { JSDOM } = require('jsdom');
function clear(url,content,option) {
    if (!url){
        url=undefined
    }
    return new Readability(new JSDOM(DOMPurify(new JSDOM('').window).sanitize(content), {  url: url}).window.document,option).parse()
}






