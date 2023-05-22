import java.util.Arrays;
import java.util.Scanner;

/**
 * 高效的任务规划
 * 题目描述：
 * 你有 n 台机器编号为 1~n，每台都需要完成完成一项工作，机器经过配置后都能完成独立完成一项工作。
 * 假设第 i 台机器你需要花 B 分钟进行设置，然后开始运行，J 分钟后完成任务。
 * 现在，你需要选择布置工作的顺序，使得用最短的时间完成所有工作。
 * 注意，不能同时对两台机器进行配置，但配置完成的机器们可以同时执行它们各自的工作。
 * 注：此题对效率有要求，请考虑高效的实现方式
 *
 * 输入描述：
 * 第一行输入代表总共有 M 组任务数据（1<M<=10）。
 * 每组数第一行为一个整数，指定机器的数量 N（0<N<=1000）。
 * 随后的 N 行每行两个整数，第一个表示 B（0<=B<=10000），第二个表示 J（0<=J<=10000）。
 * 每组数据连续输入，不会用空行分隔。
 * 各组任务单独计时。
 * 输出描述：
 * 对于每组任务，输出最短完成时间，且每组的结果独占一行。
 * 例如，两组任务就应该有两行输出。
 * 示例 1：
 * 1
 * 1
 * 2 2
 * 输出4
 */

public class AssignmentArrangement {
    public static void main(String[] args) {
        Scanner in = new Scanner(System.in);
        int M = in.nextInt();
        for(int m = 0; m < M; m++) {
            int N = in.nextInt();
            // 动态规划，让任务工作时间最长的机器先运行，这样总体时间才最短啊
            int[] dp = new int[N]; // dp[i]第i台机器工作完的时间
            int last = 0; // 之前机器配置完成的时间
            int res = 0; //总时间
            int[][] machine = new int[N][2];  // 每组任务的N台机器的配置时间和工作时间
            for (int i = 0; i < N; i++) {
                int B = in.nextInt();
                int J = in.nextInt();
                machine[i][0] = B;
                machine[i][1] = J;
            }
            // lambda按第二元素降序排序，即按工作时间降序排序
            Arrays.sort(machine, (e1, e2) -> (e2[1] - e1[1]));
            for (int i = 0; i < N; i++) {
                dp[i] = last + machine[i][0] + machine[i][1];
                last += machine[i][0];
                res = Math.max(res, dp[i]);
            }
            System.out.println(res);
        }
    }
}
