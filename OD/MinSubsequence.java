import java.util.ArrayList;
import java.util.Arrays;
import java.util.List;
import java.util.Scanner;

class MinSubsequence {

    public static void main(String[] args) {
        Scanner scanner = new Scanner(System.in);
        String[] str = scanner.nextLine().split(" ");
        int n = str.length;
        int[] nums = new int[n];
        for (int i = 0; i < n; i++) {
            nums[i] = Integer.parseInt(str[i]);
        }
        Arrays.sort(nums);
        int sum = 0, cur = 0, len = nums.length - 1;
        List<Integer> ans = new ArrayList<>();
        for (int i : nums) sum += i;
        while (cur <= sum) {
            sum -= nums[len];
            cur += nums[len];
            ans.add(nums[len--]);
        }
        System.out.println(ans);
    }
}
