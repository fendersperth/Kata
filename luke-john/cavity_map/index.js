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

var checkMap = function checkMap(n, map) {
  var cavityMap = map.slice();
  for (var y = 1; y < n - 1; y++) {
    for (var x = 1; x < n - 1; x++) {
      var cavity = map[y][x] - map[y - 1][x] > 0 && map[y][x] - map[y][x + 1] > 0 && map[y][x] - map[y + 1][x] > 0 && map[y][x] - map[y][x - 1] > 0;
      if (cavity) {
        cavityMap[y] = cavityMap[y].substr(0, x) + 'X' + cavityMap[y].substr(x + 1);
      }
    }
  }
  return cavityMap.join('\n');
};

function main() {
    var n = parseInt(readLine());
    var grid = [];
    for(var grid_i = 0; grid_i < n; grid_i++){
       grid[grid_i] = readLine();
    }
    process.stdout.write(checkMap(n, grid));
}
