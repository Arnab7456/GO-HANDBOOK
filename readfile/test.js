const fs = require('fs')
 
const ans = fs.readFileSync("data.txt").toString();
console.log(ans);
