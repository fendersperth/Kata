function main() {
    var n = parseInt(readLine());
    var grid = [];
    for(var grid_i = 0; grid_i < n; grid_i++){
       grid[grid_i] = readLine().split("");
    }

    for(var x = 0; x < n; x++) {
        for(var y = 0; y < n; y++) {
            if(isCavity(grid, x, y)) {
                grid[x][y] = 'X';
            }
        }
        
        console.log(grid[x].join(""));
    }
}

function isCavity(grid, x,y) {
    var currentValue = grid[x][y];
    var isCavity = false;
    var n = grid.length;
    
    if(x > 0 && x < (n-1) && y > 0 && y < (n-1)) {
    
        if(currentValue > getTopEdge(grid, x, y) 
            && currentValue > getBottomEdge(grid, x, y) 
            && currentValue > getRightEdge(grid, x, y) 
            && currentValue > getLeftEdge(grid, x, y)) {
            isCavity = true;
        }
    }
    
    return isCavity;
}

function getTopEdge(grid, x, y) {
    return grid[x-1][y];
}

function getBottomEdge(grid, x, y) {
    return grid[x+1][y];
}

function getRightEdge(grid, x, y) {
    return grid[x][y+1];
}

function getLeftEdge(grid, x, y) {
    return grid[x][y-1];
}