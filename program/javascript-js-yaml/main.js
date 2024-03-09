const yaml = require('js-yaml');
const fs = require('fs');

let path = process.argv[2];
let text = fs.readFileSync(path, 'utf8');

let obj = yaml.load(text);
console.log(obj);
