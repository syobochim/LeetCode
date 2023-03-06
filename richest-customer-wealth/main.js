function maximumWealth(accounts) {
    //    let richestNumber: number = 0;
    //    for (const account of accounts) {
    //        const sumAccount = account.reduce((sum, num) => sum + num, 0);
    //        if (richestNumber < sumAccount) {
    //            richestNumber = sumAccount;
    //        }
    //    }
    //    return richestNumber;
    return Math.max.apply(Math, accounts.map(function (account) { return account.reduce(function (sum, num) { return sum + num; }, 0); }));
}
;
function test() {
    console.log(maximumWealth([[1, 2, 3], [3, 2, 1]])); // -> 6
    console.log(maximumWealth([[1, 5], [7, 3], [3, 5]])); // -> 10
    console.log(maximumWealth([[2, 8, 7], [7, 1, 3], [1, 9, 5]])); // -> 17
}
;
test();
