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

function SquareGrid(gridLines) {
    var _vertexDistance;
    var _ordered;
    var _edges;
    var _cavities;
    var _matrix;
    var _gridLines;

    _gridLines = gridLines;
    setupGrid();
    findEdges();
    findCavities();
    
    function setupGrid() {
        var ordered = [];
        _gridLines.forEach(function (line, i) {
            line.split('').forEach(function (char) {
                ordered.push(parseInt(char));
            });
        });
        _vertexDistance = _gridLines[0].length;
        _ordered = ordered;

        var matrix = [];
        for (var x=0; x<_vertexDistance; x++) {
            var line = [];
            for (var y=0; y<_vertexDistance; y++) {
                line.push(_ordered[x * _vertexDistance + y]);
            }
            matrix.push(line);
        }
        _matrix = matrix;
    }

    function printGrid() {
        var strGrid = [];
        for (var i=0; i<_vertexDistance; i++) {
            var line = '';
            for (var j=0; j<_vertexDistance; j++) {
                var n = i * _vertexDistance + j;
                if (!isEdge(n) && isCavity(i, j)) {
                    line += 'X';
                } else {
                    line += _ordered[n];
                }
            }
            strGrid.push(line);
        }
        console.log(strGrid.join('\n'));
    }

    function findEdges() {
        var edges = [];
        _ordered.forEach(function (num, n) {
            if (n < _vertexDistance) edges.push(n); // top
            if (n / _vertexDistance >= (_vertexDistance - 1)) edges.push(n); // bottom
            if (!(n % _vertexDistance)) edges.push(n - 1, n); // sides
        });
        _edges = edges;
    }

    function findCavities() {
        var cavities = [];
        for (var i=0; i<_vertexDistance; i++) {
            for (var j=0; j<_vertexDistance; j++) {
                var n = i * _vertexDistance + j;
                if (!isEdge(n)) {
                    if (isCavity(i, j)) cavities.push(n);
                }
            }
        }
        _cavities = cavities;
    }

    function isCavity(i, j) {
        var surrounds = [];
        surrounds.push(
            _matrix[i][j-1],
            _matrix[i][j+1],
            _matrix[i-1][j],
            _matrix[i+1][j]
        );

        var valueAtN = _matrix[i][j];
        return Math.max.apply(0xDEADBEEF, surrounds) < valueAtN;
    }

    function isEdge(n) {
        return (_edges.indexOf(n) !== -1);
    }

    return {
        printGrid: printGrid
    };
}

function main() {
    var n = parseInt(readLine());
    var grid = [];
    for(var grid_i = 0; grid_i < n; grid_i++){
       grid[grid_i] = readLine();
    }

    var sqGd = new SquareGrid(grid);
    sqGd.printGrid();
}
