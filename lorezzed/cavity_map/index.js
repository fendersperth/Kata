'use strict'
const input = [
    [1,1,1,2],
    [1,9,1,2],
    [1,8,9,2],
    [1,2,3,4],
]
// keep input intact to compare on, but mutate render array
let render = input

const check = (height, width) => {
    // check bounds
    if(!input[height-1] || !input[height+1])
        return
    if(!input[height][width-1] || !input[height][width+1])
        return
    // compare greaterThan on neighbours
    const cell = input[height][width]
    if( cell > input[height-1][width]
    &&  cell > input[height+1][width]
    &&  cell > input[height][width-1]
    &&  cell > input[height][width+1]
    )
        render[width][height] = 'X'
}

for(let height=0; height < input.length; height++) {
    for(let width=0; width < input.length; width++) {
        check(height, width)
    }
}

render.map(row => console.log(row.join('') ))