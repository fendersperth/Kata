const min = array => array.reduce((p, c) => Math.min(p, c), Infinity)
const subtract = a => b => b-a
const positive = x => x>0

const cutTheSticks = list => {
    if(!positive(list.length))
        return
    console.log(list.length)
    const currentMin = min(list)
    const cutList = list.map(subtract(currentMin))
                        .filter(positive)
    cutTheSticks(cutList)
}

const readLine = input => {
    const sticks = input.toString().trim().split(' ')
    cutTheSticks(sticks)
    console.log('more:\n')
}

// // get input from user
// const stdin = process.openStdin()
// stdin.addListener("data", data => readLine(data) )

// // // too lazy to get user input
// const sticks = [5, 4, 4, 2, 2, 8]
// cutTheSticks(sticks)

module.exports = cutTheSticks