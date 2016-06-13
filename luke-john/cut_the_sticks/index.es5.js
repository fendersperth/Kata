"use strict";

process.stdin.resume();
process.stdin.setEncoding('ascii');

var input_stdin = "";
var input_stdin_array = "";
var input_currentline = 0;

process.stdin.on('data', function (data) {
    input_stdin += data;
});

process.stdin.on('end', function () {
    input_stdin_array = input_stdin.split("\n");
    main();    
});

function readLine() {
    return input_stdin_array[input_currentline++];
}

/////////////// ignore above this line ////////////////////

function _toConsumableArray(arr) { if (Array.isArray(arr)) { for (var i = 0, arr2 = Array(arr.length); i < arr.length; i++) { arr2[i] = arr[i]; } return arr2; } else { return Array.from(arr); } }

var cut = function cut(bundle) {
  console.log(bundle.length);
  var min = Math.min.apply(Math, _toConsumableArray(bundle));
  var cutBundle = bundle.map(function (stick) {
    return stick - min;
  }).filter(function (stick) {
    return stick > 0;
  });
  if (cutBundle.length) {
    cut(cutBundle);
  }
};

function main() {
    var n = parseInt(readLine());
    var arr = readLine().split(' ');
    arr = arr.map(Number);
    cut(arr)
}
