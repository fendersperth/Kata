const min = array => array.reduce((p, c) => Math.min(p, c), Infinity)
const subtract = a => b => b-a
const positive = x => x>0

const cutTheSticks = list => {
    console.log(list.length)
    if(list.length <= 1)
        return
    const currentMin = min(list)
    const cutList = list.map(subtract(currentMin))
                        .filter(positive)
    cutTheSticks(cutList)
}

const readLine = input => {
    const sticks = input.toString().trim().split(' ')
    cutTheSticks(sticks)
    console.log('\nmore:')
}

const stdin = process.openStdin()
stdin.addListener("data", data => readLine(data) )