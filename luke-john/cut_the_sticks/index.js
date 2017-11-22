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

const cut = (bundle) => {
  console.log(bundle.length)
  const min = Math.min(...bundle)
  const cutBundle = bundle.map(stick => stick - min)
    .filter(stick => stick > 0)
  if (cutBundle.length) {
    cut(cutBundle)
  }
}

function main() {
    var n = parseInt(readLine());
    var arr = readLine().split(' ');
    arr = arr.map(Number);
    cut(arr)
}
