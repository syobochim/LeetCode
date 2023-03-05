function runningSum(nums: number[]): number[] {
    let result: number[] = [];
    let prevSum: number = 0;
    nums.forEach((num: number, i: number) => {
        prevSum = prevSum + num
        result.push(prevSum);
    })
    return result;
};

function test(): void {
    console.log(runningSum([1,2,3,4]))
    console.log(runningSum([1,1,1,1,1]))
    console.log(runningSum([3,1,2,10,1]))
}

test()