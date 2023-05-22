import java.util.Scanner;

public class ServiceBroadcast {
    public static int count = 0;
    public static void main(String[] args) {
        Scanner in = new Scanner(System.in);
        String[] str = in.nextLine().split(" ");
        int n = str.length;
        int[][] arr = new int[n][n];
        for(int i = 0; i < n; i++) {  // 把第一行加入arr
            arr[0][i] = Integer.parseInt(str[i]);
        }
        for(int i = 1; i < n; i++) {  // 把剩下的行加入arr
            String[] s = in.nextLine().split(" ");
            for(int j = 0; j < n; j++) {
                arr[i][j] = Integer.parseInt(s[j]);
            }
        }
        boolean[] visited = new boolean[n];
        for(int i = 0; i < n; i++) {
            if(!visited[i]) {
                dfs(arr, visited, i);
            }
        }
        System.out.println(count);
    }
    public static void dfs(int[][] arr, boolean[] visited, int index) {
        visited[index] = true;
        boolean flag = true;
        for (int i = index + 1; i < arr.length; i++) {
            if (arr[index][i] == 1) {
                flag = false;
                dfs(arr, visited, i);
            }
        }
        if (flag) {
            count++;
        }
    }
}
