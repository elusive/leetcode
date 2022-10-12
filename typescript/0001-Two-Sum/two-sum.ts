
export default function TwoSum(values: number[], target: number): number[] {
    
    // outer loop thru each number until solution found
    for (let v = 0; v < values.length; v++) {
        const val1 = values[v];

        // inner loop to search from index to the right looking for other 
        // operand that when added to val1 will equal the target value
        for (let t = v + 1; t < values.length; t++) {
            const val2 = values[t];
            let currentSum = val1 + val2;
            if (currentSum === target) {
                return ([v, t]);
            }
        }
    } 

    // should not get here due to constraint
    // of one correct answer for each instance
    // see readme.md
    return [];
}
