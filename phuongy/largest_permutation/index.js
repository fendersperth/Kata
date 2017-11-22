var lookup = {};
var numbers = [];

function processData(input) {
    var lines = input.split('\n');
    var line1 = lines[0]
                    .split(' ')
                    .map(Number);

    
    numbers = lines[1]
                    .split(' ')
                    .map(Number)
                    .filter(function(n) {
                        return typeof n === "number";
                    });
    
    var n = line1[0];
    var k = line1[1];
   
    createLookup(numbers);

    var i = 0;
    
    while(k > 0 && i < n) {
        
        if(numbers[i] != lookup[i]) {  
            doPermutation(i);
            k--;
        }
        
        i++;        
    }
    
    console.log(numbers.join(' '));
    
} 

function createLookup(numbers) {
    
    lookup.numbers = numbers
                        .slice()
                        .map(function(value) {
                            return {max: 0, value: value};
                        });
    
    lookup.max = numbers
                    .slice()
                    .map(function(value, index) {
                        return {index: index, value: value};
                    })
                    .sort(function(a, b) {
                        return b.value - a.value;
                    });
    
    lookup.max.forEach(function(value, index) {
       lookup.numbers[value.index].max = index;
    });
}

function doPermutation(i) {
    var max = lookup.max[i];
    var number = lookup.numbers[i];
    
    var current_number_object = lookup.numbers[i];
    var max_number_object = lookup.numbers[max.index];
    
    // swap
    numbers[i] = max_number_object.value; 
    numbers[max.index] = current_number_object.value;
    
    // update numbers lookup table indexes
    lookup.numbers[i] = max_number_object;
    lookup.numbers[max.index] = current_number_object;
    
    // update max numbers lookup table indexes
    lookup.max[i].index = i;
    lookup.max[number.max].index = max.index;
}