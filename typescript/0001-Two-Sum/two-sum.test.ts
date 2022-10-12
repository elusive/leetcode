import TwoSum from "./two-sum";

test("TwoSum function should return expected value.", () => {
    expect(TwoSum([2,7,11,15], 9)).toEqual([0,1]);
    expect(TwoSum([3,2,4], 6)).toEqual([1,2]);
    expect(TwoSum([3,3], 6)).toEqual([0,1]);
});
