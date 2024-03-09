const fs = require('fs');
const YAML = require('yaml');

let path = process.argv[2];
let text = fs.readFileSync(path, 'utf8');

let obj = YAML.parse(text);
console.log(obj);
