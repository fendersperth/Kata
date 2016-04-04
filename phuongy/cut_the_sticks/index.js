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
    
    var sticks = arr;
    
    while(sticks.length > 0) {
        console.log(sticks.length);
        sticks = performCut(sticks);
    }
}

function performCut(sticks) {
    var min = getMin(sticks);
    var next = [];
    
    for(var i = 0; i < sticks.length; i++) {
        if(sticks[i] - min > 0) {
            next.push(sticks[i] - min);
        } 
    }
    
    return next;
}

function getMin(numbers) {
    var min = numbers[0];
    
    for(var i = 0; i < numbers.length; i++) {
        if(numbers[i] < min) {
            min = numbers[i];
        }
    }
    
    return min;
}