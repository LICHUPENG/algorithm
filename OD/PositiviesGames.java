import java.util.ArrayList;
import java.util.Scanner;

public class PositiviesGames {
    public static void main(String[] args) {
        Scanner sc = new Scanner(System.in);
        String s = sc.nextLine();
        String[] split = s.split(" ");
        int m = Integer.parseInt(split[0]);
        int n = Integer.parseInt(split[1]);

        if (n == 1 || m == 1) {
            System.out.println(n);
        }
        ArrayList<Integer> list = new ArrayList<>();
        for (int i = 1; i <= m; i++) {
            list.add(i);
        }
        // 构造一个全局计数 从1 开始 一直数 知道 list的长度为1
        int count = 1;
        while (list.size() > 1) {
            int tmpSize;
            tmpSize = list.get(0);
            list.remove(0);
            if (count % n == 0 || Integer.toString(count).endsWith(String.valueOf(n))) { // 当前报的数 为 m的倍数 或者 以m结尾
            } else {
                list.add(tmpSize);
            }
            count++;
        }
        System.out.println(list.get(0));
    }
}
