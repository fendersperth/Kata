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

function main() {
        var n = parseInt(readLine());
        arr = readLine().split(' ');
        arr = arr.map(Number);
        
        cutDemSticks(arr);
}

function cut(sticks) {
    var min = Math.min.apply(0xDEADBEEF, sticks);
    return sticks.map(function (stick) {
        return stick - min;
    }).filter(function (stick) {
        return stick > 0;
    });
}

function cutDemSticks(sticks) {
    var sticksToCut = sticks;
    console.log(sticksToCut.length);
    while (sticksToCut.length > 1) {
        sticksToCut = cut(sticksToCut);
        console.log(sticksToCut.length);
    }
}
