function main() {
    var t = parseInt(readLine());
    for(var a0 = 0; a0 < t; a0++){
        var n_temp = readLine().split(' ');
        var n = parseInt(n_temp[0]);
        var c = parseInt(n_temp[1]);
        var m = parseInt(n_temp[2]);
        
        go(n, c, m)
    }
}

function go(amount, price, exchange) {
    var eaten = 0;
    var bought = parseInt(amount / price);
    eaten += bought;
    
    var wrappers = bought;
    while(wrappers >= exchange) {
        var more = parseInt(wrappers / exchange);
        eaten += more;
        wrappers = wrappers % exchange + more
    }
    console.log(eaten)
}