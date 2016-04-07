function main() {
    
    var t = parseInt(readLine());
    for(var a0 = 0; a0 < t; a0++){
        var n_temp = readLine().split(' ');
        var n = parseInt(n_temp[0]);
        var c = parseInt(n_temp[1]);
        var m = parseInt(n_temp[2]);
        
        var paid_chocolates = parseInt(n / c);
        var free_chocolates = parseInt(paid_chocolates / m);
        var total_chocolates = paid_chocolates + free_chocolates;
        
        while(parseInt(total_chocolates / m) > free_chocolates) {
            free_chocolates = parseInt(total_chocolates / m);
            total_chocolates = paid_chocolates + free_chocolates;
        }
        
        console.log(total_chocolates);
    }

}