"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
var express = require("express");
var app = express();
app.get('/', function (req, res) {
    res.end('hello, Mars');
});
app.listen(3000, function () {
    console.log('server is listening');
});
//# sourceMappingURL=server.js.map