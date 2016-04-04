const Benchmark = require('benchmark')

const suite = new Benchmark.Suite
const cutTheSticks  = require('./index')

const tests = [
    {in: [5, 4, 4, 2, 2, 8]},
    {in: [1, 2, 3, 4, 3, 3, 2, 1]},
]

suite.add('blabla', () =>
        cutTheSticks(tests[0].in))
    .on('complete', event => {
        console.log(event)
    })
    .run()