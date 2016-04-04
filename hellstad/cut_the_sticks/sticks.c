#include <math.h>
#include <stdio.h>
#include <string.h>
#include <stdlib.h>
#include <assert.h>
#include <limits.h>
#include <stdbool.h>

int calc_umin(int *p_arr, int arr_size) {
    if (arr_size < 1) return -1;
    if (arr_size == 1) return p_arr[0];
    
    int min = p_arr[0];
    for (int i=1; i<arr_size; i++) {
        if (p_arr[i] < min) {
            min = p_arr[i];
        }
    }
    return min;
}

void cut_sticks(int *p_arr, int *p_arr_size) {
    int min = calc_umin(p_arr, *p_arr_size);
    int new_size = 0;

    for (int i=0; i<*p_arr_size; i++) {
        p_arr[i] = p_arr[i] - min;
        if (p_arr[i] > 0) new_size++;
    }

    int *new_arr = malloc(new_size * sizeof(int));
    int new_arr_count = 0;
    for (int i=0; i<*p_arr_size; i++) {
        if (p_arr[i] > 0) {
            new_arr[new_arr_count++] = p_arr[i];
        }
    }

    realloc(p_arr, new_size * sizeof(int));
    for (int i=0; i<new_size; i++) {
        p_arr[i] = new_arr[i];
    }
    free(new_arr);
    *p_arr_size = new_size;
}

void i_will_cut_you(int *p_arr, int arr_size) {
    int p_arr_size = arr_size;
    while (p_arr_size > 0) {
        printf("%d\n", p_arr_size);
        cut_sticks(p_arr, &p_arr_size);
    }
}

int main() {
    int n; 
    scanf("%d", &n);
    int *arr = malloc(n * sizeof(int));
    for (int arr_i=0; arr_i<n; arr_i++) {
       scanf("%d", (arr + arr_i));
    }
    
    i_will_cut_you(arr, n);
    
    return 0;
}

